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
	"strings"
	"syscall"
	"time"

	"github.com/ArtemVoronov/indefinite-studies-utils/pkg/services/auth"
	"github.com/ArtemVoronov/indefinite-studies-utils/pkg/utils"
	"github.com/gin-gonic/gin"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type FuncSetup func()

type FuncShutdown func()

type FuncRegisterService func(s *grpc.Server)

type FuncVerifyToken func(token string) (*auth.VerificationResult, error)

func StartHTTP(setup FuncSetup, shutdown FuncShutdown, host string, router *gin.Engine) {
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

func StartGRPC(setup FuncSetup, shutdown FuncShutdown, host string, registerServices FuncRegisterService, creds *credentials.TransportCredentials, logger *logrus.Logger) {
	setup()
	defer shutdown()
	lis, err := net.Listen("tcp", host)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption = []grpc.ServerOption{}
	if creds != nil {
		opts = append(opts, grpc.Creds(*creds))
	}
	if logger != nil {
		logrusEntry := logrus.NewEntry(logger)
		// TODO
		// Shared options for the logger, with a custom gRPC code to log level function.
		lorgusOpts := []grpc_logrus.Option{
			// grpc_logrus.WithLevels(customFunc),
		}
		// Make sure that log statements internal to gRPC library are logged using the logrus Logger as well.
		grpc_logrus.ReplaceGrpcLogger(logrusEntry)
		opts = append(opts, grpc_middleware.WithUnaryServerChain(
			grpc_logrus.UnaryServerInterceptor(logrusEntry, lorgusOpts...),
		))
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

func AuthReqired(f FuncVerifyToken) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer") {
			c.JSON(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}

		token := authHeader[len("Bearer "):]
		verificationResult, err := f(token)

		if err != nil {
			c.JSON(http.StatusInternalServerError, "Internal Server Error")
			log.Printf("error during verifying access token: %v\n", err)
			c.Abort()
			return
		}

		if (*verificationResult).IsExpired {
			c.JSON(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}

		c.Next()
	}
}

func NewLogrusLogger() *logrus.Logger {
	logrusLogger := logrus.New()
	logrusLogger.SetFormatter(&logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime: "@timestamp",
			logrus.FieldKeyMsg:  "message",
		},
	})
	logrusLogger.SetLevel(logrus.DebugLevel)
	return logrusLogger
}

func GetDurationInMillseconds(start time.Time) float64 {
	end := time.Now()
	duration := end.Sub(start)
	milliseconds := float64(duration) / float64(time.Millisecond)
	rounded := float64(int(milliseconds*100+.5)) / 100
	return rounded
}

func GetClientIP(c *gin.Context) string {
	// first check the X-Forwarded-For header
	requester := c.Request.Header.Get("X-Forwarded-For")
	// if empty, check the Real-IP header
	if len(requester) == 0 {
		requester = c.Request.Header.Get("X-Real-IP")
	}
	// if the requester is still empty, use the hard-coded address from the socket
	if len(requester) == 0 {
		requester = c.Request.RemoteAddr
	}

	// if requester is a comma delimited list, take the first one
	// (this happens when proxied via elastic load balancer then again through nginx)
	if strings.Contains(requester, ",") {
		requester = strings.Split(requester, ",")[0]
	}

	return requester
}

// GetUserID gets the current_user ID as a string
func GetUserID(c *gin.Context) string {
	userID, exists := c.Get("userID")
	if exists {
		return userID.(string)
	}
	return ""
}

// JSONLogMiddleware logs a gin HTTP request in JSON format, with some additional custom key/values
func JSONLogMiddleware(logger *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()

		// Process Request
		c.Next()

		// Stop timer
		duration := GetDurationInMillseconds(start)

		entry := logger.WithFields(logrus.Fields{
			"client_ip":  GetClientIP(c),
			"duration":   duration,
			"method":     c.Request.Method,
			"path":       c.Request.RequestURI,
			"status":     c.Writer.Status(),
			"user_id":    GetUserID(c),
			"referrer":   c.Request.Referer(),
			"request_id": c.Writer.Header().Get("Request-Id"),
			// "api_version": util.ApiVersion,
		})

		if c.Writer.Status() >= 500 {
			entry.Error(c.Errors.String())
		} else {
			entry.Info("")
		}
	}
}
