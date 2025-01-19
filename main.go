package main

import (
	"database/sql"
	"fmt"
	"log/slog"
	"time"

	"github.com/bwmarrin/snowflake"
	"github.com/chenmuyao/url_shortener/config"
	"github.com/chenmuyao/url_shortener/internal/events"
	"github.com/chenmuyao/url_shortener/internal/repo"
	"github.com/chenmuyao/url_shortener/internal/repo/dao"
	"github.com/chenmuyao/url_shortener/internal/service"
	"github.com/chenmuyao/url_shortener/internal/web"
	"github.com/go-playground/validator/v10"
	"github.com/redis/go-redis/v9"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	slog.SetLogLoggerLevel(slog.LevelDebug)

	// init DB
	dbConn, err := initDB()
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()
	slog.Info("DB init")

	// init Redis
	redis := redis.NewClient(&redis.Options{
		Addr: config.Redis.Addr,
	})

	v := validator.New(validator.WithRequiredStructEnabled())

	node, err := snowflake.NewNode(1)
	if err != nil {
		panic(err)
	}

	dao := dao.New(dbConn)
	counter := events.NewRedisCountProducer(redis)
	repo := repo.NewUrlShortenerRepo(node, dao, counter)
	svc := service.NewUrlShortenerSvc(repo)
	url := web.NewUrlShortenerHdl(v, svc)

	// start consumer
	c := events.NewRedisCountConsumer(dao, redis)
	c.Start()

	// init web server
	app := fiber.New()

	app.Use(pprof.New())

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
