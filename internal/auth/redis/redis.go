package redisserver

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/unknownn17/Internship_Task/internal/models"
)

type Redis struct {
	R *redis.Client
	C context.Context
}

func (u *Redis) Register_User(req *models.Register_User) error {
	byted, err := json.Marshal(req)
	if err != nil {
		log.Println(err)
		return err
	}
	if err := u.R.Set(u.C, req.Email, byted, time.Duration(time.Minute*5)).Err(); err != nil {
		return err
	}
	return nil
}

func (u *Redis) VerificationCodeSave(req *models.Verify_User) error {
	if err := u.R.Set(u.C, req.Email+"code", req.Code, time.Duration(time.Minute*5)).Err(); err != nil {
		return err
	}
	return nil
}

func (u *Redis) VerificationCodeGet(req string) (string, error) {
	var code string

	if err := u.R.Get(u.C, req+"code").Scan(&code); err != nil {
		return "", err
	}
	return code, nil
}

func (u *Redis) CheckUser(req string) error {
	_, err := u.R.Get(u.C, req).Result()
	if err == redis.Nil {
		return nil
	} else if err != nil {
		return err
	}
	return errors.New("user already exists with this email")
}


func (u *Redis) GetUser(req string) (*models.Register_User, error) {

	res, err := u.R.Get(u.C, req).Result()
	if err != nil {
		return nil, err
	}
	var user models.Register_User

	if err := json.Unmarshal([]byte(res), &user); err != nil {
		return nil, err
	}
	return &user, nil
}
