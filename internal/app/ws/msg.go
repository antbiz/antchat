package ws

type Msg struct {
	SeqId        string `json:"seq" v:"required"`  // sequence number chosen by client
	Body         []byte `json:"body" v:"required"` // binary body bytes
	SenderID     string
	SenderRole   int
	SenderNick   string
	ReceiverID   string
	ReceiverNick string
}
