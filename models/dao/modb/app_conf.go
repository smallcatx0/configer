package modb

import (
	"context"
	"fmt"
	"gtank/models/dao"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ctx = context.TODO()
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

func (e *AppEnv) Save() (id string, err error) {
	// 先检查是否存在
	err = e.Col().FindOne(ctx, bson.D{
		{"sign", e.Sign},
	}).Err()
	if err == nil {
		return "", fmt.Errorf("环境已存在")
	}
	e.ID = primitive.NewObjectID()
	res, err := e.Col().InsertOne(ctx, e)
	if err != nil {
		return "", err
	}
	id = res.InsertedID.(primitive.ObjectID).Hex()
	return
}

func (e *AppEnv) Edit() (count int64, err error) {
	up := bson.D{}
	if e.Name != "" {
		up = append(up, bson.E{"name", e.Name})
	}
	if e.Desc != "" {
		up = append(up, bson.E{"desc", e.Desc})
	}
	if e.Principal.Name != "" {
		up = append(up, bson.E{"principal", e.Principal})
	}
	res, err := e.Col().UpdateOne(ctx, bson.D{
		{"sign", e.Sign},
	}, bson.D{
		{"$set", up},
	})
	if err != nil {
		return
	}
	count = res.ModifiedCount
	return
}

func (e *AppEnv) List() (list []AppEnv, err error) {
	q := bson.D{}
	res, err := e.Col().Find(ctx, q)
	if err != nil {
		return nil, err
	}
	err = res.All(ctx, &list)
	if err != nil {
		return nil, err
	}
	return
}

func (e *AppEnv) Del() (count int64, err error) {
	q := bson.M{
		"sign": e.Sign,
	}
	res, err := e.Col().DeleteOne(ctx, q)
	if err != nil {
		return
	}
	count = res.DeletedCount
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
