package main

import (
	"assignment2/config"
	"assignment2/controllers"
	"assignment2/repository"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func GetEnvironmentVariables() *gorm.DB {
	dbUsername := os.Getenv("DB_USERNAME")
	dbHost := os.Getenv("DB_HOST")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	db := config.DBInit(dbUsername, dbPassword, dbPort, dbName, dbHost)
	return db

}
func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Loading .env file")
	}
	serverAddress := os.Getenv("SERVICE_PORT")
	db := GetEnvironmentVariables()
	repoOrder := repository.NewRepository(db)
	inDB := &controllers.InDB{
		Orderepository: repoOrder,
	}

	router := gin.Default()

	router.GET("/orders", inDB.GetOrders)
	router.POST("/orders", inDB.CreateOrder)
	router.PUT("/orders", inDB.UpdateOrder)
	router.DELETE("/orders/:id", inDB.DeleteOrder)
	router.Run(serverAddress)
}
