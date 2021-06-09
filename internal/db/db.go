package db

import (
	"context"
	"sync"
	"time"

	"github.com/gogf/gf/frame/g"
	"github.com/qiniu/qmgo"
	"github.com/qiniu/qmgo/field"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	cli  *qmgo.Client
	once sync.Once
)

type DefaultField struct {
	ID        primitive.ObjectID `bson:"_id"`
	CreatedAt time.Time          `bson:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt"`
}

func (DefaultField) CustomFields() field.CustomFieldsBuilder {
	return field.NewCustom().SetId("ID").SetCreateAt("CreatedAt").SetUpdateAt("UpdatedAt")
}

// Cli is mongo client
func Cli() *qmgo.Client {
	once.Do(func() {
		var err error
		mongoURI := g.Cfg().GetString("mongo.uri")
		cli, err = qmgo.NewClient(
			context.Background(),
			&qmgo.Config{
				Uri: mongoURI,
			},
		)
		if err != nil {
			g.Log().Errorf("failed to connect mongo: %s", mongoURI)
		} else {
			g.Log().Debugf("connected mongo: %s", mongoURI)
		}
	})
	return cli
}

// DB 数据库实例
func DB(database ...string) *qmgo.Database {
	var dbName string
	if len(database) > 0 {
		dbName = database[0]
	} else {
		dbName = g.Cfg().GetString("mongo.default")
	}
	return Cli().Database(dbName)
}
