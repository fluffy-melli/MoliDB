package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type TokenBucket struct {
	capacity   int
	tokens     float64
	refillRate float64
	lastRefill time.Time
	lastSeen   time.Time
	mutex      sync.Mutex
}

type MiddlewareConfig struct {
	Capacity   int
	RefillRate float64
}

func NewTokenBucket(capacity int, refillRate float64) *TokenBucket {
	now := time.Now()
	return &TokenBucket{
		capacity:   capacity,
		tokens:     float64(capacity),
		refillRate: refillRate,
		lastRefill: now,
		lastSeen:   now,
	}
}

func (tb *TokenBucket) AllowRequest(currentTime time.Time) bool {
	tb.mutex.Lock()
	defer tb.mutex.Unlock()

	elapsed := currentTime.Sub(tb.lastRefill).Seconds()
	tb.tokens += elapsed * tb.refillRate
	if tb.tokens > float64(tb.capacity) {
		tb.tokens = float64(tb.capacity)
	}
	tb.lastRefill = currentTime
	tb.lastSeen = currentTime

	if tb.tokens >= 1 {
		tb.tokens -= 1
		return true
	}
	return false
}

func (tb *TokenBucket) Reset(capacity int, refillRate float64) {
	tb.mutex.Lock()
	defer tb.mutex.Unlock()

	tb.capacity = capacity
	tb.refillRate = refillRate
	tb.tokens = float64(capacity)
	now := time.Now()
	tb.lastRefill = now
	tb.lastSeen = now
}

var TokenBucketPool = sync.Pool{
	New: func() interface{} {
		return NewTokenBucket(100, 100.0/60.0)
	},
}

var tokenBuckets sync.Map

const CleanupInterval = 5 * time.Minute

const InactiveDuration = 10 * time.Minute

func init() {
	go cleanupTokenBuckets()
}

func cleanupTokenBuckets() {
	ticker := time.NewTicker(CleanupInterval)
	defer ticker.Stop()

	for {
		<-ticker.C
		now := time.Now()
		tokenBuckets.Range(func(key, value interface{}) bool {
			tb := value.(*TokenBucket)

			tb.mutex.Lock()
			lastSeen := tb.lastSeen
			tb.mutex.Unlock()

			if now.Sub(lastSeen) > InactiveDuration {
				tokenBuckets.Delete(key)
				TokenBucketPool.Put(tb)
			}
			return true
		})
	}
}

func getOrCreateTokenBucket(ip string, capacity int, refillRate float64) *TokenBucket {
	if v, ok := tokenBuckets.Load(ip); ok {
		return v.(*TokenBucket)
	}

	newBucket := TokenBucketPool.Get().(*TokenBucket)
	newBucket.Reset(capacity, refillRate)

	actual, loaded := tokenBuckets.LoadOrStore(ip, newBucket)
	if loaded {
		TokenBucketPool.Put(newBucket)
		return actual.(*TokenBucket)
	}
	return actual.(*TokenBucket)
}

func TokenBucketMiddleware(config MiddlewareConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		currentTime := time.Now()

		tb := getOrCreateTokenBucket(ip, config.Capacity, config.RefillRate)

		if !tb.AllowRequest(currentTime) {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "Too many requests"})
			c.Abort()
			return
		}
		c.Next()
	}
}
