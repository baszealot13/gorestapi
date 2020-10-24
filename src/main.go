package main

import (
	"fmt"
	"gorestapi/src/config"
	"gorestapi/src/entities"
	"gorestapi/src/routes"
	"gorestapi/src/services"

	"github.com/jinzhu/gorm"
)

var err error

func main() {
	var appConfig entities.Configuration
	services.ConfigInit(&appConfig)

	config.DB, err = gorm.Open("mysql", config.DbURL(config.BuildDBConfig(&appConfig.Database)))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer config.DB.Close()

	// config.DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&entities.User{}, &entities.About{})
	r := routes.SetupRouter()

	r.Run()

}
