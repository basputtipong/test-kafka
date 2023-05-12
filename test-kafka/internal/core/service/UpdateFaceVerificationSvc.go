package service

import (
	"gitdev.devops.krungthai.com/hexagonal/log"
	"gitdev.devops.krungthai.com/ktb-next/backend/channel-api/nexterr-go.git/v2"
	"test-kafka/internal/core/domain"
	"test-kafka/internal/core/port"
)

type updateFaceVerificationSvc struct {
	faceVerRepo  port.FaceVerificationRepo
	transferRepo port.TransferFaceVerificationRepo
}

func NewUpdateFaceVerificationSvc(faceVerRepo port.FaceVerificationRepo, transferRepo port.TransferFaceVerificationRepo) domain.UpdateFaceVerificationSvc {
	return &updateFaceVerificationSvc{
		faceVerRepo:  faceVerRepo,
		transferRepo: transferRepo,
	}
}

func (s *updateFaceVerificationSvc) Execute(req domain.UpdateFaceVerReq, l *log.Logger) error {
	faceVerReq := req.BuildFaceVerReq()
	err := s.faceVerRepo.Insert(faceVerReq, l)

	transferFaceVerReq := req.BuildTranferFaceReq()
	err = s.transferRepo.Insert(transferFaceVerReq, l)

	errRes := err.(nexterr.NextError)
	return errRes
}
