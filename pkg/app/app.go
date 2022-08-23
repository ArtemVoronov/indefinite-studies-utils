package app

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ArtemVoronov/indefinite-studies-utils/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type FuncSetup func()
type FuncShutdown func()

func Start(setup FuncSetup, shutdown FuncShutdown, host string, router *gin.Engine) {
	setup()
	defer shutdown()
	srv := &http.Server{
		Addr:    host,
		Handler: router,
	}

	go func() {
		log.Printf("App starting at localhost%s ...\n", srv.Addr)
		err := srv.ListenAndServe()
		if err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Println("Server was closed")
		} else if err != nil {
			log.Fatalf("Unable to start app: %v\n", err)
		}
	}()

	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be caught, so don't need to add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server ...")

	ctx, cancel := context.WithTimeout(context.Background(), ShutdownTimeout())
	defer cancel()

	err := srv.Shutdown(ctx)
	if err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server has been shutdown")
}

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Print("No .env file found")
	}
}

func Cors() gin.HandlerFunc {
	cors := utils.EnvVarDefault("CORS", "*")
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", cors)
		c.Next()
	}
}

func Host() string {
	port := utils.EnvVarDefault("APP_PORT", "3005")
	host := ":" + port
	return host
}

func Mode() string {
	return utils.EnvVarDefault("APP_MODE", "debug")
}

func ShutdownTimeout() time.Duration {
	return utils.EnvVarDurationDefault("APP_SHUTDOWN_TIMEOUT_IN_SECONDS", time.Second, 5*time.Second)
}
