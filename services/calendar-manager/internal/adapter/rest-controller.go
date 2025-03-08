package adapter

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	core "github.com/FelipeStillner/ProjectPilot/services/calendar-manager/internal/core/service"
	"github.com/gin-gonic/gin"
)

type RestController struct {
	calendarService core.CalendarService
}

func NewRestController(calendarService core.CalendarService) *RestController {
	return &RestController{calendarService: calendarService}
}

func (c *RestController) Run() error {
	gin.SetMode(gin.ReleaseMode)
	r := c.createEngine()
	port := os.Getenv("PORT_HTTP_CALENDAR_MANAGER")
	fmt.Println("Starting HTTP server on port " + port)
	return r.Run(":" + port)
}

func (c *RestController) createEngine() *gin.Engine {
	r := gin.Default()
	setRouteCreateEvent(r, c)
	setRouteReadEvent(r, c)
	setRouteUpdateEvent(r, c)
	setRouteDeleteEvent(r, c)
	return r
}

func setRouteCreateEvent(r *gin.Engine, c *RestController) {
	r.POST("/event", func(ctx *gin.Context) {
		var requestBody core.CreateEventInput
		if err := ctx.ShouldBindJSON(&requestBody); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		event, err := c.calendarService.CreateEvent(requestBody)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, event)
	})
}

func setRouteReadEvent(r *gin.Engine, c *RestController) {
	r.GET("/event/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		idUint, err := strconv.ParseUint(id, 10, 32)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		input := core.ReadEventInput{Id: uint32(idUint)}
		event, err := c.calendarService.ReadEvent(input)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, event)
	})
}

func setRouteUpdateEvent(r *gin.Engine, c *RestController) {
	r.PATCH("/event", func(ctx *gin.Context) {
		var requestBody core.UpdateEventInput
		if err := ctx.ShouldBindJSON(&requestBody); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		event, err := c.calendarService.ReadEvent(core.ReadEventInput{Id: requestBody.Id})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, event)
	})
}

func setRouteDeleteEvent(r *gin.Engine, c *RestController) {
	r.DELETE("/event/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		idUint, err := strconv.ParseUint(id, 10, 32)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		input := core.DeleteEventInput{Id: uint32(idUint)}
		err = c.calendarService.DeleteEvent(input)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"status": "event deleted"})
	})
}
