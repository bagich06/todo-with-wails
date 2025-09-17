package main

import (
	"context"
	"log"
	"todo/internal/bindings"
	"todo/internal/models"
	"todo/internal/repository"
	"todo/internal/services"
)

type App struct {
	ctx          context.Context
	authBinding  *bindings.AuthBinding
	taskBindings *bindings.TaskBindings
}

func NewApp() *App {
	return &App{}
}

func (a *App) OnStartup(ctx context.Context) {
	a.ctx = ctx

	db, err := repository.NewPGRepo("postgres://postgres:postgres@localhost:5432/intern_test")
	if err != nil {
		log.Fatal(err)
	}

	jwtService := services.NewJWTService()
	userService := services.NewUserService(db, jwtService)
	taskService := services.NewTaskService(db)

	a.authBinding = bindings.NewAuthBinding(userService)
	a.taskBindings = bindings.NewTaskBindings(taskService)
}

func (a *App) Login(email, password string) (*models.AuthResponse, error) {
	return a.authBinding.Login(email, password)
}

func (a *App) Register(username, email, password string) (*models.AuthResponse, error) {
	return a.authBinding.Register(username, email, password)
}

func (a *App) Create(userID int, description string) (int, error) {
	return a.taskBindings.Create(userID, description)
}

func (a *App) GetAllTasks(userID int) ([]models.Task, error) {
	return a.taskBindings.GetAllTasks(userID)
}

func (a *App) GetTaskByID(taskID, userID int) (models.Task, error) {
	return a.taskBindings.GetTaskByID(taskID, userID)
}

func (a *App) DeleteTaskByID(taskID, userID int) error {
	return a.taskBindings.DeleteTaskByID(taskID, userID)
}

func (a *App) MarkAsDone(taskID int, userID int) error {
	return a.taskBindings.MarkAsDone(taskID, userID)
}

func (a *App) MarkAsUndone(taskID int, userID int) error {
	return a.taskBindings.MarkAsUndone(taskID, userID)
}
