package team

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestProfileHandler_UpdateProfile_Success(t *testing.T) {
	// Create a new instance of the TeamBase struct
	tb := &TeamBase{}

	// Create a new request with POST method and set the form values
	req, err := http.NewRequest("POST", "/profile/1", strings.NewReader("updateButton=Update&firstName=John&surname=Doe&email=john.doe@example.com"))
	if err != nil {
		t.Fatal(err)
	}

	// Set the Content-Type header to simulate form data
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Create a new ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Call the ProfileHandler function with the mock database and the new request and response recorders
	http.HandlerFunc(tb.ProfileHandler).ServeHTTP(rr, req)

	// Check if the response status code is 200 OK
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check if the response body contains the expected message
	expected := "Profile updated successfully"
	if !strings.Contains(rr.Body.String(), expected) {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
