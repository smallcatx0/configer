package modb

import (
	"errors"
	"gtank/models/dao"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	DbCronStatus_on  = "online"
	DbCronStatus_off = "offline"
)

type DbTTL struct {
	ID     primitive.ObjectID `bson:"_id" json:"id"`
	DSN    string             `bson:"dsn" json:"dsn"`
	DB     string             `bson:"db" json:"db"`
	Table  string             `bson:"table" json:"table"`
	Field  string             `bson:"field" json:"field"`
	Cron   string             `bson:"cron" json:"cron"`
	TTL    int                `bson:"ttl" json:"ttl"`
	Limit  int                `bson:"limit" json:"limit"`
	Status string             `bson:"status" json:"status"`
	Desc   string             `bson:"desc" json:"desc"`
}

func (DbTTL) Col() *mongo.Collection {
	return dao.MongoDb.Collection("dbcron_ttl")
}

func (d *DbTTL) IsExist() (ok bool, err error) {
	// 先检查是否存在
	q := bson.D{
		{Key: "db", Value: d.DB},
		{Key: "table", Value: d.Table},
	}
	err = d.Col().FindOne(ctx, q).Err()
	if err == nil {
		return true, nil
	}
	if errors.Is(err, mongo.ErrNoDocuments) {
		return false, nil
	} else {
		return false, err
	}

}

func (d *DbTTL) NewOne() (id string, err error) {
	d.ID = primitive.NewObjectID()
	d.Status = DbCronStatus_on
	res, err := d.Col().InsertOne(ctx, d)
	if err != nil {
		return "", err
	}
	id = res.InsertedID.(primitive.ObjectID).Hex()
	return
}

func (d *DbTTL) Edit(up bson.D) (cont int64, err error) {
	res, err := d.Col().UpdateByID(ctx, d.ID, bson.D{{
		Key: "$set", Value: up,
	}})
	if err != nil {
		return
	}
	cont = res.ModifiedCount
	return
}

func (d *DbTTL) Del() error {
	_, err := d.Col().DeleteOne(ctx, bson.D{
		{Key: "_id", Value: d.ID},
	})
	return err
}

func (d *DbTTL) List(q bson.D, skip, limit int) (list []DbTTL, err error) {
	opt := options.Find()
	opt.SetSkip(int64(skip))
	opt.SetLimit(int64(limit))
	cur, err := d.Col().Find(ctx, q, opt)
	if err != nil {
		return nil, err
	}
	err = cur.All(ctx, &list)
	if err != nil {
		return nil, err
	}
	return
}
