package main

import (
	"Main.go/model"
	"Main.go/task"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
    

    db, err := gorm.Open(sqlite.Open("tasks.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    db.AutoMigrate(&model.Task{})   

    handler := task.Handler{DB: db}

    app := fiber.New()

    api := app.Group("/api")
    
    api.Get("/", handler.GetTasks)

    api.Get("/task/:id", handler.GetTaskByID)

    api.Post("/task", handler.CreateTask)

    api.Put("/task/:id", handler.UpdateTask)
    
    api.Delete("/task/:id", handler.DeleteTask)    

    app.Listen(":3000")
}