package middleware

import (
	"encoding/json"
	"fmt"
	"go-server/models"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var tasks = models.Tasks

func Init() {
	tasks = append(tasks, models.Task{
		ID:          "1",
		Name:        "Learn RestfulAPI",
		Description: "Do freecodecamp course: RestfulAPI",
		Status:      true,
	}, models.Task{
		ID:          "2",
		Name:        "Learn Security",
		Description: "Do freecodecamp course: Security",
		Status:      false,
	})
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
	json.NewEncoder(w).Encode((&models.Task{}))
}

func AddTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var taskSample models.Task
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
	json.NewEncoder(w).Encode((&models.Task{}))
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
	json.NewEncoder(w).Encode((&models.Task{}))
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
	json.NewEncoder(w).Encode(&models.Task{})
}

func DeleteAllTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	tasks = tasks[:0]
	json.NewEncoder(w).Encode(tasks)

}
