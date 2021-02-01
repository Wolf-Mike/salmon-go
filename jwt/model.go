package sgcjwt

import (
	"github.com/dgrijalva/jwt-go"
)

//TokenUserInput 生成token的输入参数
type TokenUserInput struct {
	ID       int64  `json:"Id"`
	Username string `json:"user"`
	Password string `json:"token"`
	//按秒计算
	Expire int64 `json:"expire"`
}

//SgcClaims 继承jwt.StandardClaims
type SgcClaims struct {
	ID       int64  `json:"Id"`
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}
