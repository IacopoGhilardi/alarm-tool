package bootstrap

import "github.com/iacopoghilardi/alarm-tool/internals/services"

type Services struct {
	UserService  *services.UserService
	AuthService  *services.AuthService
	AlarmService *services.AlarmService
}

func SetupServices(repositories *Repositories) *Services {
	return &Services{
		UserService:  services.NewUserService(repositories.UserRepository),
		AuthService:  services.NewAuthService(repositories.UserRepository),
		AlarmService: services.NewAlarmService(repositories.AlarmRepository),
	}
}
