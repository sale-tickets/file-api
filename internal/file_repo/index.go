package file_repo

import (
	"github.com/godev-lib/golang/orm"
	"github.com/sale-tickets/file-api/internal/model"
	"gorm.io/gorm"
)

type FileRepo interface {
	orm.DataMethod[model.File]
}

func NewFileRepo(db *gorm.DB) FileRepo {
	return orm.NewOrm[model.File](db)
}
