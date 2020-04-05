package common

import (
	"github.com/dgrijalva/jwt-go"
	"time"
	"ziogie.top/gin/model"
)

var jwtKey = []byte("a_secret_crect")

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

func ReleaseToken(user model.User) (string, error)  {
	expirationTime := time.Now().Add(7*24*time.Hour)
	//荷载部分
	claims := &Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt: time.Now().Unix(),
			Issuer: "ziogie.top",
			Subject: "user token",
		},
	}
	//协议头hs256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//前两部分加key的hash
	tokenString,err := token.SignedString(jwtKey)
	if err != nil {
		return "",err
	}
	return tokenString,nil
}

func ParseToken(tokenString string)  (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	
	token,err := jwt.ParseWithClaims(tokenString,claims, func(token *jwt.Token) (i interface{}, e error) {
		return jwtKey,nil
	})
	return token, claims, err
}