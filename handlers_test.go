package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPayLoadHandler(t *testing.T) {

	router := setUpRouter()

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

	expectedResult, _ := json.Marshal(resultParam)
	req, _ := http.NewRequest("POST", "/payload", bytes.NewBuffer(expectedResult))
	req.Header.Set("X-GitHub-Event", "issues")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}
