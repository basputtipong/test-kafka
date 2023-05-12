package repo

import (
	"gitdev.devops.krungthai.com/hexagonal/gormlog"
	"gitdev.devops.krungthai.com/hexagonal/log"
	"gorm.io/gorm"
	"test-kafka/internal/core/port"
)

type transferFaceVerificationRepo struct {
	dbTransfer *gorm.DB
}

func NewTransferFaceVerificationRepo(dbTransfer *gorm.DB) port.TransferFaceVerificationRepo {
	return &transferFaceVerificationRepo{dbTransfer}
}

func (r *transferFaceVerificationRepo) Insert(req port.TransferFaceVerReq, l *log.Logger) error {
	ll := gormlog.NewGormLogger(l)

	tx := r.dbTransfer.Session(&gorm.Session{Logger: ll})
	err := tx.Save(req).Error

	return err
}
