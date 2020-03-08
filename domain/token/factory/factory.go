package factory

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/mitchellh/mapstructure"
	It "github.com/purwantogz/go-impact/domain/token/interfaces"
	"github.com/purwantogz/go-impact/models"
)

//JwtConfig jwt config secret
type JwtConfig struct {
	SecretKey string
	Issuer    string
	ExpiredIn int
}

//customClaims struct custom claims
type customClaims struct {
	jwt.StandardClaims
	Email string        `json:"email"`
	Roles *models.Roles `json:"roles"`
}

//New for inialisasi jwt token
func New(secretKey, issuer string, expiredIn int) It.ITokenFactory {
	return &JwtConfig{
		SecretKey: secretKey,
		ExpiredIn: expiredIn,
		Issuer:    issuer,
	}
}

//Build Build(roleType string) (map[string]string, error)
func (j *JwtConfig) Build(email string, role *models.Roles) (map[string]string, error) {
	if email != "" {
		resp, _ := j.generateToken(email, role)
		return resp, nil
	}
	return nil, errors.New("email cannot be empty")
}

//Refresh Refresh(refreshToken string) (map[string]string, error)
func (j *JwtConfig) Refresh(refreshToken string) (map[string]string, error) {

	//Parse taken the token string and function for looking up the key
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		//Dont forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(j.SecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		//Get the user record from database or
		//run through your bussiness logic to verifiy if the user can log in
		if string(claims["iss"].(string)) == j.Issuer {
			if string(claims["email"].(string)) != "" {

				result := models.Roles{}
				mapstructure.Decode(claims["roles"], &result)
				newToken, _ := j.generateToken(string(claims["email"].(string)), &result)
				return newToken, nil

			}
			return nil, errors.New("email invalid")
		}
		return nil, errors.New("issuer invalid")
	}
	return nil, errors.New("token invalid")
}

func (j *JwtConfig) generateToken(email string, role *models.Roles) (map[string]string, error) {
	claims := customClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    j.Issuer,
			ExpiresAt: time.Now().Add(time.Minute * time.Duration(j.ExpiredIn)).Unix(),
		},
		Email: email,
		Roles: role,
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	// Generate encoded token and send it as response.
	// The signing string should be secret (a generated UUID works too)
	t, err := token.SignedString([]byte(j.SecretKey))
	if err != nil {
		return nil, err
	}

	refreshToken := jwt.New(jwt.SigningMethodHS256)
	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["sub"] = 1
	rtClaims["exp"] = time.Now().Add(time.Hour * 120).Unix()

	rt, err := refreshToken.SignedString([]byte(j.SecretKey))
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"type":          "bearer",
		"access_token":  t,
		"refresh_token": rt,
	}, nil
}
