package tasks

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTaskHandler(t *testing.T) {

	tb := &TasksBase{}

	// Test creating a new task
	t.Run("CreateNewTask", func(t *testing.T) {
		// Create a test http request with a JSON payload
		taskJSON := `{"planStart": "", "planDue": "", "assignee": "1", "topic": "test task", "content": "test", "fileList": "", "createButton": "Создать"}`
		req, err := http.NewRequest(http.MethodPost, "http://127.0.0.1:8090/tasks/task/new", bytes.NewBufferString(taskJSON))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()

		// Call the task handler function
		handler := http.HandlerFunc(tb.TaskHandler)
		handler.ServeHTTP(rr, req)

		// Confirm the response status
		if st := rr.Result().StatusCode; st != http.StatusSeeOther {
			t.Errorf("handler returned wrong status code: got %v expected %v", st, http.StatusSeeOther)
		}
	})

	// Test creating a new task
	t.Run("UpdateTask", func(t *testing.T) {
		// Create a test http request with a JSON payload
		taskJSON := `{"planStart": "", "planDue": "", "assignee": "1", "topic": "update task", "content": "test", "fileList": "", "updateButton": "Сохранить"}`
		req, err := http.NewRequest(http.MethodPost, "http://127.0.0.1:8090/tasks/task/10", bytes.NewBufferString(taskJSON))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()

		// Call the task handler function
		handler := http.HandlerFunc(tb.TaskHandler)
		handler.ServeHTTP(rr, req)

		// Confirm the response status
		if st := rr.Result().StatusCode; st != http.StatusSeeOther {
			t.Errorf("handler returned wrong status code: got %v expected %v", st, http.StatusSeeOther)
		}
	})

}
