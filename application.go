package httbackend

import (
	"fmt"
	"htt/httbackend/config"

	"gorm.io/driver/mysql"
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
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", config.DbUsername, config.DbPassword, config.DbHost, config.DbPort, config.DbName)
	_, err := gorm.Open(mysql.New(mysql.Config{
		DSN: connectionString,
		DefaultStringSize: 256,
		DisableDatetimePrecision: true,
		DontSupportRenameIndex: true,
		DontSupportRenameColumn: true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{})

	if err != nil {
		defer fmt.Printf("DATABASE ERROR: %v\n", err)
	}
}
