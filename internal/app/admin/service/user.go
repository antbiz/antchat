package service

import "github.com/gogf/gf/crypto/gmd5"

var User = &userSrv{}

type userSrv struct{}

// EncryptPwd 加密账号密码
func (srv *userSrv) EncryptPwd(username, password string) string {
	return gmd5.MustEncrypt(username + password)
}
