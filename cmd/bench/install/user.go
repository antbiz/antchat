package install

import (
	"context"

	"github.com/antbiz/antchat/internal/app/system/service"
	"github.com/antbiz/antchat/internal/db"
	"github.com/gogf/gf/frame/g"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// initAdmin 初始化超级管理员账号
func initAdmin() {
	ctx := context.Background()
	username := g.Cfg().GetString("admin.username")
	password := g.Cfg().GetString("admin.password")
	email := g.Cfg().GetString("admin.email")
	if username == "" || password == "" {
		panic("admin username or password is empty")
	}
	password = service.User.EncryptPwd(username, password)

	admin := &db.User{
		Email:    email,
		Username: username,
		Password: password,
		Nickname: username,
		IsAdmin:  true,
	}
	adminUser, _ := db.GetUserByUsername(ctx, username)
	if adminUser != nil {
		admin.ID = adminUser.ID
	}
	res, err := db.GetUserCollection().
		UpsertId(ctx, admin.ID, admin)
	if err != nil {
		panic(err)
	}
	if res.UpsertedID != nil {
		admin.ID = res.UpsertedID.(primitive.ObjectID)
	}
	g.Log().Debug("Init admin account successfully!")

	adminStrUID := admin.ID.Hex()
	agent := &db.Agent{
		UserID: adminStrUID,
	}
	adminAgent, _ := db.GetAgentByUID(ctx, adminStrUID)
	if adminAgent != nil {
		agent.ID = adminAgent.ID
	}
	_, err = db.GetAgentCollection().
		UpsertId(ctx, agent.ID, agent)
	if err != nil {
		panic(err)
	}
	g.Log().Debug("Init admin agent successfully!")
}
