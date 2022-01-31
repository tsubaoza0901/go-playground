package main

import (
	"log"
	"net/http"
	"net/http/pprof"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// --------
// model↓
// --------

// User user info
type User struct {
	ID          string     `gorm:"primary_key;not null" json:"id_str"`
	ScreenName  string     `gorm:"not null" json:"screen_name"`
	Name        string     `gorm:"not null" json:"name"`
	URL         string     `gorm:"not null" json:"url"`
	Description string     `gorm:"null" json:"description"`
	IsSignedIn  bool       `gorm:"not null" json:"is_signed_in"`
	CreatedAt   time.Time  `gorm:"null" json:"create_at"`
	UpdatedAt   time.Time  `gorm:"null" json:"update_at"`
	DeletedAt   *time.Time `gorm:"null" json:"-"`
}

// --------
// router↓
// --------

// InitRouting ...
func InitRouting(e *echo.Echo) {
	pf := e.Group("/debug/pprof/")
	pf.Any("/", echo.WrapHandler(http.HandlerFunc(pprof.Index)))
	pf.Any("/cmdline", echo.WrapHandler(http.HandlerFunc(pprof.Cmdline)))
	pf.Any("/profile", echo.WrapHandler(http.HandlerFunc(pprof.Profile)))
	pf.Any("/symbol", echo.WrapHandler(http.HandlerFunc(pprof.Symbol)))
	pf.Any("/heap", echo.WrapHandler(http.HandlerFunc(pprof.Handler("heap").ServeHTTP)))

	e.GET("/", Top)
	e.GET("/logout", Logout)
}

func Top(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hello, World!")
}

func Logout(c echo.Context) error {
	return c.Redirect(http.StatusFound, "/")
}

// // --------
// // db↓
// // --------

// var db *gorm.DB

// // InitDB ...
// func InitDB() *gorm.DB {
// 	dsn := "root:root@tcp(db:3306)/goplayground?parseTime=True&loc=Local"
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return db
// }

// --------
// middleware↓
// --------

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

// --------
// main↓
// --------

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	// db = InitDB()

	// sqlDB, err := db.DB()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer sqlDB.Close()

	e := echo.New()

	InitMiddleware(e)

	InitRouting(e)

	log.Println(e.Start(":8444"))
}
