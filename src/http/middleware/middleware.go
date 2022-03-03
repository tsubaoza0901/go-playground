package middleware

import (
	"fmt"
	"go-playground/m/v1/src/http/handler"
	"log"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitMiddleware(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
		AllowHeaders:     []string{echo.HeaderAccessControlAllowHeaders, echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowCredentials: true,
	}))
}

func NewJwtConfig() echo.MiddlewareFunc {
	config := middleware.JWTConfig{
		Claims: &handler.JwtCustomClaims{},
		// SigningKey: []byte("secret"),
		KeyFunc: getKey,
		// ParseTokenFunc: parseToken,
	}
	return middleware.JWTWithConfig(config)
}

func getKey(token *jwt.Token) (interface{}, error) {
	log.Printf("token: %+v", token)
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}

	return []byte("secret"), nil
}

// func parseToken(auth string, c echo.Context) (interface{}, error) {
// 	log.Printf("auth: %v", auth)
// 	token, err := jwt.Parse(auth, getKey)
// 	if err != nil {
// 		return nil, err
// 	}

// 	log.Printf("token: %+v", token)
// 	if !token.Valid {
// 		return nil, errors.New("invalid token")
// 	}
// claims, ok := token.Claims.(*handler.JwtCustomClaims)
// if !ok {
// 	return nil, errors.New("cann't convert token.Claims")
// }

// nbf := time.Unix(claims.NotBefore, 0)
// log.Printf("nbf: %v", nbf)

// if time.Now().Before(nbf) {
// 	log.Println("invalid token")
// 	return nil, errors.New("Before")
// }

// exp := time.Unix(claims.ExpiresAt, 0)
// log.Printf("exp: %v", exp)

// if time.Now().After(exp) {
// 	log.Println("expired token")
// 	return nil, errors.New("it's expired")
// }

// 	return token, nil
// }
