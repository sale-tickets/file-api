package main

import (
	"context"

	"github.com/godev-lib/golang/config"
	minio_custom "github.com/godev-lib/golang/minio"
	"github.com/godev-lib/golang/psql"
	"github.com/minio/minio-go/v7"
	"github.com/sale-tickets/file-api/internal/file_repo"
	grpc_handle "github.com/sale-tickets/file-api/internal/grpc"
	file_handle "github.com/sale-tickets/file-api/internal/grpc/file"
	file_api "github.com/sale-tickets/golang-common/file-api/proto"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

func main() {
	components := []fx.Option{}

	components = append(components, fx.Module("config",
		fx.Provide(func() *config.Config {
			return config.NewConfig()
		}),
	))

	components = append(components, fx.Module(
		"connection",
		fx.Provide(func(config *config.Config) *gorm.DB {
			return psql.NewConnectionPsql(config)
		}),
		fx.Provide(func(config *config.Config) *minio.Client {
			return minio_custom.NewMinioClient(config)
		}),
	))

	components = append(components, fx.Module(
		"repo",
		fx.Provide(func(db *gorm.DB) file_repo.FileRepo {
			return file_repo.NewFileRepo(db)
		}),
	))

	components = append(components, fx.Module(
		"handle",
		fx.Provide(func(
			fileRepo file_repo.FileRepo,
			minioClient *minio.Client,
		) file_api.FileServer {
			return file_handle.NewFileHanle(fileRepo, minioClient)
		}),
	))

	components = append(components, fx.Module(
		"run",
		fx.Invoke(func(
			lc fx.Lifecycle,
			config *config.Config,
			fileHandle file_api.FileServer,
		) {
			lc.Append(fx.Hook{
				OnStart: func(ctx context.Context) error {
					go grpc_handle.GrpcServer(config, fileHandle)
					go grpc_handle.HttpServer(config, fileHandle)
					return nil
				},
				OnStop: func(ctx context.Context) error {
					return nil
				},
			})
		}),
	))

	fx.New(components...).Run()
}
