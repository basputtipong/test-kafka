package handler

import (
	"encoding/json"
	"gitdev.devops.krungthai.com/hexagonal/foundation.git/v2/port/mq"
	"gitdev.devops.krungthai.com/hexagonal/log"
	"test-kafka/internal/core/domain"
)

type updateFaceVerificationHdl struct {
	svc domain.UpdateFaceVerificationSvc
}

func NewUpdateFaceVerificationHdl(service domain.UpdateFaceVerificationSvc) *updateFaceVerificationHdl {
	return &updateFaceVerificationHdl{svc: service}
}

func (hdl *updateFaceVerificationHdl) Handle(message mq.Message, l *log.Logger) error {
	req := domain.UpdateFaceVerReq{}
	if err := json.Unmarshal(message.Value(), &req); err != nil {
		log.Errorf("error unmarshal,  %v", err)
		return err
	}
	err := hdl.svc.Execute(req, l)
	return err
}
