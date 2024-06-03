package mongo

import (
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"subscriptions/core/errors"
	"time"
)

const (
	ConnectionFailed = "Db:Conn:Fail"
)

type Manager struct {
	Client *mongo.Client
	Random string
}

type SubscriptionPackage struct {
	Id       uuid.UUID
	Name     string
	Features []uuid.UUID
}

func NewManager(setting ConnectionSettings) Manager {
	con := context.TODO()
	context.WithTimeout(con, time.Second*30)

	println(setting.GetConnectionString())

	client, err := mongo.Connect(
		con,
		options.Client().ApplyURI(setting.GetConnectionString()))

	if err != nil {
		panic(errors.CodeError{
			Code:    ConnectionFailed,
			Message: "Failed To Create The Manager",
		})
	}

	return Manager{
		Random: "stf",
		Client: client,
	}
}
