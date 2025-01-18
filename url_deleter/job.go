package main

import (
	"context"
	"log/slog"
	"time"
)

func (a *Application) DeleteJob() {
	ctx, cancel := context.WithTimeout(context.Background(), a.Config.JobTimeout)
	defer cancel()

	cutTime := time.Now().Add(-a.Config.DataTimeout)

	res, err := a.DBConn.ExecContext(ctx, `DELETE FROM "urls" WHERE created_at < $1`, cutTime)
	if err != nil {
		slog.Error("Failed to delete expired lines", slog.Any("err", err))
	}
	rows, err := res.RowsAffected()
	slog.Info("Removed lines", slog.Any("rows", rows), slog.Any("err", err))
}
