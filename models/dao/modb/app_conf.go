package modb

import (
	"context"
	"errors"
	"fmt"
	"gtank/middleware/resp"
	"gtank/models/dao"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	ctx = context.TODO()
)

type Contact struct {
	Name  string `bson:"name" json:"name"`
	Phone string `bson:"phone" json:"phone"`
}

type AppEnv struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	Name      string             `bson:"name" json:"name"`
	Type      string             `bson:"type" json:"type"`
	Sign      string             `bson:"sign" json:"sign"`
	Desc      string             `bson:"desc" json:"desc"`
	Principal Contact            `bson:"principal" json:"principal"`
}

func (AppEnv) Col() *mongo.Collection {
	return dao.MongoDb.Collection("env")
}

func (e *AppEnv) NewOne() (id string, err error) {
	// 先检查是否存在
	err = e.Col().FindOne(ctx, bson.D{
		{Key: "sign", Value: e.Sign},
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
		up = append(up, bson.E{Key: "name", Value: e.Name})
	}
	if e.Desc != "" {
		up = append(up, bson.E{Key: "desc", Value: e.Desc})
	}
	if e.Principal.Name != "" {
		up = append(up, bson.E{Key: "principal", Value: e.Principal})
	}
	res, err := e.Col().UpdateOne(ctx, bson.D{
		{Key: "sign", Value: e.Sign},
	}, bson.D{
		{Key: "$set", Value: up},
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

func (e *AppEnv) GetbySign(sign string) error {
	res := e.Col().FindOne(ctx, bson.D{{Key: "sign", Value: sign}})
	if res.Err() != nil {
		return res.Err()
	}
	if err := res.Decode(e); err != nil {
		return err
	}
	return nil
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
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	Name      string             `bson:"name" json:"name"`
	Sign      string             `bson:"sign" json:"sign"`
	Desc      string             `bson:"desc" json:"desc"`
	CreateAt  time.Time          `bson:"create_at" json:"create_at"`
	Principal Contact            `bson:"principal" json:"principal"`
}

func (App) Col() *mongo.Collection {
	return dao.MongoDb.Collection("app")
}

func (a *App) NewOne() (id string, err error) {
	q := bson.D{
		{Key: "sign", Value: a.Sign},
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

func (a *App) GetbySign(sign string) error {
	q := bson.D{{Key: "sign", Value: sign}}
	res := a.Col().FindOne(ctx, q)
	if res.Err() != nil {
		return res.Err()
	}
	if err := res.Decode(a); err != nil {
		return err
	}
	return nil
}

func (a *App) Edit() (count int64, err error) {
	up := bson.D{}
	if a.Name != "" {
		up = append(up, bson.E{Key: "name", Value: a.Name})
	}
	if a.Desc != "" {
		up = append(up, bson.E{Key: "desc", Value: a.Desc})
	}
	if a.Principal.Name != "" {
		up = append(up, bson.E{Key: "principal", Value: a.Principal})
	}
	res, err := a.Col().UpdateOne(ctx, bson.D{
		{Key: "sign", Value: a.Sign},
	}, bson.D{
		{Key: "$set", Value: up},
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

type ConfFile struct {
	Name    string `bson:"name" json:"name"`
	Content string `bson:"content" json:"content"`
	Type    string `bson:"type" json:"type"`
}

type AppFile struct {
	ID      primitive.ObjectID `bson:"_id" json:"id"`
	AppSign string             `bson:"app_sign" json:"app_sign"`
	EnvSign string             `bson:"env_sign" json:"env_sign"`
	Files   map[string]string  `bson:"files" json:"files"`
}

func (AppFile) Col() *mongo.Collection {
	return dao.MongoDb.Collection("app_file")
}

var (
	FileStatus_Enable  = "enable"
	FileStatus_Disable = "disable"
)

func (f *AppFile) AddFile(name string) error {
	res := f.Col().FindOne(ctx, bson.D{
		{Key: "env_sign", Value: f.EnvSign},
		{Key: "app_sign", Value: f.AppSign},
	})
	if res.Err() == nil {
		// 更新
		if err := res.Decode(f); err != nil {
			return err
		}
		if f.Files[name] != "" {
			return nil
		}
		f.Files[name] = FileStatus_Enable
		_, err := f.Col().UpdateByID(ctx, f.ID, bson.D{
			{Key: "$set", Value: bson.D{{Key: "files", Value: f.Files}}},
		})
		if err != nil {
			return err
		}
	}
	if errors.Is(res.Err(), mongo.ErrNoDocuments) {
		// 新增
		f.ID = primitive.NewObjectID()
		f.Files = map[string]string{
			name: FileStatus_Enable,
		}
		_, err := f.Col().InsertOne(ctx, f)
		return err
	} else {
		return res.Err()
	}
}

func (f *AppFile) DelFile(name string) error {
	res := f.Col().FindOne(ctx, bson.D{
		{Key: "env_sign", Value: f.EnvSign},
		{Key: "app_sign", Value: f.AppSign},
	})
	if res.Err() == nil {
		if err := res.Decode(f); err != nil {
			return err
		}
		delete(f.Files, name)
		_, err := f.Col().UpdateByID(ctx, f.ID, bson.D{
			{Key: "$set", Value: bson.D{{Key: "files", Value: f.Files}}},
		})
		if err != nil {
			return err
		}
	} else if errors.Is(res.Err(), mongo.ErrNoDocuments) {
		return nil
	}
	return res.Err()
}

func (f *AppFile) GetFiles() (err error) {
	res := f.Col().FindOne(ctx, bson.D{
		{Key: "env_sign", Value: f.EnvSign},
		{Key: "app_sign", Value: f.AppSign},
	})
	if errors.Is(res.Err(), mongo.ErrNoDocuments) {
		return resp.NotFund
	}
	if err := res.Decode(f); err != nil {
		return err
	}
	return
}

type AppBaseInfo struct {
	Name    string `bson:"name" json:"name"`
	Sign    string `bson:"sign" json:"sign"`
	Contact string `bson:"contact" json:"contact"`
}

type EnvBaseInfo struct {
	Type    string `bson:"type" json:"type"`
	Sign    string `bson:"sign" json:"sign"`
	Name    string `bson:"name" json:"name"`
	Contact string `bson:"contact" json:"contact"`
}

type AppConf struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	App       AppBaseInfo        `bson:"app" json:"app"`
	Env       EnvBaseInfo        `bson:"env" json:"env"`
	File      ConfFile           `bson:"file" json:"file"`
	Principal Contact            `bson:"principal" json:"principal"`
	CreateAt  time.Time          `bson:"create_at" json:"create_at"`
	Status    string             `bson:"status" json:"status"`
}

func (AppConf) Col() *mongo.Collection {
	return dao.MongoDb.Collection("appconf")
}

func (f *AppConf) SaveOne() (id string, err error) {
	f.ID = primitive.NewObjectID()
	f.CreateAt = time.Now()
	res, err := f.Col().InsertOne(ctx, f)
	if err != nil {
		return "", err
	}
	id = res.InsertedID.(primitive.ObjectID).Hex()
	return
}

func (f *AppConf) FileHistory(env string, app string, filename string) (list []AppConf, err error) {
	q := bson.D{
		{Key: "app.sign", Value: app},
		{Key: "env.sign", Value: env},
		{Key: "file.name", Value: filename},
	}
	opt := options.Find()
	opt.SetLimit(500)
	opt.SetSort(bson.D{{Key: "create_at", Value: -1}})
	cur, err := f.Col().Find(ctx, q, opt)
	if err != nil {
		return nil, err
	}
	err = cur.All(ctx, &list)
	if err != nil {
		return nil, err
	}
	return
}

func (f *AppConf) Top(env, app, file string) (*AppConf, error) {
	q := bson.D{
		{Key: "env.sign", Value: env},
		{Key: "app.sign", Value: app},
		{Key: "file.name", Value: file},
	}
	opt := options.FindOne()
	opt.SetSort(bson.D{{Key: "create_at", Value: -1}})
	cur := f.Col().FindOne(ctx, q, opt)
	conf := &AppConf{}
	if err := cur.Decode(conf); err != nil {
		return nil, err
	}
	return conf, nil
}
