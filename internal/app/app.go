package app

import (
	"context"
	"net/http"
	"time"
)

type App struct {
	httpServer *http.Server
}

func (app *App) Run(port string) error {
	app.httpServer = &http.Server{
		Addr: ":" + port,
		MaxHeaderBytes: 1 << 20, // 1 Mb
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	return app.httpServer.ListenAndServe()
}

func (app *App)  Shutdown(ctx context.Context) error {
	return app.httpServer.Shutdown(ctx)
}