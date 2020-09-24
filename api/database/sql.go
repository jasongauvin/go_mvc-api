package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB
var err error

// DBConfig represents db configuration
type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func buildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		Host:     "localhost",
		Port:     3306,
		User:     "root",
		Password: "root",
		DBName:   "db",
	}
	return &dbConfig
}

func dbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}

func InitializeDB() error {
	// Connect Db
	DB, err := gorm.Open("mysql", dbURL(buildDBConfig()))

	if err != nil {
		fmt.Println("Status:", err)
	}
	defer DB.Close()
	return err
}
