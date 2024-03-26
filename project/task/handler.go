package task

import (
	"Main.go/model"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Handler struct {
    DB *gorm.DB 
}

func (h *Handler) GetTaskByID(c *fiber.Ctx) error {
    var task model.Task
    id := c.Params("id")
    result := h.DB.Where("id = ?", id).First(&task)
    if result.Error != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error" : "Erro ao buscar tarefa",
        })
    } 
    return c.JSON(task)
}

func (h *Handler) GetTasks(c *fiber.Ctx) error {
    var tasks []model.Task
	result := h.DB.Find(&tasks)
	if result.Error != nil {
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
	if task.Title == "" && task.Description == "" {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "O título e a descrição não podem estar vazios",
        })
    }

	h.DB.Create(&task)
	return c.Status(fiber.StatusCreated).JSON(task)
}

func (h *Handler) UpdateTask(c *fiber.Ctx) error {
    id := c.Params("id")
    task := new(model.Task)
    if err := c.BodyParser(task); err != nil {
        return c.Status(fiber.StatusBadRequest).SendString(err.Error())
    }
    var existingTask model.Task
	result := h.DB.First(&existingTask, id)
    if result.Error != nil {
        return c.Status(404).SendString("Sem tarefas com esse ID" + id)
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
		return c.Status(404).SendString("Sem tarefas com esse ID" + id)
	}
	h.DB.Delete(&task)
	return c.SendString("Tarefa deletada com sucesso")
}