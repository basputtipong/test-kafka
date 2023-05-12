package listener

import (
	"context"
	"encoding/json"
	"gitdev.devops.krungthai.com/hexagonal/foundation.git/v2/port/mq"
	"gitdev.devops.krungthai.com/ktb-next/backend/channel-api/nexterr-go.git/v2"
	"gitdev.devops.krungthai.com/ktb-next/backend/channel-api/nextfnd-go.git/v2/exchangelog/mqlog"
	"gitdev.devops.krungthai.com/ktb-next/backend/channel-api/nextnet-go.git/v2"
	"test-kafka/infrastructure"
	"test-kafka/internal/adaptor/handler"
	"test-kafka/internal/adaptor/repo"
	"test-kafka/internal/core/service"
)

var client = nextnet.NewClient(nextnet.ClientConfig{})

func Run() {
	bindSomeThing(infrastructure.FaceReflectMqReader, infrastructure.FaceLeadMqReader)
}

func bindSomeThing(faceReflectReader, faceLeadReader mq.MqReader) {
	faceVerRepo := repo.NewFaceVerificationRepo(infrastructure.FaceVerDB)
	transferRepo := repo.NewTransferFaceVerificationRepo(infrastructure.TransferDB)
	svc := service.NewUpdateFaceVerificationSvc(faceVerRepo, transferRepo)
	hdl := handler.NewUpdateFaceVerificationHdl(svc)

	PublishMessage()

	ctx := context.Background()
	faceReflectReader.ReadMessage(ctx, mqlog.TraceSubscriber(hdl.Handle))
	//faceLeadReader.ReadMessage(ctx, mqlog.TraceSubscriber(hdl.Handle))
}

func PublishMessage() nexterr.NextError {
	header := mq.Header{Key: "traceId", Value: "testTraceId"}
	req := UpdateFaceVerReq{
		Code:        "0000",
		Description: "Success",
		Data: Data{
			TencentCode:    "0",
			TencentDesc:    "ok",
			TencentVersion: "1.0.0",
			ResultFlag:     true,
			RiskResult: RiskResult{
				DeviceRiskCode:      "0",
				DeviceRiskLevelCode: 3,
				DeviceRiskTag:       "",
				StreamRiskLevelCode: "",
				StreamRiskTag:       "",
			},
			LiveResult: LiveResult{
				Code:            "0",
				Message:         "ok",
				Similarity:      90,
				SimilarityFloat: 90.0,
			},
		},
	}

	b, err := json.Marshal(req)

	if err != nil {
		return nexterr.NewTechnicalErr("80000", err)
	}

	if err = infrastructure.MqWriter.WriteMessage(b, header); err != nil {
		return nexterr.NewTechnicalErr("80000", err)
	}

	return nil
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
