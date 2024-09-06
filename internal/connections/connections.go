package connections

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"github.com/redis/go-redis/v9"
	"github.com/unknownn17/Internship_Task/internal/api/handler"
	redisserver "github.com/unknownn17/Internship_Task/internal/auth/redis"
	"github.com/unknownn17/Internship_Task/internal/config"
	mongodb "github.com/unknownn17/Internship_Task/internal/database/mongoDB"
	interface17 "github.com/unknownn17/Internship_Task/internal/interface"
	"github.com/unknownn17/Internship_Task/internal/interface/impliment"
	services "github.com/unknownn17/Internship_Task/internal/service"
	"github.com/unknownn17/Internship_Task/internal/database/sqlc/storage"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoCollection() *mongodb.Mongo {
	opt, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://mongo:27017"))
	if err != nil {
		log.Println(err)
	}
	if err := opt.Ping(context.Background(), options.Client().ReadPreference); err != nil {
		log.Println(err)
	}
	ctx := context.Background()
	data := opt.Database("internship").Collection("tasks")
	return &mongodb.Mongo{M: data, C: ctx}
}

func Redis() *redisserver.Redis {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})
	ctx := context.Background()
	_, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatal(err)
	}
	return &redisserver.Redis{R: client, C: ctx}
}

func Adjust() interface17.TaskManagement {
	c := config.Configuration()

	db, err := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", c.Database.User, c.Database.Password, c.Database.Host, c.Database.DBname))
	if err != nil {
		log.Println(err)
	}
	if err := db.Ping(); err != nil {
		log.Println(err)
	}
	queries := storage.New(db)

	mongo := MongoCollection()
	redis := Redis()
	return &services.Service{Storage: queries, Mongo: mongo, Redis: redis}
}

func NewService() *impliment.Service {
	a := Adjust()
	return &impliment.Service{I: a}
}

func NewHandler() *handler.Handler {
	a := NewService()
	ctx := context.Background()
	return &handler.Handler{S: a, C: ctx}
}
