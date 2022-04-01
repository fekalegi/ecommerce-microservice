package interfaces

import (
	"ecommerce-microservice/order/common/models"
	"github.com/google/uuid"
)

type HelperInterface interface {
	CreateJwtTokenLogin(userID, username string, authID uuid.UUID) (token string, err error)
	CreateRefreshJwtTokenLogin(userID string, authID uuid.UUID) (token string, err error)
	ParseJwt(token string) (claims models.TokenClaims, err error)
	ParseRefreshJwt(token string) (claims models.TokenClaims, err error)
}
