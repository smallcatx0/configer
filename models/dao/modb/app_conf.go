package modb

import (
	"context"
	"fmt"
	"gtank/models/dao"
	"time"

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

func (e *AppEnv) NewOne() (id string, err error) {
	// 先检查是否存在
	err = e.Col().FindOne(ctx, bson.D{
		{"sign", e.Sign},
	}).Err()
	if err == nil {
		return "", fmt.Errorf("资源已存在")
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

type App struct {
	ID        primitive.ObjectID `bson:"_id"`
	Name      string
	Sign      string
	Desc      string
	CreateAt  time.Time
	Principal Contact
}

func (App) Col() *mongo.Collection {
	return dao.MongoDb.Collection("app")
}

func (a *App) NewOne() (id string, err error) {
	q := bson.D{
		{"sign", a.Sign},
	}
	err = a.Col().FindOne(ctx, q).Err()
	if err == nil {
		return "", fmt.Errorf("资源已存在")
	}
	a.ID = primitive.NewObjectID()
	a.CreateAt = time.Now()
	res, err := a.Col().InsertOne(ctx, a)
	if err != nil {
		return "", err
	}
	id = res.InsertedID.(primitive.ObjectID).Hex()
	return
}

func (a *App) Edit() (count int64, err error) {
	up := bson.D{}
	if a.Name != "" {
		up = append(up, bson.E{"name", a.Name})
	}
	if a.Desc != "" {
		up = append(up, bson.E{"desc", a.Desc})
	}
	if a.Principal.Name != "" {
		up = append(up, bson.E{"principal", a.Principal})
	}
	res, err := a.Col().UpdateOne(ctx, bson.D{
		{"sign", a.Sign},
	}, bson.D{
		{"$set", up},
	})
	if err != nil {
		return
	}
	count = res.ModifiedCount
	return
}
func (a *App) List() (list []App, err error) {
	q := bson.D{}
	res, err := a.Col().Find(ctx, q)
	if err != nil {
		return nil, err
	}
	err = res.All(ctx, &list)
	if err != nil {
		return nil, err
	}
	return
}

func (a *App) Del() (count int64, err error) {
	q := bson.M{
		"sign": a.Sign,
	}
	res, err := a.Col().DeleteOne(ctx, q)
	if err != nil {
		return
	}
	count = res.DeletedCount
	return
}
