package main

import (
	"database/sql"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/chenmuyao/url_shortener/url_deleter/config"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/robfig/cron/v3"
)

type Application struct {
	DBConn *sql.DB
	Config config.AppConfig
}

func NewApp(c config.AppConfig) *Application {
	return &Application{
		Config: c,
	}
}

func main() {
	app := NewApp(config.App)
	err := app.initDB()
	if err != nil {
		panic(err)
	}
	defer app.DBConn.Close()

	slog.Info("Start the deleter")
	c := cron.New(cron.WithSeconds())

	id, err := c.AddFunc(app.Config.Schedule, app.DeleteJob)
	if err != nil {
		panic(err)
	}

	c.Start()

	slog.Info("job started", slog.Any("id", id))

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
	<-sigchan

	slog.Info("Stop the deleter, wait for the last job to terminate")
	ctx := c.Stop()

	<-ctx.Done()
	slog.Info("Stopped")
}

func (a *Application) initDB() error {
	conn, err := sql.Open("pgx", a.Config.DBAddr)
	if err != nil {
		return fmt.Errorf("Unable to connect to DB: %s", err)
	}

	err = conn.Ping()
	if err != nil {
		return err
	}

	a.DBConn = conn

	return nil
}
