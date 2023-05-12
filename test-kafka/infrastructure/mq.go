package infrastructure

import (
	"strings"

	fnd "gitdev.devops.krungthai.com/hexagonal/foundation.git/v2/port/mq"
	"gitdev.devops.krungthai.com/hexagonal/kafka/v2/mq"
	"gitdev.devops.krungthai.com/hexagonal/log"
	"github.com/spf13/viper"
)

var (
	FaceReflectMqReader fnd.MqReader
	FaceLeadMqReader    fnd.MqReader
	MqWriter            fnd.MqWriter
)

func InitMQ() {
	cfg, err := getKafkaConfig()
	if err != nil {
		panic("get kafka config error")
	}

	faceReflect := strings.ToLower("faceReflect")
	faceLead := strings.ToLower("faceLead")

	FaceReflectMqReader = mq.NewMqReader(cfg, cfg.Topics[faceReflect])
	FaceLeadMqReader = mq.NewMqReader(cfg, cfg.Topics[faceLead])
	MqWriter = mq.NewWriter(cfg, cfg.Topics[faceReflect])
}

func getKafkaConfig() (*mq.KafkaCfg, error) {
	config := mq.KafkaCfg{}
	err := viper.UnmarshalKey("kafka", &config)
	if err != nil {
		log.Errorf("can not load config kafka, error: %v", err)
		return &config, err
	}
	if config.CertPath == "" {
		log.Errorf("can not load config kafka, error: %v", err)
		return nil, err
	}
	return &config, nil
}
