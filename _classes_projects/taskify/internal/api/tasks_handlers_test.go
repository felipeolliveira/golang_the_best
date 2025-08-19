package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandleCreateTask(t *testing.T) {
	api := Application{
		// TaskService: ,
	}

	payload := map[string]any{
		"title":       "Task",
		"description": "A testing for task",
		"priority":    8000,
	}

	body, err := json.Marshal(payload)
	assert.NoError(t, err)

	req := httptest.NewRequest("Post", "/api/v1/tasks", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(api.handleCreateTask)
	handler.ServeHTTP(rec, req)

	// t.Logf("Rec body %s\n", rec.Body.Bytes())
	assert.Equal(t, rec.Code, http.StatusCreated)

	var responseBody map[string]any
	err = json.Unmarshal(rec.Body.Bytes(), &responseBody)
	assert.NoError(t, err)
	assert.Equal(t, responseBody["title"], payload["title"])
}
