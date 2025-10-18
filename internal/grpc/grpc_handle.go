package grpc_handle

import (
	"fmt"
	"log"
	"net"

	"github.com/godev-lib/golang/config"
	file_api "github.com/sale-tickets/golang-common/file-api/proto"
	"google.golang.org/grpc"
)

func GrpcServer(
	config *config.Config,
	fileHanle file_api.FileServer,
) {
	port := fmt.Sprintf(":%s", config.App.GrpcPort)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer(
		grpc.ChainUnaryInterceptor(),
	)
	file_api.RegisterFileServer(s, fileHanle)

	log.Printf("gRPC server running on %s", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalln("error start grpc server: ", err.Error())
	}
}
