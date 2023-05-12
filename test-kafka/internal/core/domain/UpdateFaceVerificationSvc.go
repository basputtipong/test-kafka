package domain

import (
	"gitdev.devops.krungthai.com/hexagonal/log"
	"test-kafka/internal/core/port"
)

type UpdateFaceVerificationSvc interface {
	Execute(req UpdateFaceVerReq, l *log.Logger) error
}

type UpdateFaceVerReq struct {
	Code        string `json:"code"`
	Description string `json:"description"`
	Data        Data   `json:"data"`
}

type Data struct {
	TencentCode    string     `json:"tencentCode"`
	TencentDesc    string     `json:"tencentDesc"`
	TencentVersion string     `json:"tencentVersion"`
	ResultFlag     bool       `json:"resultFlag"`
	RiskResult     RiskResult `json:"riskResult"`
	LiveResult     LiveResult `json:"liveResult"`
}

type RiskResult struct {
	DeviceRiskCode      string `json:"deviceRiskCode"`
	DeviceRiskLevelCode int    `json:"deviceRiskLevelCode"`
	DeviceRiskTag       string `json:"deviceRiskTag"`
	StreamRiskLevelCode string `json:"streamRiskLevelCode"`
	StreamRiskTag       string `json:"streamRiskTag"`
}

type LiveResult struct {
	Code            string  `json:"code"`
	Message         string  `json:"message"`
	Similarity      int     `json:"similarity"`
	SimilarityFloat float64 `json:"similarityFloat"`
}

func (req *UpdateFaceVerReq) BuildFaceVerReq() port.FaceVerReq {
	//mapping fields
	return port.FaceVerReq{}
}

func (req *UpdateFaceVerReq) BuildTranferFaceReq() port.TransferFaceVerReq {
	//mapping fields
	return port.TransferFaceVerReq{}
}
