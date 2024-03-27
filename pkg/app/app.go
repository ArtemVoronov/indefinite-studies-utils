package app

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/ArtemVoronov/indefinite-studies-utils/pkg/log"
	"github.com/ArtemVoronov/indefinite-studies-utils/pkg/services/auth"
	"github.com/ArtemVoronov/indefinite-studies-utils/pkg/services/db/entities"
	"github.com/ArtemVoronov/indefinite-studies-utils/pkg/utils"
	"github.com/gin-gonic/gin"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	CTX_TOKEN_ID_KEY   = "Id"
	CTX_TOKEN_TYPE_KEY = "Type"
	CTX_TOKEN_ROLE_KEY = "Role"
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
		log.Info(fmt.Sprintf("http server listening at %v\n", srv.Addr))
		err := srv.ListenAndServe()
		if err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Info("http server was closed")
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

	log.Info("shutting down http server ...")

	ctx, cancel := context.WithTimeout(context.Background(), ShutdownTimeout())
	defer cancel()

	err := srv.Shutdown(ctx)
	if err != nil {
		log.Fatalf("http server forced to shutdown: %s", err)
	}

	log.Info("http server has been shutdown")
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
		log.Info(fmt.Sprintf("grpc server listening at %v", lis.Addr()))
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

	log.Info("grpc server shutting down server ...")

	grpc.GracefulStop()

	log.Info("grpc server has been shutdown")
}

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Error("Unable to load environment vars", "No .env file found")
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
		log.Fatalf("unable to load TLS credentials: %s", err.Error())
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
	pemServerCA, err := os.ReadFile(certPath)
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
			log.Error("error during verifying access token", err.Error())
			c.Abort()
			return
		}

		if (*verificationResult).IsExpired {
			c.JSON(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}

		c.Set(CTX_TOKEN_ID_KEY, verificationResult.Uuid)
		c.Set(CTX_TOKEN_TYPE_KEY, verificationResult.Type)
		c.Set(CTX_TOKEN_ROLE_KEY, verificationResult.Role)

		c.Next()
	}
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

func NewLoggerMiddleware(logger *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()

		duration := GetDurationInMillseconds(start)

		entry := logger.WithFields(logrus.Fields{
			"client_ip":  GetClientIP(c),
			"duration":   duration,
			"method":     c.Request.Method,
			"path":       c.Request.RequestURI,
			"status":     c.Writer.Status(),
			"user_id":    c.GetInt(CTX_TOKEN_ID_KEY),
			"user_type":  c.GetString(CTX_TOKEN_TYPE_KEY),
			"user_role":  c.GetString(CTX_TOKEN_ROLE_KEY),
			"referrer":   c.Request.Referer(),
			"request_id": c.Writer.Header().Get("Request-Id"),
		})

		if c.Writer.Status() >= 500 {
			entry.Error(c.Errors.String())
		} else {
			entry.Info("")
		}
	}
}

func RequiredRoles(reqiredRoles []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role := c.GetString(CTX_TOKEN_ROLE_KEY)

		if !utils.Contains(reqiredRoles, role) {
			c.JSON(http.StatusForbidden, "Forbidden")
			c.Abort()
			return
		}

		c.Next()
	}
}

func RequiredOwnerRole() gin.HandlerFunc {
	return func(c *gin.Context) {
		role := c.GetString(CTX_TOKEN_ROLE_KEY)

		if role != entities.USER_ROLE_OWNER {
			c.JSON(http.StatusForbidden, "Forbidden")
			c.Abort()
			return
		}

		c.Next()
	}
}

func IsTokenOfUserType(c *gin.Context) bool {
	tokenType := c.GetString(CTX_TOKEN_TYPE_KEY)
	return tokenType == entities.TOKEN_TYPE_USER
}

func IsSameUser(c *gin.Context, userUuid string) bool {
	userUuidFromCtx, ok := c.Get(CTX_TOKEN_ID_KEY)
	if !ok {
		return false
	}

	return userUuidFromCtx == userUuid
}

func HasOwnerRole(c *gin.Context) bool {
	return HasRole(c, entities.USER_ROLE_OWNER)
}

func HasRole(c *gin.Context, role string) bool {
	userRoleFromCtx := c.GetString(CTX_TOKEN_ROLE_KEY)

	return userRoleFromCtx == role
}
