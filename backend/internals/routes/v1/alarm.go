package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/iacopoghilardi/alarm-tool/internals/handlers"
	"github.com/iacopoghilardi/alarm-tool/internals/routes/middlewares"
)

func SetupAlarmRoutes(r *gin.RouterGroup, handler *handlers.AlarmHandler) {
	r.Use(middlewares.AuthMiddleware())

	r.GET("/", handler.FindAll)
	r.GET("/:id", handler.FindById)
	r.POST("/", handler.Create)
	r.PUT("/:id", handler.Update)
	r.DELETE("/:id", handler.Delete)
}
