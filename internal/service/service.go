package services

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	email1 "github.com/unknownn17/Internship_Task/internal/auth/email"
	jwttoken "github.com/unknownn17/Internship_Task/internal/auth/jwt"
	redisserver "github.com/unknownn17/Internship_Task/internal/auth/redis"
	mongodb "github.com/unknownn17/Internship_Task/internal/database/mongoDB"
	"github.com/unknownn17/Internship_Task/internal/database/sqlc/storage"
	"github.com/unknownn17/Internship_Task/internal/models"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	Storage *storage.Queries
	Mongo   *mongodb.Mongo
	Redis   *redisserver.Redis
}

// Register(ctx context.Context, req *models.Register_User) (string, error)
// 	Verify(ctx context.Context, req *models.Verify_User) (string, error)
// 	LogIn(ctx context.Context, req *models.LogIn) (string, error)
// 	CreateTask(ctx context.Context, req *models.Task) (*models.GetTaskResponse, error)
// 	GetTask(ctx context.Context, req *models.GetTaskRequest) (*models.GetTaskResponse, error)
// 	GetTasks(ctx context.Context, req int) ([]*models.GetTaskResponse, error)
// 	UpdateTask(ctx context.Context, req *models.Task) (*models.GetTaskResponse, error)
// 	DeleteTask(ctx context.Context, req *models.GetTaskRequest) (string, error)

func (u *Service) Register(ctx context.Context, req *models.Register_User) (string, error) {
	if err := u.Redis.CheckUser(req.Email); err != nil {
		return err.Error(), nil
	}
	if err := u.Redis.Register_User(req); err != nil {
		return err.Error(),nil
	}
	code, err := email1.Sent(req.Email)
	if err != nil {
		return err.Error(),nil
	}
	if err := u.Redis.VerificationCodeSave(&models.Verify_User{Email: req.Email, Code: code}); err != nil {
		return  err.Error(),nil
	}
	return "Confirmation code has been sent to your email check it's delivered to verify your account", nil
}

func (u *Service) Verify(ctx context.Context, req *models.Verify_User) (string, error) {
	code, err := u.Redis.VerificationCodeGet(req.Email)
	if err != nil {
		return "email address doesn't match", nil
	}
	if code != req.Code {
		return "confirmation code isn't matching", nil
	}
	res, err := u.Redis.GetUser(req.Email)
	if err != nil {
		log.Println(err)
	}
	_, err = u.Storage.GetUser(ctx, sql.NullString{String: req.Email, Valid: true})
	if err == nil {
		return "this email is already existed", nil
	}
	password := Hashing(res.Password)
	fmt.Println(password)
	if err := u.Storage.CreateUser(ctx, storage.CreateUserParams{
		Username: sql.NullString{String: res.Username, Valid: true},
		Email:    sql.NullString{String: res.Email, Valid: true},
		Password: sql.NullString{String: password, Valid: true}}); err != nil {
		log.Println("error is there", err)
	}
	return "Successfully verified your account and now please LOGIN", nil
}

func (u *Service) LogIn(ctx context.Context, req *models.LogIn) (string, error) {
	res, err := u.Storage.GetUser(ctx, sql.NullString{String: req.Email, Valid: true})
	if err == sql.ErrNoRows {
		return "email address is wrong", nil
	}
	if err != nil {
		return err.Error(), nil
	}
	check := ComparePassword(res.Password.String, req.Password)
	if !check {
		return "password ain't match", nil
	}
	token, err := jwttoken.CreateToken(&models.Register_User{Username: res.Username.String, Email: res.Email.String})
	if err != nil {
		log.Println(err)
	}
	return token, nil
}

func (u *Service) CreateTask(ctx context.Context, req *models.Task) (*models.GetTaskResponse, error) {
	res, err := u.Storage.CreateTask(ctx, storage.CreateTaskParams{UserID: sql.NullInt32{
		Int32: int32(req.UserID), Valid: true},
		Title:     sql.NullString{String: req.Title, Valid: true},
		CreatedAt: sql.NullString{String: time.Now().Format(time.RFC3339), Valid: true},
	})
	if err != nil {
		return nil, err
	}
	if err := u.Mongo.CreateTask(&models.MonogTask{ID: int(res.ID), UserID: int(res.UserID.Int32), Status: req.Status, Description: req.Description, Important: req.Important}); err != nil {
		return nil, err
	}
	res1, err := u.Mongo.GetTask(&models.GetTaskRequest{ID: int(res.ID), UserID: req.UserID})
	if err != nil {
		return nil, err
	}
	return &models.GetTaskResponse{
		ID:          int(res.ID),
		UserID:      int(res.UserID.Int32),
		Title:       res.Title.String,
		Status:      res1.Status,
		Description: res1.Description,
		Important:   res1.Important,
		CreatedAt:   res.CreatedAt.String,
		UpdatedAt:   res.UpdatedAt.String}, nil
}

