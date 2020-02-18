package main

import (
	"bytes"
	"encoding/json"
	"testing"
)

//TestSlackMessageGeneration testcase for Slack Message Generation
func TestSlackMessageGeneration(t *testing.T) {
	requiredJSON := map[string]interface{}{
		"blocks": []*message{
			&message{
				TextType: "section",
				TextData: map[string]string{
					"type": "mrkdwn",
					"text": "Label *testing* was added to the issue <https://google.com|201>",
				},
			},
		},
	}
	expectedResult, _ := json.Marshal(requiredJSON)
	actualResult := generateSlackMessage("https://google.com", 201, "testing", "added")

	res := bytes.Compare(actualResult, expectedResult)
	if res != 0 {
		t.Errorf("Mismatch")
	}
}

func TestSlackResponseConstruction(t *testing.T) {
	requiredJSON := map[string]interface{}{
		"blocks": []*message{
			&message{
				TextType: "section",
				TextData: map[string]string{
					"type": "mrkdwn",
					"text": "Label *testing* was added to the issue <https://google.com|201>",
				},
			},
		},
	}
	expectedResult, _ := json.Marshal(requiredJSON)

	resultParam := map[string]interface{}{
		"action": "labeled",
		"issue": map[string]interface{}{
			"html_url": "https://google.com",
			"number":   float64(201),
		},
		"label": map[string]interface{}{
			"name": "testing",
		},
	}

	actualResult := constructSlackResponseMessage(resultParam)

	res := bytes.Compare(actualResult, expectedResult)
	if res != 0 {
		t.Errorf("Mismatch")
	}

}

func TestNilForSlackResponse(t *testing.T) {
	resultParam := map[string]interface{}{
		"action": "edited",
	}
	result := constructSlackResponseMessage(resultParam)

	if result != nil {
		t.Errorf("Mismatch")
	}
}

func TestActionMap(t *testing.T) {
	actionType := getActionType("labeled")
	if actionType != "added" {
		t.Errorf("Mismatch")
	}
}
