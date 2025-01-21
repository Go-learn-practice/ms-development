package jwts

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type JwtToken struct {
	AccessToken  string
	RefreshToken string
	AccessExp    int64
	RefreshExp   int64
}

func CreateToken(val string, exp time.Duration, secret string, refreshExp time.Duration, refreshSecret string) *JwtToken {
	aExp := time.Now().Add(exp).Unix()
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"token": val,
		"exp":   aExp,
	})
	aToken, _ := accessToken.SignedString([]byte(secret))
	rExp := time.Now().Add(refreshExp).Unix()
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"token": val,
		"exp":   rExp,
	})
	rToken, _ := refreshToken.SignedString([]byte(refreshSecret))

	return &JwtToken{
		AccessToken:  aToken,
		RefreshToken: rToken,
		AccessExp:    aExp,
		RefreshExp:   rExp,
	}
}

func ParseToken(tokenString string, secret string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secret), nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// interface断言
		val := claims["token"].(string)
		exp := int64(claims["exp"].(float64))
		if time.Now().Unix() > exp {
			return "", fmt.Errorf("token已过期")
		}
		return val, nil
	} else {
		return "", fmt.Errorf("token无效")
	}
}
