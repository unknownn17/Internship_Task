package interface17

import (
	"context"

	"github.com/unknownn17/Internship_Task/internal/models"
)

type TaskManagement interface {
	Register(ctx context.Context, req *models.Register_User) (string, error)
	Verify(ctx context.Context, req *models.Verify_User) (string, error)
	LogIn(ctx context.Context, req *models.LogIn) (string, error)
	CreateTask(ctx context.Context, req *models.Task) (*models.GetTaskResponse, error)
	GetTask(ctx context.Context, req *models.GetTaskRequest) (*models.GetTaskResponse, error)
	GetTasks(ctx context.Context, req int) ([]*models.GetTaskResponse, error)
	UpdateTask(ctx context.Context, req *models.Task) (*models.GetTaskResponse, error)
	DeleteTask(ctx context.Context, req *models.GetTaskRequest) (string, error)
}

