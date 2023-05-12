package infrastructure

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	TransferDB *gorm.DB
	FaceVerDB  *gorm.DB
)

type MysqlDbConfig struct {
	Database string `mapstructure:"database"`
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

func InitDB() {
	TransferDB = generateDsn("db.transfers")
	FaceVerDB = generateDsn("db.faceVerificationResult")
}

func generateDsn(key string) *gorm.DB {
	var dbCfg MysqlDbConfig
	err := viper.UnmarshalKey(key, &dbCfg)
	if err != nil {
		return nil
	}

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true&loc=Local",
		dbCfg.Username,
		dbCfg.Password,
		dbCfg.Host,
		dbCfg.Port,
		dbCfg.Database,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}

func pingDb(dsn *gorm.DB) error {
	db, err := dsn.DB()
	if err != nil {
		return err
	} else {
		return db.Ping()
	}
}

func PingAllDb() error {
	dbs := []*gorm.DB{
		TransferDB,
		FaceVerDB,
	}

	for _, db := range dbs {
		if err := pingDb(db); err != nil {
			return err
		}
	}
	return nil
}
