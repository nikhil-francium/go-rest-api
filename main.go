package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type message struct {
	TextType string            `json:"type"`
	TextData map[string]string `json:"text"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	router := setUpRouter()
	router.Run(":4567")
}

func setUpRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/payload", GithubIssueLabelNotifierToSlack)
	return router
}
