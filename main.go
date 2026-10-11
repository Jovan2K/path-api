package main

import (
	"context"
	"log"
	"path-api/internal/auth"
	"path-api/internal/db"
	"path-api/internal/kelas"
	"path-api/internal/pengguna"
	"path-api/internal/postingan"
	"path-api/internal/utils"

	_ "github.com/joho/godotenv/autoload"
	echojwt "github.com/labstack/echo-jwt/v5"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

// @title PATH API
// @version 1.0
// @description Personal Adaptive Teaching Hub

// @contact.name MBKM PATH
// @contact.email mbkmpath@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	ctx := context.Background()
	q, pool, dbConn, err := db.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	defer dbConn.Close()
	defer pool.Close()

	e := echo.New()

	e.Validator = utils.NewValidator()

	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())

	authHandler := auth.New(q)
	authHandler.Routes(e)

	protected := e.Group("")
	protected.Use(echojwt.WithConfig(utils.JWTConfig))

	penggunaHandler := pengguna.New(q)
	penggunaHandler.Routes(protected)

	kelasHandler := kelas.New(q)
	kelasHandler.Routes(protected)

	postinganHandler := postingan.New(q)
	postinganHandler.Routes(protected)

	port := utils.GetPort()
	err = e.Start(port)
}
