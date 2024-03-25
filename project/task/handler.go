package task

import (
	"fmt"

	"Main.go/model"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Handler struct {
    DB *gorm.DB
}

func (h *Handler) GetTasks(c *fiber.Ctx) error {
    var tasks []model.Task
	err := h.DB.Find(&tasks)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error" : " Erro ao buscar tarefas",
		})
	} 
    return c.JSON(tasks)
}

func (h *Handler) CreateTask(c *fiber.Ctx) error{
	task := new(model.Task)
	if err := c.BodyParser(task); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Dados da tarefa inválidos",
		})	
	}
	if task.Title == "" || task.Description == "" {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "O título e a descrição não podem estar vazios",
        })
    }

	h.DB.Create(&task)
	return c.Status(fiber.StatusCreated).JSON(task)
}

func (h *Handler) UpdateTask(c *fiber.Ctx) error {
    id := c.Params("id")
	fmt.Println("ID: " , id)
    task := new(model.Task)
    if err := c.BodyParser(task); err != nil {
        return c.Status(400).SendString(err.Error())
    }
    var existingTask model.Task
	result := h.DB.First(&existingTask, id)
    if result.Error != nil {
        return c.Status(404).SendString("No task found with ID " + id)
    }

    existingTask.Title = task.Title
    existingTask.Description = task.Description
    h.DB.Save(&existingTask)
    return c.JSON(existingTask)
}


func (h *Handler) DeleteTask(c *fiber.Ctx) error{
	id := c.Params("id")
	var task model.Task
	result := h.DB.First(&task, id)
	if result.Error != nil {
		return c.Status(404).SendString("No task found with ID " + id)
	}
	h.DB.Delete(&task)
	return c.SendString("Task successfully deleted")
}
