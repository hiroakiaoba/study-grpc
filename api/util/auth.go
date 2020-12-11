package util

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Auther interface {
	GenToken(userID int32) (token string, err error)
	Verify(token string) (userID int32, err error)
}

type JWTAuth struct{}

func NewJWTAuth() Auther {
	return &JWTAuth{}
}

const (
	userIDKey = "user_id"
	// iat と exp は登録済みクレーム名。それぞれの意味は https://tools.ietf.org/html/rfc7519#section-4.1 を参照
	iatKey = "iat"
	expKey = "exp"
	// lifetime は jwt の発行から失効までの期間を表す。
	lifetime = 72 * time.Hour
)

// var signingkey = os.Getenv("SIGNINGKEY")
var signingkey = "sample_key"

// GenToken userIDを受け取ってJWTを作成する関数
func (j *JWTAuth) GenToken(userID int32) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		userIDKey: userID,
		iatKey:    time.Now().Unix(),
		expKey:    time.Now().Add(lifetime).Unix(),
	})

	return token.SignedString([]byte(signingkey))
}

// Verify Tokenを受け取って検証してOKならuserIDを返す関数
func (j *JWTAuth) Verify(token string) (int32, error) {
	decodedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return "", fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(signingkey), nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := decodedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, err
	}
	userID, ok := claims["user_id"].(float64)
	if !ok {
		return 0, errors.New("type cast error. userId to float64")
	}

	return int32(userID), nil
}
