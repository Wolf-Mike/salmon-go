package sgcjwt

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//EncryToken 创建jwt-token
func EncryToken(user TokenUserInput, secretKey string) (r string, err error) {
	claims := SgcClaims{
		ID:       user.ID,
		Username: user.Username,
		Password: user.Password,
		StandardClaims: jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix()),               // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + user.Expire), // 过期时间 一小时

		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	r, err = token.SignedString([]byte(secretKey))
	return
}

//DecrypToken 解析jwt-token
func DecrypToken(token string, secretKey string) (claims *SgcClaims, err error) {
	t, err := jwt.ParseWithClaims(token, &SgcClaims{}, func(jt *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				err = errors.New("无效token")
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				err = errors.New("token已过期")
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				err = errors.New("token未激活") //还没有到达token开始时间，一般用不到
			} else {
				err = errors.New("token未激活")
			}
			return
		}
	}
	claims, ok := t.Claims.(*SgcClaims)
	if !ok || !t.Valid {
		err = errors.New("无效token")
	}
	return
}
