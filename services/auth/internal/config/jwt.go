package config

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

type JwtConfig struct {
	config *viper.Viper
}

func NewJWT(viperConfig *viper.Viper) *JwtConfig {
	return &JwtConfig{
		config: viperConfig,
	}
}

func (j *JwtConfig) ClaimAccessToken(email string) (string, error) {
	claims := jwt.MapClaims{
		"kid":   "sim2",
		"email": email,
		"exp":   time.Now().Add(time.Second * time.Duration(3600)).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token.Header["kid"] = "sim2"
	// Generate encoded token and send it as response.
	fmt.Println("kid ", j.config.GetString("jwt.kid"))
	t, err := token.SignedString([]byte(j.config.GetString("jwt.key")))
	if err != nil {
		return "", err
	}
	return t, nil
}

func (j *JwtConfig) ClaimRefreshToken(email string) (string, error) {
	claims := jwt.MapClaims{
		"key":   "rahasia",
		"email": email,
		"exp":   time.Now().Add(time.Hour * 7200).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}
	return t, nil
}

func (j *JwtConfig) DecodeToken(token *jwt.Token) jwt.MapClaims {
	claims := token.Claims.(jwt.MapClaims)
	return claims
}

/*
*
decode token from token string

param string
return jwt.MapClaims
*/
func (j *JwtConfig) DecodeTokenString(tokenString string) jwt.MapClaims {
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte("rahasia"), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims
	} else {
		return nil
	}
}

// verify jwt token by token string in header or payload
func VerifyToken(tokenString string) (*jwt.Token, bool) {
	token, _ := jwt.Parse("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFkaXRAbWFpbC5jb20iLCJleHAiOjE3MjI4NTA4MjcsImtleSI6InJhaGFzaWEifQ.GGgqjc4z_lQUUedAVeYQz4tjh4WrG_QAt5cpRjvB-7s", func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte("rahasia"), nil
	})

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return token, true
	} else {
		return nil, false
	}
}
