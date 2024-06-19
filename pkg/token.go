package pkg

import (
	"fmt"
	"os"
	"time"

	"github.com/clerijr/teste-picpay-go/entities/user/dto"
	"github.com/go-chi/jwtauth"
	"github.com/golang-jwt/jwt"
)

type UserAuthToken struct {
	AccessToken string `json:"access_token"`
}

type AuthEncoder struct {
	secret string
}

func NewAuthEncoder() *AuthEncoder {
	return &AuthEncoder{
		secret: os.Getenv("SECRET"),
	}
}

func (a *AuthEncoder) GenerateToken(user *dto.UserAuth) (*UserAuthToken, error) {

	tokenAuth := jwtauth.New("HS256", []byte(os.Getenv("SECRET")), nil)

	_, token, _ := tokenAuth.Encode(map[string]any{
		"id":       user.ID,
		"username": user.Name,
		"email":    user.Email,
		"exp":      time.Now().Add(time.Second * 60000).Unix(),
	})

	userAuthToken := UserAuthToken{
		AccessToken: token,
	}

	return &userAuthToken, nil
}

func (a *AuthEncoder) ParseJWTToken(tkn string) (*dto.UserAuth, error) {

	var userClaims dto.UserAuth

	_, err := jwt.ParseWithClaims(tkn, &userClaims, func(token *jwt.Token) (interface{}, error) {
		return []byte(a.secret), nil
	})
	if err != nil {
		err = fmt.Errorf("error decoding token: %v", err)
		return nil, err
	}

	return &userClaims, nil
}
