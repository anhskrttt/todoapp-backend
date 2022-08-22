package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Structs (Models)
// type User struct {
// 	ID       int       `json: "id"`
// 	Username string    `json: "username"`
// 	Tasks    *TaskList `json: "tasks"`
// }

// type TaskList struct {
// 	ID    int    `json: "id"`
// 	Name  string `json: "name"`
// 	Tasks *Task  `json:"tasks"`
// }

type Task struct {
	ID          string `json: "id"`
	Name        string `json: "name"`
	Description string `json: "description"`
	Status      bool   `json: "status"`
}

// Future task: Connect to DB
var tasks []Task

func main() {
	// Init Router.
	r := mux.NewRouter()

	// Mock data.
	tasks = append(tasks, Task{
		ID:          "1",
		Name:        "Learn RestfulAPI",
		Description: "Do freecodecamp course: RestfulAPI",
		Status:      true,
	}, Task{
		ID:          "2",
		Name:        "Learn Security",
		Description: "Do freecodecamp course: Security",
		Status:      false,
	})

	// Route Handlers / Endpoints
	// Test API.
	r.HandleFunc("/", HomePageHandler).Methods("GET")
	// My APIs.
	// Get all tasks.
	// Future task: logic for "/api/dodo?id="
	r.HandleFunc("/api/task", GetAllTasks).Methods("GET")
	// Get task info by id.
	r.HandleFunc("/api/task/{id}", GetTaskById).Methods("GET")
	// Future tasks.
	// Get all done tasks.
	// Get all undone tasks.
	// Get all tasks from startDate to endDate.

	// Add a new task.
	r.HandleFunc("/api/task", AddTask).Methods("POST")

	// Set task's status to done.
	r.HandleFunc("/api/task/done/{id}", SetDone).Methods("PUT")
	// Set task's status to undone.
	r.HandleFunc("/api/task/undone/{id}", SetUndone).Methods("PUT")
	// Future tasks.
	// Modify task's name.
	// Modify task's description.

	// Delete a task by id.
	r.HandleFunc("/api/task/{id}", DeleteTaskById).Methods("DELETE")
	// Delete all tasks.
	r.HandleFunc("/api/task", DeleteAllTasks).Methods("DELETE")

	// Start server.
	fmt.Println("Listening and serving.....")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	welcome := "Welcome to my homepage"
	json.NewEncoder(w).Encode(&welcome)
}

func GetAllTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func GetTaskById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Get params
	params := mux.Vars(r)
	fmt.Println(params)

	for _, t := range tasks {
		if params["id"] == t.ID {
			json.NewEncoder(w).Encode(t)
			return
		}
	}
	// Future work: Return error invalid id.
	json.NewEncoder(w).Encode((&Task{}))
}

func AddTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var taskSample Task
	// Why "_" ?
	_ = json.NewDecoder(r.Body).Decode(&taskSample)
	taskSample.ID = strconv.Itoa((rand.Intn(1000000)))
	tasks = append(tasks, taskSample)
	json.NewEncoder(w).Encode(taskSample)
}

func SetDone(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, t := range tasks {
		if params["id"] == t.ID {
			tasks[i].Status = true
			json.NewEncoder(w).Encode(tasks[i])
			return
		}
	}
	json.NewEncoder(w).Encode((&Task{}))
}

func SetUndone(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, t := range tasks {
		if params["id"] == t.ID {
			tasks[i].Status = false
			json.NewEncoder(w).Encode(tasks[i])
			return
		}
	}
	json.NewEncoder(w).Encode((&Task{}))
}

func DeleteTaskById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, t := range tasks {
		if params["id"] == t.ID {
			// Delete this task.
			tasks = append(tasks[:i], tasks[i+1:]...)
			// Return all tasks.
			json.NewEncoder(w).Encode(tasks)
			return
		}
	}
	// Future task: Return error invalid
	json.NewEncoder(w).Encode(&Task{})
}

func DeleteAllTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	tasks = tasks[:0]
	json.NewEncoder(w).Encode(tasks)

}
