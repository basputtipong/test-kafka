package port

import "gitdev.devops.krungthai.com/hexagonal/log"

type TransferFaceVerificationRepo interface {
	Insert(req TransferFaceVerReq, l *log.Logger) error
}

type TransferFaceVerReq struct{}
