package projs

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetProjAPIUnauthorized(t *testing.T) {
	// create a new request with JSON body containing project ID
	requestBody := reqProj{Proj: 3}
	reqBodyBytes, err := json.Marshal(requestBody)
	if err != nil {
		t.Errorf("Failed to create request body: %v", err)
	}
	req, err := http.NewRequest("GET", "/projects", bytes.NewBuffer(reqBodyBytes))
	if err != nil {
		t.Errorf("Failed to create request: %v", err)
	}

	// create a new ProjsBase instance with mock database and memorydb
	pb := &ProjsBase{}

	// create a new response recorder
	rr := httptest.NewRecorder()

	// call GetProjAPI handler function
	http.HandlerFunc(pb.GetProjAPI).ServeHTTP(rr, req)
}
