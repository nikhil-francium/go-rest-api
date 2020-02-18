package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func constructSlackResponseMessage(result map[string]interface{}) []byte {
	action := result["action"].(string)
	actionType, exists := getActionMap()[action]
	if !exists {
		return nil
	}
	issueDetails := result["issue"].(map[string]interface{})
	labelDetails := result["label"].(map[string]interface{})
	issueURL := issueDetails["html_url"].(string)
	issueNumber := issueDetails["number"].(float64)
	labelName := labelDetails["name"].(string)
	return generateSlackMessage(issueURL, issueNumber, labelName, actionType)
}

func generateSlackMessage(issueURL string, issueNumber float64, labelName string, actionType string) []byte {
	slackMessage := &message{
		TextType: "section",
		TextData: map[string]string{
			"type": "mrkdwn",
			"text": fmt.Sprintf("Label *%s* was %s to the issue <%s|%s>", labelName, actionType, issueURL, strconv.FormatFloat(issueNumber, 'f', -1, 64)),
		},
	}
	slackPostBody := map[string]interface{}{
		"blocks": []*message{slackMessage},
	}
	finalResult, _ := json.Marshal(slackPostBody)
	return finalResult
}

func decodeGithubPayloadMessage(request *http.Request) map[string]interface{} {
	var result map[string]interface{}
	eventType := request.Header.Get("X-GitHub-Event")
	if eventType == "issues" {
		body, _ := ioutil.ReadAll(request.Body)
		json.Unmarshal(body, &result)
	}
	return result
}

func getActionMap() map[string]string {
	actionsMap := map[string]string{
		"labeled":   "added",
		"unlabeled": "removed",
	}
	return actionsMap
}

func getActionType(action string) string {
	return getActionMap()[action]
}
