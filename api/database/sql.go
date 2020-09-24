package database

import (
	"fmt"
	"log"

	"github.com/jasongauvin/go_mvc-api/api/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func InitializeDB(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {

	var err error

	// Connect Db
	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
	DB, err = gorm.Open(Dbdriver, DBURL)
	if err != nil {
		fmt.Printf("Cannot connect to %s database", Dbdriver)
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("We are connected to the %s database", Dbdriver)
	}

	DB.AutoMigrate(&models.Activity{})
	fmt.Println("\n Migrations are OK ...")
}
