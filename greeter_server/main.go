package main

import (
	"context"
	"flag"
	"log"
	"net/http"

	"github.com/ArtemVoronov/indefinite-studies-utils/pkg/app"
	pb "github.com/ArtemVoronov/indefinite-studies-utils/pkg/services/greeter"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

var (
	tls      = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	certFile = flag.String("cert_file", "", "The TLS cert file")
	keyFile  = flag.String("key_file", "", "The TLS key file")
	port     = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func defaultCertFilePath() string {
	// TODO
	return "TODO"
}

func defaultKeyFilePath() string {
	// TODO
	return "TODO"
}

func main() {
	// flag.Parse()
	// lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	// if err != nil {
	// 	log.Fatalf("failed to listen: %v", err)
	// }
	// // TLS --------------------
	// var opts []grpc.ServerOption
	// if *tls {
	// 	if *certFile == "" {
	// 		*certFile = defaultCertFilePath()
	// 	}
	// 	if *keyFile == "" {
	// 		*keyFile = defaultKeyFilePath()
	// 	}
	// 	creds, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)
	// 	if err != nil {
	// 		log.Fatalf("Failed to generate credentials %v", err)
	// 	}
	// 	opts = []grpc.ServerOption{grpc.Creds(creds)}
	// }
	// // ------------------------
	// s := grpc.NewServer(opts...)
	// pb.RegisterGreeterServer(s, &server{})
	// log.Printf("server listening at %v", lis.Addr())
	// if err := s.Serve(lis); err != nil {
	// 	log.Fatalf("failed to serve: %v", err)
	// }

	serviceServer := &server{}

	registerServices := func(s *grpc.Server) {
		pb.RegisterGreeterServer(s, serviceServer)
	}
	go func() {
		app.StartGRPC(setup, shutdown, ":50051", registerServices)
	}()
	app.StartHTTP(setup, shutdown, ":3000", router())

}

func setup() {

}

func shutdown() {

}

func router() *gin.Engine {
	router := gin.Default()
	gin.SetMode(app.Mode())
	router.Use(app.Cors())
	router.Use(gin.Logger())
	v1 := router.Group("/api/v1")

	v1.GET("/ping", Ping)
	return router
}
func Ping(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Pong!")
}
