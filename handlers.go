package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CRUD operations on database
func loadTasks() ([]Task, error) {
	var tasks []Task
	rows, err := db.Query("SELECT * FROM tasks")
	if err != nil {
		return nil, fmt.Errorf("loading tasks: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var task Task
		err := rows.Scan(&task.ID, &task.Title, &task.Completed)
		if err != nil {
			return nil, fmt.Errorf("scanning task: %v", err)
		}
		tasks = append(tasks, task)
	}

	// Check for errors from iterating over rows
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterating tasks: %v", err)
	}
	return tasks, nil
}

func getTasks(c *gin.Context) {
	tasks, err := loadTasks()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, tasks)
}

func storeTask(task Task) (int64, error) {
	result, err := db.Exec("INSERT INTO tasks (title, completed) VALUES (?, ?)", task.Title, task.Completed)
	if err != nil {
		return 0, fmt.Errorf("addTask: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addTask: %v", err)
	}
	return id, nil
}
func postTask(c *gin.Context) {
	var newTask Task

	if err := c.BindJSON(&newTask); err != nil {
		return
	}

	storeTask(newTask)
	c.IndentedJSON(http.StatusCreated, newTask)
}
