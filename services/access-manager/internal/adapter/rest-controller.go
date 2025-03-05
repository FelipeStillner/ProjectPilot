package adapter

import (
	"fmt"
	"net/http"
	"os"

	core "github.com/FelipeStillner/ProjectPilot/services/access-manager/internal/core/service"
	"github.com/gin-gonic/gin"
)

type RestController struct {
	accessManagerService core.AccessService
}

func NewRestController(accessManagerService core.AccessService) *RestController {
	return &RestController{accessManagerService}
}

func (c *RestController) Run() error {
	gin.SetMode(gin.ReleaseMode)
	r := c.createEngine()
	port := os.Getenv("PORT_HTTP_ACCESS_MANAGER")
	fmt.Println("Starting HTTP server on port " + port)
	return r.Run(":" + port)
}

func (c *RestController) createEngine() *gin.Engine {
	r := gin.Default()
	setRouteCreateTeam(r, c)
	setRouteCreateUser(r, c)
	setRouteLogin(r, c)
	return r
}

func setRouteCreateTeam(r *gin.Engine, c *RestController) {
	r.POST("/team", func(ctx *gin.Context) {
		var requestBody core.CreateTeamInput
		if err := ctx.ShouldBindJSON(&requestBody); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		task, err := c.accessManagerService.CreateTeam(requestBody)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, task)
	})
}

func setRouteCreateUser(r *gin.Engine, c *RestController) {
	r.POST("/user", func(ctx *gin.Context) {
		var requestBody core.CreateUserInput
		if err := ctx.ShouldBindJSON(&requestBody); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		task, err := c.accessManagerService.CreateUser(requestBody)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, task)
	})
}

func setRouteLogin(r *gin.Engine, c *RestController) {
	r.POST("/login", func(ctx *gin.Context) {
		var requestBody core.LoginInput
		if err := ctx.ShouldBindJSON(&requestBody); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		task, err := c.accessManagerService.Login(requestBody)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, task)
	})
}
