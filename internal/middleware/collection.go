package middleware

import (
	"os"

	"github.com/fluffy-melli/MoliDB/internal/handlers"
	"github.com/fluffy-melli/MoliDB/internal/runtime"
	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) bool {
	token := c.GetHeader("Authorization")
	if token == "" {
		handlers.SendErrResponse(c, 400, "Authorization token is required")
		return true
	}
	if token != os.Getenv("API_KEY") {
		handlers.SendErrResponse(c, 401, "Invalid authorization token")
		return true
	}
	return false
}

func Collection(c *gin.Context) {
	if i := Auth(c); i {
		return
	}
	cy, err := handlers.CryptoEncrypt(runtime.DB.GetStore())
	if err != nil {
		handlers.SendErrResponse(c, 500, err.Error())
	}
	handlers.SendResponse(c, 200, cy)
}

func CollectionID(c *gin.Context) {
	if i := Auth(c); i {
		return
	}
	id := c.Param("id")
	store, ex := runtime.DB.Get(id)
	if !ex {
		handlers.SendErrResponse(c, 404, "Collection not found")
	}
	cy, err := handlers.CryptoSingleEncrypt(store)
	if err != nil {
		handlers.SendErrResponse(c, 500, err.Error())
	}
	handlers.SendResponse(c, 200, cy)
}

func CollectionPost(c *gin.Context) {
	if i := Auth(c); i {
		return
	}
	id := c.Param("id")
	_, ex := runtime.DB.Get(id)
	if ex {
		handlers.SendErrResponse(c, 409, "Collection already exists")
	}
	runtime.DB.Set(id, map[string]any{})
	store, _ := runtime.DB.Get(id)
	cy, err := handlers.CryptoSingleEncrypt(store)
	if err != nil {
		handlers.SendErrResponse(c, 500, err.Error())
	}
	handlers.SendResponse(c, 200, cy)
}

func CollectionPut(c *gin.Context) {
	if i := Auth(c); i {
		return
	}
	id := c.Param("id")
	var data string
	if err := c.ShouldBind(&data); err != nil {
		handlers.SendErrResponse(c, 400, "Invalid request body")
		return
	}
	decryptedData, err := handlers.CryptoSingleDecrypt(data)
	if err != nil {
		handlers.SendErrResponse(c, 500, "Failed to decrypt data")
		return
	}
	runtime.DB.Set(id, decryptedData)
	store, _ := runtime.DB.Get(id)
	cy, err := handlers.CryptoSingleEncrypt(store)
	if err != nil {
		handlers.SendErrResponse(c, 500, err.Error())
	}
	handlers.SendResponse(c, 200, cy)
}

func CollectionDelete(c *gin.Context) {
	if i := Auth(c); i {
		return
	}
	id := c.Param("id")
	_, ex := runtime.DB.Get(id)
	if !ex {
		handlers.SendErrResponse(c, 404, "Collection not found")
	}
	runtime.DB.Del(id)
	cy, err := handlers.CryptoSingleEncrypt(map[string]any{})
	if err != nil {
		handlers.SendErrResponse(c, 500, err.Error())
	}
	handlers.SendResponse(c, 200, cy)
}
