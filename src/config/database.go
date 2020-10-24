package config

import (
	"fmt"
	"gorestapi/src/entities"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

// DBConfig represents db configuration
type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func BuildDBConfig(db *entities.DatabaseConfiguration) *DBConfig {
	fmt.Println("database.BuildDBConfig.db ", db)
	dbConfig := DBConfig{
		Host:     db.Host,
		Port:     db.Port,
		User:     db.User,
		Password: db.Password,
		DBName:   db.Dbname,
	}
	return &dbConfig
}
func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}
