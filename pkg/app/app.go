package app

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ArtemVoronov/indefinite-studies-utils/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type FuncSetup func()
type FuncShutdown func()

func StartHTTP(setup FuncSetup, shutdown FuncShutdown, host string, router *gin.Engine) {
	LoadEnv()
	setup()
	defer shutdown()
	srv := &http.Server{
		Addr:    host,
		Handler: router,
	}

	go func() {
		log.Printf("http server listening at %v\n", srv.Addr)
		err := srv.ListenAndServe()
		if err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Println("http server was closed")
		} else if err != nil {
			log.Fatalf("unable to start http server: %v\n", err)
		}
	}()

	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be caught, so don't need to add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("shutting down http server ...")

	ctx, cancel := context.WithTimeout(context.Background(), ShutdownTimeout())
	defer cancel()

	err := srv.Shutdown(ctx)
	if err != nil {
		log.Fatal("http server forced to shutdown:", err)
	}

	log.Println("http server has been shutdown")
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

func HostHTTP() string {
	port := utils.EnvVarDefault("APP_HTTP_API_PORT", "3005")
	host := ":" + port
	return host
}

func HostGRPC() string {
	port := utils.EnvVarDefault("APP_GRPC_API_PORT", "50051")
	host := ":" + port
	return host
}

func Mode() string {
	return utils.EnvVarDefault("APP_MODE", "debug")
}

func ShutdownTimeout() time.Duration {
	return utils.EnvVarDurationDefault("APP_SHUTDOWN_TIMEOUT_IN_SECONDS", time.Second, 5*time.Second)
}

func TLSCredentials() credentials.TransportCredentials {
	creds, err := LoadTLSCredentialsForServer(utils.EnvVar("APP_TLS_CERT_PATH"), utils.EnvVar("APP_TLS_KEY_PATH"))
	if err != nil {
		log.Fatalf("unable to load TLS credentials")
	}
	return creds
}

type FuncRegisterService func(s *grpc.Server)

func StartGRPC(setup FuncSetup, shutdown FuncShutdown, host string, registerServices FuncRegisterService, creds *credentials.TransportCredentials) {
	setup()
	defer shutdown()
	lis, err := net.Listen("tcp", host)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	if creds != nil {
		opts = []grpc.ServerOption{grpc.Creds(*creds)}
	}

	grpc := grpc.NewServer(opts...)

	registerServices(grpc)

	go func() {
		log.Printf("grpc server listening at %v", lis.Addr())
		err := grpc.Serve(lis)
		if err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be caught, so don't need to add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("grpc server shutting down server ...")

	grpc.GracefulStop()

	log.Println("grpc server has been shutdown")
}

func LoadTLSCredentialsForServer(certPath, keyPath string) (credentials.TransportCredentials, error) {
	serverCert, err := tls.LoadX509KeyPair(certPath, keyPath)
	if err != nil {
		return nil, err
	}

	config := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.NoClientCert,
	}

	return credentials.NewTLS(config), nil
}

func LoadTLSCredentialsForClient(certPath string) (credentials.TransportCredentials, error) {
	pemServerCA, err := ioutil.ReadFile(certPath)
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemServerCA) {
		return nil, fmt.Errorf("failed to add server CA's certificate")
	}

	config := &tls.Config{
		RootCAs: certPool,
	}

	return credentials.NewTLS(config), nil
}
