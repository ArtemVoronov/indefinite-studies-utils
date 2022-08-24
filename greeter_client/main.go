package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/ArtemVoronov/indefinite-studies-utils/pkg/services/greeter"
)

const (
	defaultName = "world"
)

var (
	tls                = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	caFile             = flag.String("ca_file", "", "The file containing the CA root cert file")
	serverHostOverride = flag.String("server_host_override", defaultServerNameOverride(), "The server name used to verify the hostname returned by the TLS handshake")
	addr               = flag.String("addr", "localhost:50051", "the address to connect to")
	name               = flag.String("name", defaultName, "Name to greet")
)

func defaultCertFilePath() string {
	// TODO
	return "TODO"
}
func defaultServerNameOverride() string {
	// TODO
	return "some.domain.name.com"
}

func main() {
	// flag.Parse()

	// // TLS --------------------
	// var opts []grpc.DialOption
	// if *tls {
	// 	if *caFile == "" {
	// 		*caFile = defaultCertFilePath()
	// 	}
	// 	creds, err := credentials.NewClientTLSFromFile(*caFile, *serverHostOverride)
	// 	if err != nil {
	// 		log.Fatalf("Failed to create TLS credentials %v", err)
	// 	}
	// 	opts = append(opts, grpc.WithTransportCredentials(creds))
	// } else {
	// 	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	// }
	// // ------------------------

	// conn, err := grpc.Dial(*addr, opts...)
	// if err != nil {
	// 	log.Fatalf("did not connect: %v", err)
	// }
	// defer conn.Close()
	// c := pb.NewGreeterClient(conn)

	// ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	// defer cancel()
	// r, err := c.SayHello(ctx, &pb.HelloRequest{Name: *name})
	// if err != nil {
	// 	log.Fatalf("could not greet: %v", err)
	// }
	// log.Printf("Greeting: %s", r.GetMessage())

	service := greeter.CreateHelloService("localhost:50051")
	defer service.Shutdown()

	names := []string{"Bob", "Alice", "Mary"}
	for i := 0; i < 100; i++ {
		err := service.Hello(names[i%3])
		if err != nil {
			fmt.Println(err)
		}
		time.Sleep(2 * time.Second)
	}
}
