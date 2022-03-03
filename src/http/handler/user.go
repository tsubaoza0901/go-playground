package handler

import (
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type User struct {
	ID   uint
	Name string
	Age  uint
}

type JwtCustomClaims struct {
	UserID uint `json:"user_id"`
	Admin  bool `json:"admin"`
	jwt.StandardClaims
}

func Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	// Throws unauthorized error
	if username != "jon" || password != "shhh!" {
		log.Printf("username: %s / password: %s\n", username, password)
		return echo.ErrUnauthorized
	}

	// Set custom claims
	claims := &JwtCustomClaims{
		UserID: 1,
		Admin:  true,
		StandardClaims: jwt.StandardClaims{ // StandardClaimsを使用している場合、IssuedAt、NotBefore、ExpiresAtについては各フィールドに対するデフォルトのバリデーションが設定されている
			Id:       "000001",          // JWT ID：JWTのユニーク性を担保するためのID
			Issuer:   "xxx.co.jp",       // Issuer：JWTの発行者
			Subject:  "check user",      // Subject：JSTの使用用途
			Audience: "service user",    // Audience：JWTの想定利用者
			IssuedAt: time.Now().Unix(), // Issued At：数値表現によるJWT発行日
			// NotBefore: time.Now().Add(time.Minute * 5).Unix(), // Not Before：数値表現によるJWTの利用開始日時（例の場合、JWTが発行された5min後から有効）
			ExpiresAt: time.Now().Add(time.Hour).Unix(), // Expiration Time：数値表現によるJWT失効日
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}

func Private(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*JwtCustomClaims)
	id := claims.UserID
	user := getUser(id)
	return c.JSON(http.StatusOK, user)
}

func getUser(userID uint) User {
	return User{
		ID:   userID,
		Name: "山田",
		Age:  30,
	}
}

// // Custom ParseTokenFuncを使用する場合
// func Private(c echo.Context) error {
// 	token := c.Get("user").(*jwt.Token)
// 	claims := token.Claims.(jwt.MapClaims)
// 	id := claims["user_id"]
// 	user := getUser(id.(float64))
// 	return c.JSON(http.StatusOK, user)
// }

// func getUser(userID float64) User {
// 	return User{
// 		ID:   uint(userID),
// 		Name: "山田",
// 		Age:  30,
// 	}
// }

func Logout(c echo.Context) error {
	return c.Redirect(http.StatusFound, "/")
}
