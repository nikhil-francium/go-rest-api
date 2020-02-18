package main

import (
	"bytes"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

//GithubIssueLabelNotifierToSlack handler for /payload POST from Github for issues event
func GithubIssueLabelNotifierToSlack(c *gin.Context) {
	result := decodeGithubPayloadMessage(c.Request)
	if result != nil {
		finalResult := constructSlackResponseMessage(result)
		if finalResult != nil {
			url := os.Getenv("SLACK_WEBHOOK_URL")
			http.Post(url, "application/json", bytes.NewBuffer(finalResult))
		}
	}
}
