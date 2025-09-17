package services

import (
	"todo/internal/models"
	"todo/internal/repository"
)

type TaskService struct {
	repo repository.TaskRepositoryInterface
}

func NewTaskService(repo repository.TaskRepositoryInterface) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(userID int, description string) (int, error) {
	task := models.Task{
		UserID:      userID,
		Description: description,
		IsDone:      false,
	}

	taskID, err := s.repo.CreateTask(task, userID)
	if err != nil {
		return 0, err
	}

	return taskID, nil
}

func (s *TaskService) GetAllTasks(userID int) ([]models.Task, error) {
	tasks, err := s.repo.GetAllTasks(userID)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (s *TaskService) GetTaskByID(taskID, userID int) (models.Task, error) {
	task, err := s.repo.GetTaskByID(taskID, userID)
	if err != nil {
		return models.Task{}, err
	}

	return task, nil
}

func (s *TaskService) DeleteTaskByID(taskID int, userID int) error {
	return s.repo.DeleteTaskByID(taskID, userID)
}

func (s *TaskService) MarkAsDone(taskID int, userID int) error {
	return s.repo.MarkAsDone(taskID, userID)
}

func (s *TaskService) MarkAsUndone(taskID int, userID int) error {
	return s.repo.MarkAsUndone(taskID, userID)
}
