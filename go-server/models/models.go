package models

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
var Tasks []Task