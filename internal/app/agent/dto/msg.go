package dto

import "github.com/antbiz/antchat/internal/app/ws"

type SendMsgReq struct {
	ws.ChatMsg
	SenderID     string
	SenderNick   string
	ReceiverID   string `json:"receiverID" v:"required"`
	ReceiverNick string `json:"receiverNick" v:"required"`
}

type PullMsgReq struct {
	VisitorID string `json:"visitorID" v:"required"`
	PageNum   int64  `json:"pageNum" v:"min:1" d:"1"`
	PageSize  int64  `json:"pageSize" v:"max:50" d:"20"`
}
