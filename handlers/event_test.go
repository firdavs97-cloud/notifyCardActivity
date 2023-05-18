package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCardActivityEventSuccess(t *testing.T) {
	var err error
	ctx := context.WithValue(context.TODO(), "MockActivityEvent", map[string][]interface{}{})
	h := Init(ctx)
	// Create a sample user payload
	body := map[string]interface{}{
		"orderType": "Purchase",
		"card":      "4433**1409",
	}
	payload, _ := json.Marshal(body)

	// Create a request with the payload
	req, err := http.NewRequest("POST", "/api/v1/events/card", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// Create a response recorder to capture the response
	rr := httptest.NewRecorder()

	// Call the handler function directly
	h.CardEvent(rr, req)

	// Check the status code
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status %d but got %d", http.StatusOK, rr.Code)
	}
}

func TestCardActivityEventError(t *testing.T) {
	ctx := context.WithValue(context.TODO(), "MockActivityEvent", map[string][]interface{}{
		"Save": {int64(0), errors.New("SaveError")},
	})
	h := Init(ctx)
	// Create a sample user payload
	body := map[string]interface{}{
		"orderType": "Purchase",
		"card":      "4433**1409",
	}
	payload, _ := json.Marshal(body)

	// Create a request with the payload
	req, err := http.NewRequest("POST", "/api/v1/events/card", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// Create a response recorder to capture the response
	rr := httptest.NewRecorder()

	// Call the handler function directly
	h.CardEvent(rr, req)

	// Check the status code
	if rr.Code != http.StatusInternalServerError {
		t.Errorf("Expected status %d but got %d", http.StatusInternalServerError, rr.Code)
	}

	// Parse the response body
	var response map[string]interface{}
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to parse response body: %v", err)
	}

	// Validate the response
	if response["message"] != "Failed to save card activity event data: SaveError" {
		t.Errorf("Unexpected message: %v", response["message"])
	}
}
