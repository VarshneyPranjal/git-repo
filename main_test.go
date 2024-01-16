package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestHandleCreate(t *testing.T) {
	clearStudents()
	payload := []byte(`{"id": 4, "name": "TestStudent", "age": 18, "grade": 90.5}`)

	req, err := http.NewRequest("POST", "/create", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handleCreate)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handleCreate returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	// Verify that the new student is added
	if len(students) != 1 {
		t.Errorf("Expected 1 student, got %d", len(students))
	}
}
func TestHandleGet(t *testing.T) {
	clearStudents()
	students = append(students, &Student{4, "TestStudent", 18, 90.5})

	req, err := http.NewRequest("GET", "/get?id=4", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handleGet)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handleGet returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Verify the response body contains the expected student data
	expectedResponse := `{"id":4,"name":"TestStudent","age":18,"grade":90.5}`
	if rr.Body.String() != expectedResponse {
		t.Errorf("handleGet returned unexpected body: got %v want %v", rr.Body.String(), expectedResponse)
	}
}

// Similar testing functions can be added for handleUpdate and handleDelete.

func clearStudents() {
	students = nil
}

func TestMain(m *testing.M) {
	// Run tests
	code := m.Run()

	// Clean up after all tests
	clearStudents()

	// Exit with the test result code
	os.Exit(code)
}
