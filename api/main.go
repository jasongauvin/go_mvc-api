package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jasongauvin/go_mvc-api/api/database"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jasongauvin/go_mvc-api/api/routes"
	"github.com/joho/godotenv"
)

func main() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	// Connected DB
	database.InitializeDB(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	// Server start
	fmt.Println("\n Start serving ...")
	r := gin.Default()
	routes.InitializeRouter(r)
	r.Run()

}
