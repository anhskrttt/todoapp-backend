# CRUD GoLang ToDoApp

## Techstack
* Current: 
    * HTML
    * CSS 
    * GoLang
    * gorilla/mux
    * joho/godotenv 
    * MongoDB
* Future tasks: 
    * ReactJs
    * PostGre

## Setup
- Clone project
```
git clone https://github.com/anhskrttt/todoapp-backend.git
```
- Run backend
```
cd go-server
go mod tidy
go run main.go
```

## Routes
| Routes                | HTTP Methods  | Description   |
| -------------         | ------------- | ------------- |
| /api/task             | GET           | Get all tasks |
| /api/task/{id}        | GET           | Get a task by id |
| /api/task             | POST          | Add a new task |
| /api/task/done/{id}   | PUT           | Set a task's status to done |
| /api/task/{id}        | PUT           | Set a task's status to undone |
| /api/task/{id}        | DELETE        | Delete a task by id |
| /api/task             | DELETE  | Delete all tasks |

## API Refs
