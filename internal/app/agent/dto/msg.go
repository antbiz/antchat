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
}
