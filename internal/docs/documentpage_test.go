package docs

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"
)

func TestDocumentHandler(t *testing.T) {
	// Initialize DocsBase object with mock values
	db := &DocsBase{}

	// Test creating a new task
	t.Run("CreateNewDocument", func(t *testing.T) {
		// Create a test http request with a JSON payload
		DocumentJSON := `{"regNo": "", "regDate": "", "incNo": "1", "incDate": "test task", "docType": "0", "category": "0", 
      "about": "", "endDate": "", "currencyCode": "0", "docSum": "", "authors": "", "addressee": "", "note": "", "fileList": "", "createButton": "Создать"}`
		req, err := http.NewRequest(http.MethodPost, "http://127.0.0.1:8090/docs/document/new", bytes.NewBufferString(DocumentJSON))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()

		// Call the task handler function
		handler := http.HandlerFunc(db.DocumentHandler)
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
		handler := http.HandlerFunc(db.DocumentHandler)
		handler.ServeHTTP(rr, req)

		// Confirm the response status
		if st := rr.Result().StatusCode; st != http.StatusSeeOther {
			t.Errorf("handler returned wrong status code: got %v expected %v", st, http.StatusSeeOther)
		}
	})
	t.Run("AddInfo", func(t *testing.T) {
		dd := &DocsBase{} // pass required fields
		w := httptest.NewRecorder()
		file, _ := os.Open("testdata/file.pdf") // sample file for upload
		defer file.Close()

		docForm := url.Values{}
		docForm.Add("regNo", "123")
		docForm.Add("category", "3")
		docForm.Add("docType", "2")
		docForm.Add("about", "test document")
		docForm.Add("authors", "Vova")
		docForm.Add("fileList", "") // assumes an already uploaded file

		r, _ := http.NewRequest("POST", "/docs/document/new", strings.NewReader(docForm.Encode()))
		r.Header.Add("Content-Type", "application/json")

		dd.DocumentHandler(w, r)
		assert.Equal(t, http.StatusSeeOther, w.Code)

	})
}
