package helper

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"ecommerce-microservice/common/interfaces"
	"ecommerce-microservice/common/models"
	"time"
)

func NewHelperFunction() interfaces.HelperInterface {
	return helperImplementation{}
}

type helperImplementation struct{}

func (h helperImplementation) CreateJwtTokenLogin(userID, username string, authID uuid.UUID) (token string, err error) {
	rtClaims := jwt.MapClaims{}
	rtClaims["user_id"] = userID
	rtClaims["username"] = username
	rtClaims["auth_id"] = authID
	rtClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	token, err = rt.SignedString([]byte("secret"))

	if err != nil {
		return "", err
	}
	return
}

func (h helperImplementation) CreateRefreshJwtTokenLogin(userID string, authID uuid.UUID) (token string, err error) {

	refreshToken := jwt.MapClaims{}
	refreshToken["user_id"] = userID
	refreshToken["auth_id"] = authID
	refreshToken["exp"] = time.Now().Add(time.Hour * 24 * 7).Unix()

	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshToken)
	token, err = rt.SignedString([]byte("refresh"))

	if err != nil {
		return "", err
	}
	return
}

func (h helperImplementation) ParseJwt(token string) (claims models.TokenClaims, err error) {
	_, err = jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("secret"), nil
	})
	return
}

func (h helperImplementation) ParseRefreshJwt(token string) (claims models.TokenClaims, err error) {
	_, err = jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("refresh"), nil
	})
	return
}
