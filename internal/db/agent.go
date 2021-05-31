package db

import "github.com/qiniu/qmgo/field"

// Agent 客服
type Agent struct {
	field.DefaultField `bson:",inline"`
	UserID             string `bson:"userID"`
	Status             int    `bson:"status"`
}
