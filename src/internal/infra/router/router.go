package gin

import "github.com/gin-gonic/gin"

func NewRouter() *gin.Engine {
	router := gin.Default()
	router.Group("/api/v1/item")
	return router
}
