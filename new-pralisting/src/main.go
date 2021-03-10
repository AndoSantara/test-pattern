package main

import (
	"log"
	"os"

	"new-pralisting/src/config"

	_entity "new-pralisting/src/api/entity"
	_handler "new-pralisting/src/api/handler"
	_repo "new-pralisting/src/api/repo"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error getting env")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "1111"
	}

	config.Init()
	db := config.GetDB()

	repoPralisting := _repo.NewPralistingRepository(db)
	entityPralisting := _entity.NewPralistingEntity(repoPralisting)

	r := gin.Default()
	api := r.Group("/v1")

	_handler.NewPralistingHandler(api, entityPralisting)

	r.Run(":" + port)
}
