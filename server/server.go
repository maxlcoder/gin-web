package server

import (
	"crypto/tls"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/maxlcoder/gin-web/pkg/util"
	pb "github.com/maxlcoder/gin-web/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
	"net/http"
)

var (
	ServerPort string
	CertName string
	CertPemPath string
	CertKeyPath string
	EndPoint string

	tlsConfig *tls.Config
)

func Serve() (err error){
	EndPoint = ":" + ServerPort
	tlsConfig = util.GetTLSConfig(CertPemPath, CertKeyPath)

	conn, err := net.Listen("tcp", EndPoint)
	if err != nil {
		log.Printf("TCP Listen err:%v\n", err)
	}

	srv := createInternalServer(conn)

	log.Printf("gRPC and https listen on: %s\n", ServerPort)
	if err = srv.Serve(conn); err != nil {
		log.Printf("ListenAndServe: %v\n", err)
	}

	log.Println("debug NewListener")

	return err
}

func createInternalServer(conn net.Listener) (*http.Server) {
	var opts []grpc.ServerOption

	// grpc server
	creds, err := credentials.NewServerTLSFromFile(CertPemPath, CertKeyPath)
	if err != nil {
		log.Printf("Failed to create server TLS credentials %v", err)
	}

	opts = append(opts, grpc.Creds(creds))
	grpcServer := grpc.NewServer(opts...)

	// register grpc pb
	pb.RegisterHelloWorldServer(grpcServer, NewHelloService())

	// gw server
	ctx := context.Background()
	dcreds, err := credentials.NewClientTLSFromFile(CertPemPath, CertName)
	if err != nil {
		log.Printf("Failed to create client TLS credentials %v", err)
	}
	dopts := []grpc.DialOption{grpc.WithTransportCredentials(dcreds)}
	gwmux := runtime.NewServeMux()

	// register grpc-gateway pb
	if err := pb.RegisterHelloWorldHandlerFromEndpoint(ctx, gwmux, EndPoint, dopts); err != nil {
		log.Printf("Failed to register gw server: %v\n", err)
	}

	// http服务
	mux := http.NewServeMux()
	mux.Handle("/", gwmux)

	return &http.Server{
		Addr:      EndPoint,
		Handler:   util.GrpcHandlerFunc(grpcServer, mux),
		TLSConfig: tlsConfig,
	}
}