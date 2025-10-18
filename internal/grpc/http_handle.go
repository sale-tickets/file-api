package grpc_handle

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/godev-lib/golang/config"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	file_api "github.com/sale-tickets/golang-common/file-api/proto"
	"google.golang.org/grpc"
)

func HttpServer(
	config *config.Config,
	fileHandle file_api.FileServer,
) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	grpcAddr := fmt.Sprintf(
		"%s:%s",
		config.App.Host,
		config.App.GrpcPort,
	)
	file_api.RegisterFileHandlerFromEndpoint(ctx, mux, grpcAddr, opts)

	httpAddr := fmt.Sprintf(
		":%s",
		config.App.HttpPort,
	)
	log.Printf("REST gateway running on %s", httpAddr)
	server := http.Server{
		Addr:    httpAddr,
		Handler: allowCORS(mux),
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatalln("error start http server: ", err.Error())
	}
}

func allowCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Nếu là request "OPTIONS" (preflight), trả về luôn
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		h.ServeHTTP(w, r)
	})
}
