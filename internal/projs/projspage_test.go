package projs

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDocumentHandler(t *testing.T) {
	// Initialize DocsBase object with mock values
	pb := &ProjsBase{}

	// Test creating a new task
	t.Run("CreateNewDocument", func(t *testing.T) {
		// Create a test http request with a JSON payload
		DocumentJSON := `{"projName": "", "description": "", "createButton": "Создать"}`
		req, err := http.NewRequest(http.MethodPost, "http://127.0.0.1:8090/docs/document/new", bytes.NewBufferString(DocumentJSON))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()

		// Call the task handler function
		handler := http.HandlerFunc(pb.ProjectHandler)
		handler.ServeHTTP(rr, req)

		// Confirm the response status
		if st := rr.Result().StatusCode; st != http.StatusSeeOther {
			t.Errorf("handler returned wrong status code: got %v expected %v", st, http.StatusSeeOther)
		}
	})

	// Test creating a new document
	t.Run("UpdateDocument", func(t *testing.T) {
		// Create a test http request with a JSON payload
		DocumentJSON := `{"regNo": "", "regDate": "", "incNo": "1", "incDate": "test task", "docType": "0", "category": "0", 
      "about": "", "endDate": "", "currencyCode": "0", "docSum": "", "authors": "", "addressee": "", "note": "", "fileList": "", "createButton": "Создать"}`
		req, err := http.NewRequest(http.MethodPost, "http://127.0.0.1:8090/tasks/task/10", bytes.NewBufferString(DocumentJSON))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()

		// Call the task handler function
		handler := http.HandlerFunc(pb.ProjectHandler)
		handler.ServeHTTP(rr, req)

		// Confirm the response status
		if st := rr.Result().StatusCode; st != http.StatusSeeOther {
			t.Errorf("handler returned wrong status code: got %v expected %v", st, http.StatusSeeOther)
		}
	})
}
