package bootstrap

import (
	"github.com/iacopoghilardi/alarm-tool/internals/db"
	"github.com/iacopoghilardi/alarm-tool/internals/repositories"
)

type Repositories struct {
	UserRepository       *repositories.UserRepository
	AlarmRepository      *repositories.AlarmRepository
	TravelPathRepository *repositories.TravelPathRepository
}

func SetupRepositories() *Repositories {
	return &Repositories{
		UserRepository:       repositories.NewUserRepository(db.GetDB()),
		AlarmRepository:      repositories.NewAlarmRepository(db.GetDB()),
		TravelPathRepository: repositories.NewTravelPathRepository(db.GetDB()),
	}
}
