package db

import "github.com/qiniu/qmgo/field"

// Visitor 访客
type Visitor struct {
	field.DefaultField `bson:",inline"`
	Domain             string `bson:"domain"`
	Refer              string `bson:"refer"`
	Token              string `bson:"token"`
	Nickname           string `bson:"nickname"`
	Email              string `bson:"email"`
	Phone              string `bson:"phone"`
	Geo                string `bson:"geo"`
	IP                 string `bson:"ip"`
	Country            string `bson:"country"`
	City               string `bson:"city"`
}
