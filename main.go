package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/mattn/go-sqlite3"

	"gitdev.devops.krungthai.com/techcoach/template/goapi.git/app"
	"gitdev.devops.krungthai.com/techcoach/template/goapi.git/app/course"
	"gitdev.devops.krungthai.com/techcoach/template/goapi.git/database"
	"gitdev.devops.krungthai.com/techcoach/template/goapi.git/logger"
)

func main() {
	db := database.NewSQLite()
	defer func() {
		_ = db.Close()
	}()

	// logger, graceful := logger.NewZap()
	// defer graceful()
	logger := logger.New()

	r := app.NewRouter(logger)

	// r.StaticFS("/public", http.Dir("static"))

	storage := course.NewStorage(db)
	courses := course.New(storage)

	r.POST("r/courses/recommends", courses.Recommended)
	r.POST("courses", courses.NewCourse)

	srv := http.Server{
		Addr:              ":" + os.Getenv("PORT"),
		Handler:           r,
		ReadHeaderTimeout: 5 * time.Second,
	}

	idleConnsClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, syscall.SIGINT, syscall.SIGTERM)
		<-sigint

		d := time.Duration(5 * time.Second)
		fmt.Printf("shutting down int %s ...", d)
		// We received an interrupt signal, shut down.
		ctx, cancel := context.WithTimeout(context.Background(), d)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			// Error from closing listeners, or context timeout:
			logger.Info("HTTP server Shutdown: " + err.Error())
		}
		close(idleConnsClosed)
	}()

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		logger.Error("HTTP server ListenAndServe: " + err.Error())
		return
	}

	<-idleConnsClosed
	fmt.Println("gracefully")
}
