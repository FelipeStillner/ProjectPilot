package port

import "github.com/gin-gonic/gin"

type RestControllerInterface interface {
	SetRoutes() *gin.Engine
}
