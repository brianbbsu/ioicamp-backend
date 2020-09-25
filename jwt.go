package main

// Modified from https://medium.com/%E4%BC%81%E9%B5%9D%E4%B9%9F%E6%87%82%E7%A8%8B%E5%BC%8F%E8%A8%AD%E8%A8%88/golang-json-web-tokens-jwt-olang-json-web-tokens-jwt-%E7%A4%BA%E7%AF%84-225b377e0f79

import (
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Claims contains the claims we want to store for each logined user
type Claims struct {
	UID uint `json:"uid"`
	jwt.StandardClaims
}

func getJWTTokenByUID(uid uint) (string, error) {
	now := time.Now()
	claims := Claims{
		UID: uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: now.Add(time.Duration(Config.GetFloat64("jwt.tokenEffectiveMinutes")) * time.Minute).Unix(),
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString([]byte(Config.GetString("jwt.secret")))
	return token, err
}

func authWithJWT(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	token := strings.TrimPrefix(auth, "Bearer ")
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(Config.GetString("jwt.secret")), nil
	})
	if err != nil {
		var message string
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				message = "token is malformed"
			} else if ve.Errors&jwt.ValidationErrorUnverifiable != 0 {
				message = "token could not be verified because of signing problems"
			} else if ve.Errors&jwt.ValidationErrorSignatureInvalid != 0 {
				message = "signature validation failed"
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				message = "token is expired"
			} else {
				message = "can not handle this token"
			}
		}
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": "failed",
			"error":  message,
		})
		c.Abort()
		return
	}
	if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
		c.Set("UID", claims.UID)
		c.Next()
	} else {
		c.Abort()
		return
	}
}
