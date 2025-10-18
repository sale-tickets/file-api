package file_handle

import (
	"context"
	"time"

	"github.com/sale-tickets/file-api/internal/view"
	file_api "github.com/sale-tickets/golang-common/file-api/proto"
)

func (h *FileHanle) Create(ctx context.Context, req *file_api.CreateTicketFileReq) (*file_api.CreateTicketFileRes, error) {
	reqData := view.CreateTicketFileReq{
		CreateTicketFileReq: req,
	}

	if err := reqData.Validate(); err != nil {
		return nil, err
	}

	paths := []*file_api.FileModel{}
	for _, item := range reqData.Data {
		policy, err := h.minioClient.PresignedPutObject(
			ctx,
			"sale-tickets",
			item.Path,
			5*time.Minute,
		)
		if err != nil {
			return nil, err
		}

		paths = append(paths, &file_api.FileModel{
			Path:     policy.Path,
			HrefEdit: policy.RequestURI(),
		})
	}

	result := &file_api.CreateTicketFileRes{
		Data: paths,
	}

	return result, nil
}
