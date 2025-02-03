package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/iacopoghilardi/alarm-tool/internals/handlers"
	"github.com/iacopoghilardi/alarm-tool/internals/routes/middlewares"
)

func SetupTravelPathRoutes(r *gin.RouterGroup, handler *handlers.TravelPathHandler) {
	r.Use(middlewares.AuthMiddleware())

	r.GET("/", handler.FindAll)
	r.POST("/", handler.Create)
	r.DELETE("/:id", handler.Delete)
	r.PUT("/:id", handler.Update)
}
