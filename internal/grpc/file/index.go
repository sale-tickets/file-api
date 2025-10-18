package file_handle

import (
	"github.com/minio/minio-go/v7"
	"github.com/sale-tickets/file-api/internal/file_repo"
	file_api "github.com/sale-tickets/golang-common/file-api/proto"
)

type FileHanle struct {
	repo        file_repo.FileRepo
	minioClient *minio.Client
	file_api.UnimplementedFileServer
}

func NewFileHanle(
	repo file_repo.FileRepo,
	minioClient *minio.Client,
) file_api.FileServer {
	return &FileHanle{
		repo:        repo,
		minioClient: minioClient,
	}
}
