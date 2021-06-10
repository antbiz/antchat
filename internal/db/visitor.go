package db

import (
	"context"
	"fmt"

	"github.com/gogf/gf/frame/g"
	"github.com/qiniu/qmgo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Visitor 访客
type Visitor struct {
	DefaultField `bson:",inline"`
	Domain       string `bson:"domain"`
	Refer        string `bson:"refer"`
	Token        string `bson:"token"`
	Nickname     string `bson:"nickname"`
	Email        string `bson:"email"`
	Phone        string `bson:"phone"`
	Geo          string `bson:"geo"`
	IP           string `bson:"ip"`
	Country      string `bson:"country"`
	City         string `bson:"city"`
	AgentID      string `bons:"agentID"`
}

func (visitor *Visitor) Address() string {
	return fmt.Sprintf("%s%s", visitor.Country, visitor.City)
}

func GetVisitorCollection() *qmgo.Collection {
	return DB().Collection("visitor")
}

func GetVisitorByID(ctx context.Context, id string) (visitor *Visitor, err error) {
	vid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	err = GetVisitorCollection().
		Find(
			ctx,
			bson.M{"_id": vid},
		).One(&visitor)
	return
}

func CreateVisitor(ctx context.Context, visitor *Visitor) (id string, err error) {
	res, err := GetVisitorCollection().InsertOne(ctx, visitor)
	if err != nil {
		g.Log().Async().Errorf("db.CreateVisitor: %v", err)
		return "", err
	}
	id = res.InsertedID.(primitive.ObjectID).Hex()
	return
}

func UpsertVisitor(ctx context.Context, visitor *Visitor) (id string, err error) {
	res, err := GetVisitorCollection().UpsertId(ctx, visitor.ID, visitor)
	if err != nil {
		g.Log().Async().Errorf("db.UpsertVisitor: %v", err)
		return "", err
	}
	if res.UpsertedID == nil {
		id = visitor.ID.Hex()
	} else {
		id = res.UpsertedID.(primitive.ObjectID).Hex()
	}
	return
}
