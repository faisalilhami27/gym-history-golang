package helpers

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

func secretKey() []byte {
	env := godotenv.Load()
	if env != nil {
		log.Fatal("Error loading .env file")
	}

	secret := os.Getenv("JWT_SECRET")
	jwtSecret := []byte(secret)
	return jwtSecret
}

func LoadEnv() error {
	env := godotenv.Load()
	if env != nil {
		log.Fatal("Error loading .env file")
	}

	return env
}

func FormatValidationError(err error) []string {
	var errors []string

	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}

	return errors
}

func GenerateToken(userId string) (string, error) {
	jwtSecret := secretKey()
	claim := jwt.MapClaims{}
	claim["user_id"] = userId
	claim["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return signedToken, err
	}

	return signedToken, err
}

func ValidateToken(token string) (*jwt.Token, error) {
	validateToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("invalid token")
		}

		jwtSecret := secretKey()
		return jwtSecret, nil
	})

	if err != nil {
		return validateToken, err
	}

	return validateToken, err
}
