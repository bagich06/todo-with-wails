package bindings

import (
	"todo/internal/models"
	"todo/internal/services"
)

type TaskBindings struct {
	taskService *services.TaskService
}

func NewTaskBindings(taskService *services.TaskService) *TaskBindings {
	return &TaskBindings{taskService}
}

func (t *TaskBindings) Create(userID int, description string) (int, error) {
	return t.taskService.CreateTask(userID, description)
}

func (t *TaskBindings) GetAllTasks(userID int) ([]models.Task, error) {
	return t.taskService.GetAllTasks(userID)
}

func (t *TaskBindings) GetTaskByID(taskID, userID int) (models.Task, error) {
	return t.taskService.GetTaskByID(taskID, userID)
}

func (t *TaskBindings) DeleteTaskByID(taskID int, userID int) error {
	return t.taskService.DeleteTaskByID(taskID, userID)
}

func (t *TaskBindings) MarkAsDone(taskID int, userID int) error {
	return t.taskService.MarkAsDone(taskID, userID)
}

func (t *TaskBindings) MarkAsUndone(taskID int, userID int) error {
	return t.taskService.MarkAsUndone(taskID, userID)
}
