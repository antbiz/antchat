package api

import (
	"github.com/antbiz/antchat/internal/app/system/dto"
	"github.com/antbiz/antchat/internal/db"
	"github.com/antbiz/antchat/internal/pkg/resp"
	"github.com/gogf/gf/net/ghttp"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Visitor 访客管理接口
var Visitor = &visitorApi{}

type visitorApi struct{}

// Get 获取访客信息
func (visitorApi) Get(r *ghttp.Request) {
	vid := r.GetString("id")
	visitor, err := db.GetVisitorByID(r.Context(), vid)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			resp.NotFound(r, "not_found_visitor", "该访客不存在")
		}
		resp.DatabaseError(r, "查询访客详情失败", err)
	}
	resp.OK(r, visitor)
}

// Update 更新访客信息
func (visitorApi) Update(r *ghttp.Request) {
	var req *dto.VisitorUpdateReq
	if err := r.Parse(&req); err != nil {
		resp.InvalidArgument(r, err.Error())
	}
	vid, _ := primitive.ObjectIDFromHex(r.GetString("id"))
	if err := db.GetVisitorCollection().
		UpdateId(
			r.Context(),
			vid,
			bson.M{"$set": bson.M{
				"nickname": req.Nickname,
				"phone":    req.Phone,
				"email":    req.Email,
			}},
		); err != nil {
		resp.DatabaseError(r, "更新访客信息失败", err)
	}
	resp.OK(r)
}
