package impliment

import (
	"context"

	interface17 "github.com/unknownn17/Internship_Task/internal/interface"
	"github.com/unknownn17/Internship_Task/internal/models"
)

type Service struct {
	I interface17.TaskManagement
}

func (u *Service) Register(ctx context.Context, req *models.Register_User) (string, error) {
	return u.I.Register(ctx, req)
}

func (u *Service) Verify(ctx context.Context, req *models.Verify_User) (string, error) {
	return u.I.Verify(ctx, req)
}

func (u *Service) LogIn(ctx context.Context, req *models.LogIn) (string, error) {
	return u.I.LogIn(ctx, req)
}

func (u *Service) CreateTask(ctx context.Context, req *models.Task) (*models.GetTaskResponse, error) {
	return u.I.CreateTask(ctx, req)
}

func (u *Service) GetTask(ctx context.Context, req *models.GetTaskRequest) (*models.GetTaskResponse, error) {
	return u.I.GetTask(ctx, req)
}

func (u *Service) GetTasks(ctx context.Context, req int) ([]*models.GetTaskResponse, error) {
	return u.I.GetTasks(ctx, req)
}

func (u *Service) UpdateTask(ctx context.Context, req *models.Task) (*models.GetTaskResponse, error) {
	return u.I.UpdateTask(ctx, req)
}

func (u *Service) DeleteTask(ctx context.Context, req *models.GetTaskRequest) (string, error) {
	return u.I.DeleteTask(ctx, req)
}
