package server

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/Alonso-Arias/test-boletia/handler"
	"github.com/Alonso-Arias/test-boletia/log"
	ca "github.com/Alonso-Arias/test-boletia/pkg/currency-api"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var loggerf = log.LoggerJSON().WithField("package", "main")

var serverPort = os.Getenv("HOST") + ":" + os.Getenv("PORT")

// Run starts the HTTP server
func Run() {
	log := loggerf.WithField("Server", "Run")

	e := echo.New()
	e.Use(middleware.Logger())

	setUpServer(e)

	// Obtener el intervalo de tiempo desde una variable de entorno
	intervalStr := os.Getenv("INTERVAL_MINUTES")
	intervalMinutes, err := strconv.Atoi(intervalStr)
	if err != nil {
		log.WithError(err).
			WithField("interval_minutes", intervalStr).
			Fatal("failed to parse interval")
	}

	// Crear un ticker que se ejecute cada X minutos
	ticker := time.NewTicker(time.Duration(intervalMinutes) * time.Minute)
	defer ticker.Stop()

	// Ejecutar la función myRoutine cada vez que el ticker se active
	go func() {
		for range ticker.C {
			myRoutine()
		}
	}()

	go func() {
		log.Info("Starting server")

		err := e.Start(serverPort)
		if err != nil {
			log.WithError(err).
				WithField("server_port", serverPort).
				Fatal("failed to start server")
		}
	}()

	// Wait for an interrupt
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)    // interrupt signal sent from terminal
	signal.Notify(sigint, syscall.SIGTERM) // sigterm signal sent from system
	<-sigint

	log.Info("Shutting down server")

	attemptGracefulShutdown(e)
}

func setUpServer(e *echo.Echo) {
	e.GET("/api/v1/currencies/:currency", handler.CurrenciesGet)
	// e.GET("/swagger/*", echoSwagger.WrapHandler)
}

func attemptGracefulShutdown(e *echo.Echo) {
	log := loggerf.WithField("Server", "attemptGracefulShutdown")
	if err := shutdownServer(e, 25*time.Second); err != nil {
		log.WithError(err).Error("failed to shutdown server")
	}
}

func shutdownServer(e *echo.Echo, maximumTime time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), maximumTime)
	defer cancel()
	return e.Shutdown(ctx)
}

func myRoutine() {
	fmt.Println("Ejecutando rutina...")
	ca.FindAndSaveCurrencyValues()
}
