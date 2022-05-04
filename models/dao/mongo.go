package dao

import (
	"context"
	"gtank/internal/conf"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MongoCli *mongo.Client

	MongoDb *mongo.Database
)

func ConnMongo(uri string) (*mongo.Client, error) {
	opt := options.Client().ApplyURI(uri)

	c, _ := context.WithTimeout(context.Background(), 5*time.Second)
	cli, err := mongo.Connect(c, opt)
	if err != nil {
		return nil, err
	}
	// 探活
	timeout, _ := context.WithTimeout(context.Background(), 2*time.Second)
	err = cli.Ping(timeout, nil)
	if err != nil {
		return nil, err
	}
	MongoCli = cli
	return cli, nil
}

// 初始化MongoDb
func InitMongo() error {
	c := conf.AppConf
	uri := c.GetString("mongo.uri")
	db := c.GetString("mongo.db")

	cli, err := ConnMongo(uri)
	if err != nil {
		return err
	}
	MongoDb = cli.Database(db)
	return nil
}