func (u *Service) GetTask(ctx context.Context, req *models.GetTaskRequest) (*models.GetTaskResponse, error) {
	res, err := u.Storage.GetTask(ctx, storage.GetTaskParams{ID: int32(req.ID), UserID: sql.NullInt32{Int32: int32(req.UserID), Valid: true}})
	if err != nil {
		return nil, err
	}
	res1, err := u.Mongo.GetTask(&models.GetTaskRequest{ID: req.ID, UserID: req.UserID})
	if err != nil {
		return nil, err
	}
	return &models.GetTaskResponse{
		ID:          int(res.ID),
		UserID:      int(res.UserID.Int32),
		Title:       res.Title.String,
		Status:      res1.Status,
		Description: res1.Description,
		Important:   res1.Important,
		CreatedAt:   res.CreatedAt.String,
		UpdatedAt:   res.UpdatedAt.String}, nil
}

func (u *Service) GetTasks(ctx context.Context, req int) ([]*models.GetTaskResponse, error) {
	res, err := u.Storage.ListTasks(ctx, sql.NullInt32{Int32: int32(req), Valid: true})
	if err != nil {
		return nil, err
	}
	res1, err := u.Mongo.GetTasks(req)
	if err != nil {
		return nil, err
	}
	fmt.Println(res)
	fmt.Println(res1)
	var tasks []*models.GetTaskResponse

	for i := 0; i < len(res); i++ {
		var all = models.GetTaskResponse{
			ID:          int(res[i].ID),
			UserID:      int(res[i].UserID.Int32),
			Title:       res[i].Title.String,
			Status:      res1[i].Status,
			Description: res1[i].Description,
			Important:   res1[i].Important,
			CreatedAt:   res[i].CreatedAt.String,
			UpdatedAt:   res[i].UpdatedAt.String,
		}
		tasks = append(tasks, &all)
	}
	return tasks, nil
}

func (u *Service) UpdateTask(ctx context.Context, req *models.Task) (*models.GetTaskResponse, error) {
	res, err := u.Storage.UpdateTask(ctx, storage.UpdateTaskParams{ID: int32(req.ID), UserID: sql.NullInt32{Int32: int32(req.UserID), Valid: true}, Title: sql.NullString{String: req.Title, Valid: true}, UpdatedAt: sql.NullString{String: time.Now().Format(time.RFC3339)}})
	if err != nil {
		return nil, err
	}
	res1, err := u.Mongo.UpdateTask(&models.MonogTask{ID: req.ID, UserID: req.UserID, Status: req.Status, Description: req.Description, Important: req.Important})
	if err != nil {
		return nil, err
	}
	return &models.GetTaskResponse{
		ID:          int(res.ID),
		UserID:      int(res.UserID.Int32),
		Title:       res.Title.String,
		Status:      res1.Status,
		Description: res1.Description,
		Important:   res1.Important,
		CreatedAt:   res.CreatedAt.String,
		UpdatedAt:   res.UpdatedAt.String}, nil
}

func (u *Service) DeleteTask(ctx context.Context, req *models.GetTaskRequest) (string, error) {
	if err := u.Storage.DeleteTask(ctx, storage.DeleteTaskParams{ID: int32(req.ID), UserID: sql.NullInt32{Int32: int32(req.UserID), Valid: true}}); err != nil {
		return "", err
	}
	if err := u.Mongo.DeleteTask(&models.GetTaskRequest{ID: req.ID, UserID: req.UserID}); err != nil {
		return "", err
	}
	return "Sucessfully deleted", nil
}

func ComparePassword(hashed, password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password)); err != nil {
		log.Println(err)
		return false
	}
	return true
}

func Hashing(password string) string {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return ""
	}
	return string(hashed)
}
