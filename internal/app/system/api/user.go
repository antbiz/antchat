package api

import (
	"github.com/antbiz/antchat/internal/app/system/dto"
	"github.com/antbiz/antchat/internal/app/system/service"
	"github.com/antbiz/antchat/internal/db"
	"github.com/antbiz/antchat/internal/pkg/resp"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"go.mongodb.org/mongo-driver/mongo"
)

// User 用户接口
var User = &userApi{}

type userApi struct{}

// SigninByAccount 用户账号（用户名/手机号/邮箱）登录
func (userApi) SigninByAccount(r *ghttp.Request) {
	var req *dto.UserSigninReq
	if err := r.Parse(&req); err != nil {
		resp.InvalidArgument(r, err.Error())
	}

	user, err := db.GetUserBySignin(r.Context(), req.Account)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			resp.NotFound(r, "account_not_found", "该账号尚未注册")
		}
		resp.DatabaseError(r, "账号查询失败", err)
	}

	if user.Password != service.User.EncryptPwd(user.Username, req.Password) {
		resp.Unauthorized(r, "incorrect_password", "密码错误")
	}

	var agentID string
	if agent, _ := db.GetAgentByUID(r.Context(), user.ID.Hex()); agent != nil {
		agentID = agent.ID.Hex()
	}

	sessionData := g.Map{
		"id":       user.ID.Hex(),
		"agentID":  agentID,
		"username": user.Username,
		"nickname": user.Nickname,
		"phone":    user.Phone,
		"email":    user.Email,
		"avatar":   user.Avatar,
		"language": user.Language,
		"sid":      r.Session.Id(),
	}
	r.Session.SetMap(sessionData)
	resp.OK(r, sessionData)
}

// LogOut 退出登录
func (userApi) LogOut(r *ghttp.Request) {
	if err := r.Session.Remove(r.GetSessionId()); err != nil {
		resp.InternalServer(r, "err_seesion_remove", "退出失败", err)
	}
	resp.OK(r)
}

// GetInfo 获取个人信息
func (userApi) GetInfo(r *ghttp.Request) {
	user, err := db.GetUserByID(r.Context(), r.Session.GetString("id"))
	if err != nil {
		if err == mongo.ErrNoDocuments {
			resp.NotFound(r, "account_not_found", "该账号尚未注册")
		}
		resp.DatabaseError(r, "账号查询失败", err)
	}
	resp.OK(r, user)
}
