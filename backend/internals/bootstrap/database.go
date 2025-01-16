package bootstrap

import (
	"log"

	"github.com/iacopoghilardi/alarm-tool/internals/config"
	"github.com/iacopoghilardi/alarm-tool/internals/db"
)

func SetupDatabase() error {
	_, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load config: ", err)
	}

	if err := db.Connect(); err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	log.Println("Connected successfully to database")
	return nil
}
