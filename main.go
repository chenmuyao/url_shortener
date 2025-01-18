package main

import (
	"database/sql"
	"fmt"
	"log/slog"
	"time"

	"github.com/chenmuyao/url_shortener/config"
	"github.com/chenmuyao/url_shortener/internal/repo"
	"github.com/chenmuyao/url_shortener/internal/service"
	"github.com/chenmuyao/url_shortener/internal/web"
	"github.com/go-playground/validator/v10"

	"github.com/gofiber/fiber/v2"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	slog.SetLogLoggerLevel(slog.LevelDebug)
	// init DB
	_, err := initDB()
	if err != nil {
		panic(err)
	}
	slog.Info("DB init")

	v := validator.New(validator.WithRequiredStructEnabled())

	repo := repo.NewUrlShortenerRepo()
	svc := service.NewUrlShortenerSvc(repo)
	url := web.NewUrlShortenerHdl(v, svc)

	// init web server
	app := fiber.New()

	url.RegisterHandlers(app)

	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.SendString("Hello, World!")
	// })

	app.Listen(":3000")
}

func initDB() (*sql.DB, error) {
	conn, err := sql.Open("pgx", config.DB.Addr)
	if err != nil {
		return nil, fmt.Errorf("Unable to connect to DB: %s", err)
	}
	conn.SetMaxOpenConns(config.DB.MaxOpenDbConn)
	conn.SetMaxIdleConns(config.DB.MaxIdleDbConn)
	conn.SetConnMaxLifetime(time.Duration(config.DB.MaxDbLifetimeSec) * time.Second)

	err = conn.Ping()
	if err != nil {
		return nil, err
	}

	return conn, nil
}
