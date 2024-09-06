package mongodb

import (
	"context"
	"errors"
	"log"

	"github.com/unknownn17/Internship_Task/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct {
	M *mongo.Collection
	C context.Context
}

func (u *Mongo) CreateTask(req *models.MonogTask) error {
	_, err := u.M.InsertOne(u.C, req)
	if err != nil {
		return err
	}
	return nil
}

func (u *Mongo) GetTask(req *models.GetTaskRequest) (*models.MonogTask, error) {
	var res models.MonogTask

	err := u.M.FindOne(u.C, bson.M{"id": req.ID, "user_id": req.UserID}).Decode(&res)
	if err == mongo.ErrNoDocuments {
		return nil, errors.New("there is no such task ")
	}
	return &res, nil
}

func (u *Mongo) GetTasks(req int) ([]*models.MonogTask, error) {
	var res []*models.MonogTask

	opts := options.Find().SetSort(bson.M{"id":1}) 
	tasks, err := u.M.Find(u.C, bson.M{"user_id": req}, opts)
	if err != nil {
		log.Println("Error finding tasks:", err)
		return nil, err
	}
	for tasks.Next(u.C) {
		var task models.MonogTask
		if err := tasks.Decode(&task); err != nil {
			log.Println("Error decoding task:", err)
			return nil, err
		}
		res = append(res, &task)
	}
	if err := tasks.Err(); err != nil {
		log.Println("Cursor error:", err)
		return nil, err
	}

	return res, nil
}


func (u *Mongo) UpdateTask(req *models.MonogTask) (*models.MonogTask, error) {
	var res models.MonogTask
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	err := u.M.FindOneAndUpdate(u.C, bson.M{"id": req.ID, "user_id": req.UserID}, bson.M{"$set": bson.M{"status": req.Status, "description": req.Description, "important": req.Important}}, opts).Decode(&res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (u *Mongo) DeleteTask(req *models.GetTaskRequest) error {
	_, err := u.M.DeleteOne(u.C, bson.M{"id": req.ID, "user_id": req.UserID})
	if err != nil {
		return errors.New("there is no such task")
	}
	return nil
}
