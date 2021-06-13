package dto

// PullMsgReq 拉取指定访客的历史消息
type PullMsgReq struct {
	VisitorID string `json:"visitorID" v:"required"`
	PageNum   int64  `json:"pageNum" v:"min:1" d:"1"`
	PageSize  int64  `json:"pageSize" v:"max:50" d:"20"`
}
