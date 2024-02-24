package security

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"online-store-application/config"
	"online-store-application/model"
	"strconv"
	"time"
)

func GenerateJwtToken(user model.Users) (string, error) {
	_, err := config.NewConfig()
	if err != nil {
		return "", err
	}

	now := time.Now()
	expirationMinutes, err := strconv.Atoi(viper.GetString("APP_EXPIRATION_TOKEN"))
	if err != nil {
		return "", fmt.Errorf("failed to parse expiration time: %v", err)
	}
	end := now.Add(time.Duration(expirationMinutes) * time.Minute)

	claims := &AppClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    viper.GetString("APP_TOKEN_NAME"),
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(end),
		},
		Username: user.Username,
	}

	tokenKey := []byte(viper.GetString("APP_TOKEN_KEY"))
	tokenJwt := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := tokenJwt.SignedString(tokenKey)
	if err != nil {
		return "", fmt.Errorf("failed to create jwt token: %v", err.Error())
	}
	return tokenString, nil
}

func VerifyJwtToken(tokenString string) (jwt.MapClaims, error) {
	cfg, err := config.NewConfig()
	if err != nil {
		return nil, err
	}

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		method, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok || method != &cfg.JwtSigningMethod {
			return nil, fmt.Errorf("invalid token signin method")
		}
		return cfg.JwtSignatureKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("error: %v", err.Error())
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid || claims["iss"] != cfg.ApplicationName {
		return nil, fmt.Errorf("invalid token")
	}
	return claims, nil
}
