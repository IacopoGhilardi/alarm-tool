package bootstrap

import (
	"github.com/iacopoghilardi/alarm-tool/internals/handlers"
)

type Handlers struct {
	UserHandler       *handlers.UserHandler
	AuthHandler       *handlers.AuthHandler
	AlarmHandler      *handlers.AlarmHandler
	TravelPathHandler *handlers.TravelPathHandler
}

func SetupHandlers(services *Services) *Handlers {
	return &Handlers{
		UserHandler:       handlers.NewUserHandler(services.UserService),
		AuthHandler:       handlers.NewAuthHandler(services.AuthService),
		AlarmHandler:      handlers.NewAlarmHandler(services.AlarmService),
		TravelPathHandler: handlers.NewTravelPathHandler(services.TravelPathService),
	}
}
