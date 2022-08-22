package router

import (
	"go-server/middleware"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	middleware.InitDatabase()

	r := mux.NewRouter()

	// Route Handlers / Endpoints
	// Test API.
	r.HandleFunc("/", middleware.HomePageHandler).Methods("GET")
	// My APIs.
	// Get all tasks.
	// Future task: logic for "/api/dodo?id="
	r.HandleFunc("/api/task", middleware.GetAllTasks).Methods("GET")
	// Get task info by id.
	r.HandleFunc("/api/task/{id}", middleware.GetTaskById).Methods("GET")
	// Future tasks.
	// Get all done tasks.
	// Get all undone tasks.
	// Get all tasks from startDate to endDate.

	// Add a new task.
	r.HandleFunc("/api/task", middleware.AddTask).Methods("POST")

	// Set task's status to done.
	r.HandleFunc("/api/task/done/{id}", middleware.SetDone).Methods("PUT")
	// Set task's status to undone.
	r.HandleFunc("/api/task/undone/{id}", middleware.SetUndone).Methods("PUT")
	// Future tasks.
	// Modify task's name.
	// Modify task's description.

	// Delete a task by id.
	r.HandleFunc("/api/task/{id}", middleware.DeleteTaskById).Methods("DELETE")
	// Delete all tasks.
	r.HandleFunc("/api/task", middleware.DeleteAllTasks).Methods("DELETE")

	return r
}
