package dto

type PullMsgReq struct {
	PageNum  int64 `json:"pageNum" v:"min:1" d:"1"`
	PageSize int64 `json:"pageSize" v:"max:50" d:"20"`
}
