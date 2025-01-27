package middleware

import (
	"fmt"
	info "runtime"
	"time"

	"github.com/fluffy-melli/MoliDB/internal/handlers"
	"github.com/fluffy-melli/MoliDB/internal/runtime"
	"github.com/gin-gonic/gin"
)

type MemoryInfo struct {
	Alloc     string  `json:"alloc"`
	HeapAlloc string  `json:"heap-alloc"`
	HeapSys   string  `json:"heap-system"`
	HeapIdle  string  `json:"heap-idle"`
	GCCPU     float64 `json:"gccpu"`
}

type SystemInfoRespond struct {
	Version string     `json:"version"`
	Runtime string     `json:"runtime"`
	Memory  MemoryInfo `json:"memory"`
}

func byteMap(bytes uint64) string {
	if bytes < 1024 {
		return fmt.Sprintf("%d (Byte)", bytes)
	} else if bytes < 1024*1024 {
		return fmt.Sprintf("%.2f (KB)", float64(bytes)/1024)
	} else if bytes < 1024*1024*1024 {
		return fmt.Sprintf("%.2f (MB)", float64(bytes)/(1024*1024))
	} else {
		return fmt.Sprintf("%.2f (GB)", float64(bytes)/(1024*1024*1024))
	}
}

func timeMap(elapsed int64) string {
	days := elapsed / (1000 * 60 * 60 * 24)
	hours := (elapsed / (1000 * 60 * 60)) % 24
	minutes := (elapsed / (1000 * 60)) % 60
	seconds := (elapsed / 1000) % 60
	milliseconds := elapsed % 1000
	return fmt.Sprintf("%dd %dh %dm %d.%ds", days, hours, minutes, seconds, milliseconds)
}

// SystemInfo returns the system information
// @Summary Get system information
// @Description Returns system memory stats, runtime duration, and version
// @Tags information
// @Accept json
// @Produce json
// @Success 200 {object} SystemInfoRespond
// @Router /system [get]
func SystemInfo(c *gin.Context) {
	var m info.MemStats
	info.ReadMemStats(&m)
	respond := SystemInfoRespond{
		Runtime: timeMap(time.Now().UnixMilli() - runtime.StartTime),
		Version: "v1.3.0",
		Memory: MemoryInfo{
			Alloc:     byteMap(m.Alloc),
			HeapAlloc: byteMap(m.HeapAlloc),
			HeapSys:   byteMap(m.HeapSys),
			HeapIdle:  byteMap(m.HeapIdle),
			GCCPU:     m.GCCPUFraction,
		},
	}
	handlers.SendResponse(c, 200, respond)
}
