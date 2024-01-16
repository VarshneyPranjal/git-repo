package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// define a struct
type Student struct {
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Age   int     `json:"age"`
	Grade float64 `json:"grade"`
}

var students []*Student

func (s *Student) displayInfo() string {
	return fmt.Sprintf("Id: %d Name: %s, Age: %d, Grade: %.2f", s.Id, s.Name, s.Age, s.Grade)
}

// handleCreate adds a new item to the students details
func handleCreate(w http.ResponseWriter, r *http.Request) {
	var newStudent Student
	// newStudent := Student{
	// 	Id:    7,
	// 	Name:  "Pranjal",
	// 	Age:   24,
	// 	Grade: 85.4,
	// }
	// if r.ContentLength == 0 {
	// 	w.Header().Set("Content-Type", "application/json")
	// 	if err := json.NewEncoder(w).Encode(newStudent); err != nil {
	// 		http.Error(w, "error encoding response", http.StatusInternalServerError)
	// 		return
	// 	}
	// 	return
	// } else {
	if err := json.NewDecoder(r.Body).Decode(&newStudent); err != nil {
		fmt.Println("error decoding request body", err)
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}
	// }
	students = append(students, &newStudent)
	w.WriteHeader(http.StatusCreated)
}

// handleGet retrieves a student by its Id
func handleGet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid student Id", http.StatusBadRequest)
		return
	}

	for _, student := range students {
		if student.Id == id {
			json.NewEncoder(w).Encode(student)
			return
		}
	}
	http.Error(w, "student not found", http.StatusNotFound)
}

// handleUpdate updates an existing item in the students details
func handleUpdate(w http.ResponseWriter, r *http.Request) {
	var updatedStudent Student
	if err := json.NewDecoder(r.Body).Decode(&updatedStudent); err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}
	for i, student := range students {
		if student.Id == updatedStudent.Id {
			students[i] = &updatedStudent
			w.WriteHeader(http.StatusOK)
			return
		}
	}
	http.Error(w, "Item not found", http.StatusNotFound)
}

// handleDelete removes an student from the students details
func handleDelete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid item Id", http.StatusBadRequest)
		return
	}
	for i, student := range students {
		if student.Id == id {
			students = append(students[:i], students[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
}

// handleGetAll retrieves all items in the student details
func handleGetAll(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(students)
}

func main() {
	//create a slice of pointers to student

	//populate the slice with student data
	students = append(students, &Student{1, "Alice", 20, 85.5})
	students = append(students, &Student{2, "bob", 22, 92.3})
	students = append(students, &Student{3, "charlie", 21, 78.9})

	//create an http server
	http.HandleFunc("/students", func(w http.ResponseWriter, r *http.Request) {
		//return student information as JSON
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(students)
	})

	//display student information
	for _, student := range students {
		student.displayInfo()
	}

	//creating handler for each operation
	http.HandleFunc("/create", handleCreate)
	http.HandleFunc("/get", handleGet)
	http.HandleFunc("/update", handleUpdate)
	http.HandleFunc("/delete", handleDelete)
	http.HandleFunc("/getall", handleGetAll)

	//start the server on port 8080
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
