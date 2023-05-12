package infrastructure

import (
	"gitdev.devops.krungthai.com/hexagonal/log"
	"github.com/spf13/viper"
)

func InitError() {
	cfg := errorBucketConfig{}
	if err := viper.UnmarshalKey("gcp.storage.errorMapping", &cfg); err != nil {
		panic(err)
	}
	log.Debugf("loaded gcs error-mapping configuration = %v", cfg)
	//nexterr.InitErrConfig(cfg.Token, cfg.Bucket, cfg.File)
}

type errorBucketConfig struct {
	Token  string `mapstructure:"token"`
	Bucket string `mapstructure:"bucket"`
	File   string `mapstructure:"file"`
}
