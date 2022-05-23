package httbackend

import (
	"fmt"
	"htt/httbackend/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ApplicationHandler struct {
	Config *config.Configuration
	DB     *gorm.DB
}

func (a *ApplicationHandler) PrepareConfig() {
	a.Config = config.New()
}

func (a *ApplicationHandler) InitDB() {
	config := a.Config.Database
	connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", config.DbHost, config.DbUsername, config.DbPassword, config.DbName, config.DbPort)
	_, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	if err != nil {
		defer fmt.Printf("DATABASE ERROR: %v\n", err)
	}
}
