package testinit

import (
	"github.com/degalaxy/gp-common/log"
	"github.com/degalaxy/gp-common/repository"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DBConnection *repository.DBConnection

// Init resource all test or testing code is needed
func init() {
	log.InitLogger(log.LogConfig{Name: "WgxFoundationTestLogger", Level: "info", FormatJson: false})
	db, err := gorm.Open(sqlite.Open("wgxFoundationTest.db"), &gorm.Config{})
	if err != nil {
		log.Logger.Fatal("failed to connect database")
	}
	DBConnection = new(repository.DBConnection)
	DBConnection.DB = db
}
