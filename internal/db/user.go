package db

import (
	"context"

	"github.com/qiniu/qmgo"
	"github.com/qiniu/qmgo/field"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User 账号
type User struct {
	field.DefaultField `bson:",inline"`
	Username           string `bson:"username" json:"username"`
	Nickname           string `bson:"nickname" json:"nickname"`
	Phone              string `bson:"phone" json:"phone"`
	Email              string `bson:"email" json:"email"`
	Password           string `bson:"password" json:"-"`
	Avatar             string `bson:"avatar" json:"avatar"`
	Language           string `bson:"language" json:"language"`
	IsAdmin            bool   `bson:"isAdmin" json:"isAdmin"`
}

func GetUserCollection() *qmgo.Collection {
	return DB().Collection("user")
}

func GetUserByLogin(ctx context.Context, login string) (u *User, err error) {
	err = GetUserCollection().
		Find(
			ctx,
			bson.M{"$or": []bson.M{{"username": login}, {"phone": login}, {"email": login}}},
		).One(&u)
	return
}

func GetUserByID(ctx context.Context, id string) (u *User, err error) {
	uid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	err = GetUserCollection().
		Find(
			ctx,
			bson.M{"_id": uid},
		).One(&u)
	return
}
