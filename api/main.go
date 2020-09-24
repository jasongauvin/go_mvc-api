package main

import (
	"fmt"

	"github.com/jasongauvin/go_mvc-api/api/database"
	"github.com/jasongauvin/go_mvc-api/api/models"

	"github.com/jasongauvin/go_mvc-api/api/routes"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Connected DB
	database.InitializeDB()
	database.DB.AutoMigrate(&models.Activity{})

	fmt.Println("\n Connected to DB and migrations OK ...")
	// Server start
	fmt.Println("\n Start serving ...")
	r := gin.Default()
	routes.InitializeRouter(r)
	r.Run()

}
