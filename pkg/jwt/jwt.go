package jwt

import (
	"github.com/golang-jwt/jwt/v5"
)

func DecodeTokenString(tokenString string) (jwt.Claims, error) {
	// token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
	// 	// Don't forget to validate the alg is what you expect:
	// 	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
	// 		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	// 	}

	// 	// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
	// 	return []byte("rahasia"), nil
	// })

	// if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
	// 	return claims
	// } else {
	// 	return nil
	// }
	signingKey := []byte("rahasia")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})
	if err != nil {
		return nil, err
	}
	return token.Claims, nil
}
