package view

import (
	"errors"

	file_api "github.com/sale-tickets/golang-common/file-api/proto"
)

type CreateTicketFileReq struct {
	*file_api.CreateTicketFileReq
}

func (v *CreateTicketFileReq) Validate() error {
	if v.Data == nil {
		return errors.New("data not found")
	}
	return nil
}
