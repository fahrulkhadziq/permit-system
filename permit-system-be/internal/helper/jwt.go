package helper

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(userID, role, unitID string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"unit_id": unitID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func GenerateAccessToken(
	userID,
	role,
	unitID string,
) (string, error) {

	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"unit_id": unitID,
		"exp": time.Now().
			Add(time.Minute * 15).
			Unix(),
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	return token.SignedString(
		[]byte(os.Getenv("JWT_SECRET")),
	)
}

func GenerateRefreshToken(
	userID string,
) (string, error) {

	claims := jwt.MapClaims{
		"user_id": userID,
		"type":    "refresh",
		"exp": time.Now().
			Add(time.Hour * 24 * 7).
			Unix(),
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	return token.SignedString(
		[]byte(os.Getenv("JWT_SECRET")),
	)
}
