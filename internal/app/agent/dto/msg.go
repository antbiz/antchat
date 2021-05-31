package dto

type SendMsgRequest struct {
	SeqId        string `json:"seq" v:"required"`  // sequence number chosen by client
	Body         []byte `json:"body" v:"required"` // binary body bytes
	SenderID     string
	SenderRole   int
	SenderNick   string
	ReceiverID   string `json:"receiverID" v:"required"`
	ReceiverNick string `json:"receiverNick" v:"required"`
}

type PullMsgRequest struct {
	VisitorID string `json:"visitorID" v:"required"`
}
