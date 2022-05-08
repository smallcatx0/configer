package modb

import (
	"context"
	"fmt"
	"gtank/models/dao"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Contact struct {
	Name  string
	Phone string
}

type AppEnv struct {
	ID        primitive.ObjectID `bson:"_id"`
	Name      string             `bson:"name"`
	Type      string             `bson:"type"`
	Sign      string             `bson:"sign"`
	Desc      string             `bson:"desc"`
	Principal Contact            `bson:"principal"`
}

func (AppEnv) Col() *mongo.Collection {
	return dao.MongoDb.Collection("env")
}

func (d *AppEnv) Save() (id interface{}, err error) {
	// 先检查是否存在
	ctx := context.TODO()
	err = d.Col().FindOne(ctx, bson.D{
		{"sign", d.Sign},
	}).Err()
	if err == nil {
		return "", fmt.Errorf("环境已存在")
	}
	d.ID = primitive.NewObjectID()
	res, err := d.Col().InsertOne(ctx, d)
	if err != nil {
		return "", err
	}
	id = res.InsertedID.(primitive.ObjectID).Hex()
	return
}

type AppConf struct {
	AppName   string
	AppSign   string
	Env       string
	FileName  string
	Content   string
	CreateAt  string
	Principal Contact
}
