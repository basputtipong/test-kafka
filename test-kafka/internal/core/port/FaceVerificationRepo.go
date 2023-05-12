package port

import "gitdev.devops.krungthai.com/hexagonal/log"

type FaceVerificationRepo interface {
	Insert(req FaceVerReq, l *log.Logger) error
}

type FaceVerReq struct{}
