package middleware

import (
	"encoding/json"
	"os"

	"github.com/fluffy-melli/MoliDB/internal/handlers"
	"github.com/fluffy-melli/MoliDB/internal/runtime"
	"github.com/fluffy-melli/MoliDB/pkg/logger"
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

// Collection godoc
// @Summary Get all collections
// @Description
// Client sends a request with data that is first compressed using gzip and then AES encrypted.
// The server will decrypt the data using AES and then decompress it using gzip.
// The server sends back the data, which is first AES encrypted and then compressed using gzip.
// Both request and response apply AES encryption and gzip compression/decompression.
// @Tags collections
// @Param Authorization header string true "Token"
// @Success 200 {object} handlers.Response "Encrypted collections data (AES encrypted, gzip compressed)"
// @Failure 500 {object} handlers.ErrResponse "Internal server error"
// @Router /collection [get]
func Collection(c *gin.Context) {
	if i := Auth(c); i {
		return
	}
	jsonData, err := json.Marshal(runtime.DB.GetStore())
	if err != nil {
		logger.WARNING("%v", err)
		handlers.SendErrResponse(c, 500, err.Error())
		return
	}
	cy, err := handlers.CryptoEncrypt(jsonData)
	if err != nil {
		logger.WARNING("%v", err)
		handlers.SendErrResponse(c, 500, err.Error())
		return
	}
	handlers.SendResponse(c, 200, cy)
}

// CollectionID godoc
// @Summary Get collection by ID
// @Description
// Client sends a request with data that is first compressed using gzip and then AES encrypted.
// The server will decrypt the data using AES and then decompress it using gzip.
// The server sends back the data, which is first AES encrypted and then compressed using gzip.
// Both request and response apply AES encryption and gzip compression/decompression.
// @Tags collections
// @Param Authorization header string true "Token"
// @Param id path string true "Collection ID"
// @Success 200 {object} handlers.Response "Encrypted collection data (AES encrypted, gzip compressed)"
// @Failure 400 {object} handlers.ErrResponse "Collection not found"
// @Failure 500 {object} handlers.ErrResponse "Internal server error"
// @Router /collection/{id} [get]
func CollectionID(c *gin.Context) {
	if i := Auth(c); i {
		return
	}
	id := c.Param("id")
	store, ex := runtime.DB.Get(id)
	if !ex {
		handlers.SendErrResponse(c, 400, "Collection not found")
		return
	}
	jsonData, err := json.Marshal(store)
	if err != nil {
		logger.WARNING("%v", err)
		handlers.SendErrResponse(c, 500, err.Error())
		return
	}
	cy, err := handlers.CryptoEncrypt(jsonData)
	if err != nil {
		logger.WARNING("%v", err)
		handlers.SendErrResponse(c, 500, err.Error())
		return
	}
	handlers.SendResponse(c, 200, cy)
}

// CollectionPut godoc
// @Summary Update collection by ID
// @Description
// Client sends a request with data that is first compressed using gzip and then AES encrypted.
// The server will decrypt the data using AES and then decompress it using gzip.
// The server sends back the updated data, which is first AES encrypted and then compressed using gzip.
// Both request and response apply AES encryption and gzip compression/decompression.
// @Tags collections
// @Param Authorization header string true "Token"
// @Param id path string true "Collection ID"
// @Param body header string true "Encrypted collection data (AES encrypted, gzip compressed)"
// @Success 200 {object} handlers.Response "Encrypted updated collection data (AES encrypted, gzip compressed)"
// @Failure 400 {object} handlers.ErrResponse "Collection not found"
// @Failure 500 {object} handlers.ErrResponse "Internal server error"
// @Router /collection/{id} [put]
func CollectionPut(c *gin.Context) {
	if i := Auth(c); i {
		return
	}
	id := c.Param("id")
	data := c.GetHeader("body")
	decryptedData, err := handlers.CryptoDecrypt(data)
	if err != nil {
		logger.WARNING("%v", err)
		handlers.SendErrResponse(c, 500, "Failed to decrypt data")
		return
	}
	var mapdata any
	err = json.Unmarshal(decryptedData, &mapdata)
	if err != nil {
		logger.WARNING("%v", err)
		handlers.SendErrResponse(c, 500, "Failed to decrypt data")
		return
	}
	runtime.DB.Set(id, mapdata)
	store, _ := runtime.DB.Get(id)
	jsonData, err := json.Marshal(store)
	if err != nil {
		logger.WARNING("%v", err)
		handlers.SendErrResponse(c, 500, err.Error())
		return
	}
	cy, err := handlers.CryptoEncrypt(jsonData)
	if err != nil {
		logger.WARNING("%v", err)
		handlers.SendErrResponse(c, 500, err.Error())
		return
	}
	handlers.SendResponse(c, 200, cy)
}

// CollectionDelete godoc
// @Summary Delete collection by ID
// @Description
// Client sends a request with data that is first compressed using gzip and then AES encrypted.
// The server will decrypt the data using AES and then decompress it using gzip.
// The server sends back a confirmation, which is first AES encrypted and then compressed using gzip.
// Both request and response apply AES encryption and gzip compression/decompression.
// @Tags collections
// @Param Authorization header string true "Token"
// @Param id path string true "Collection ID"
// @Success 200 {object} handlers.Response "Collection deleted (AES encrypted, gzip compressed)"
// @Failure 400 {object} handlers.ErrResponse "Collection not found"
// @Failure 500 {object} handlers.ErrResponse "Internal server error"
// @Router /collection/{id} [delete]
func CollectionDelete(c *gin.Context) {
	if i := Auth(c); i {
		return
	}
	id := c.Param("id")
	_, ex := runtime.DB.Get(id)
	if !ex {
		handlers.SendErrResponse(c, 400, "Collection not found")
		return
	}
	runtime.DB.Del(id)
	handlers.SendResponse(c, 200, "")
}

//func CollectionPost(c *gin.Context) {
//	if i := Auth(c); i {
//		return
//	}
//	id := c.Param("id")
//	_, ex := runtime.DB.Get(id)
//	if ex {
//		handlers.SendErrResponse(c, 409, "Collection already exists")
//		return
//	}
//	data := c.GetHeader("body")
//	decryptedData, err := handlers.CryptoSingleDecrypt(data)
//	if err != nil {
//		handlers.SendErrResponse(c, 500, "Failed to decrypt data")
//		return
//	}
//	runtime.DB.Set(id, decryptedData)
//	store, _ := runtime.DB.Get(id)
//	cy, err := handlers.CryptoSingleEncrypt(store)
//	if err != nil {
//		handlers.SendErrResponse(c, 500, err.Error())
//		return
//	}
//	handlers.SendResponse(c, 200, cy)
//}
