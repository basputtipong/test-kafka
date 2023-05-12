package repo

import (
	"gitdev.devops.krungthai.com/hexagonal/gormlog"
	"gitdev.devops.krungthai.com/hexagonal/log"
	"gorm.io/gorm"
	"test-kafka/internal/core/port"
)

type faceVerificationRepo struct {
	dbFaceVer *gorm.DB
}

func NewFaceVerificationRepo(dbFaceVer *gorm.DB) port.FaceVerificationRepo {
	return &faceVerificationRepo{dbFaceVer}
}

func (r *faceVerificationRepo) Insert(req port.FaceVerReq, l *log.Logger) error {
	ll := gormlog.NewGormLogger(l)

	tx := r.dbFaceVer.Session(&gorm.Session{Logger: ll})
	err := tx.Save(req).Error

	return err
}
