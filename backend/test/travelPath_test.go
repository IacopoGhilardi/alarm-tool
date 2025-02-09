package test

import (
	"context"
	"testing"

	"github.com/iacopoghilardi/alarm-tool/internals/repositories"
	"github.com/iacopoghilardi/alarm-tool/internals/services"
	"github.com/iacopoghilardi/alarm-tool/internals/types/dto"
	"github.com/iacopoghilardi/alarm-tool/test/producers"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type TravelPathTestSuite struct {
	suite.Suite
	db          *gorm.DB
	pgContainer *postgres.PostgresContainer
	service     *services.TravelPathService
	ctx         context.Context
}

func (s *TravelPathTestSuite) SetupSuite() {
	s.ctx = context.Background()

	pgContainer, db, err := producers.SetupGenericSuite(s.ctx)
	if err != nil {
		s.T().Fatalf("Failed to setup generic suite: %v", err)
	}
	s.pgContainer = pgContainer
	s.db = db
	s.service = services.NewTravelPathService(repositories.NewTravelPathRepository(s.db))
}

func (s *TravelPathTestSuite) TearDownSuite() {
	producers.TearDownGenericSuite(s.ctx, s.db, s.pgContainer)
}

func (s *TravelPathTestSuite) TestCreateTravelPath() {
	user := producers.CreateTestUser(s.db)

	startLocation := dto.LocationDto{
		Latitude:  1.0,
		Longitude: 1.0,
	}

	endLocation := dto.LocationDto{
		Latitude:  2.0,
		Longitude: 2.0,
	}

	createTravelPathDto := dto.CreateTravelPathDto{
		UserID:        user.ID,
		StartLocation: startLocation,
		EndLocation:   endLocation,
	}

	travelPath, err := s.service.Create(&createTravelPathDto)
	if err != nil {
		s.T().Fatalf("Failed to create travel path: %v", err)
	}

	assert.NotNil(s.T(), travelPath)
	assert.Equal(s.T(), travelPath.StartLocation, startLocation)
	assert.Equal(s.T(), travelPath.EndLocation, endLocation)
}

func TestTravelPathTestSuite(t *testing.T) {
	suite.Run(t, new(TravelPathTestSuite))
}
