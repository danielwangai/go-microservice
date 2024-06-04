package svc

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

var jwtKey = []byte(os.Getenv("SERVICE_USER_JWT_SECRET_TOKEN"))

type JWTClaim struct {
	User *UserServiceResponseType
	jwt.StandardClaims
}

func CheckPasswordHash(password, hash string) bool {
	fmt.Println("CheckPasswordHash: ", hash)
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	fmt.Printf("password hash error: %v", err)
	return err == nil
}

func GenerateJWT(user *UserServiceResponseType) (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &JWTClaim{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	return tokenString, err
}

func GetToken(signedToken string) (*JWTClaim, error) {
	var claim JWTClaim
	token, err := jwt.ParseWithClaims(
		signedToken,
		&claim,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		return nil, errors.New("couldn't parse claims")
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		return nil, errors.New("token expired")
	}

	return &claim, err
}
