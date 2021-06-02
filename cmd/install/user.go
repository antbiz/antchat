package main

import (
	"context"

	"github.com/antbiz/antchat/internal/app/admin/service"
	"github.com/antbiz/antchat/internal/db"
	"github.com/gogf/gf/frame/g"
	"go.mongodb.org/mongo-driver/bson"
)

// initAdminAccount 初始化超级管理员账号
func initAdminAccount() {
	username := g.Cfg().GetString("admin.username")
	password := g.Cfg().GetString("admin.password")
	email := g.Cfg().GetString("admin.email")
	if username == "" || password == "" {
		panic("admin username or password is empty")
	}
	password = service.User.EncryptPwd(username, password)
	_, err := db.GetUserCollection().
		Upsert(context.Background(), bson.M{"username": username}, &db.User{
			Email:    email,
			Username: username,
			Password: password,
		})
	if err != nil {
		panic(err)
	}
	g.Log().Debug("Init admin account successfully!")
}
