module test-kafka

go 1.19

require (
	gitdev.devops.krungthai.com/hexagonal/foundation.git/v2 v2.1.1
	gitdev.devops.krungthai.com/hexagonal/gin.git/v2 v2.1.0
	gitdev.devops.krungthai.com/hexagonal/kafka/v2 v2.1.0
	gitdev.devops.krungthai.com/hexagonal/log v1.5.1
	gitdev.devops.krungthai.com/hexagonal/middleware.git/v2 v2.0.0
	gitdev.devops.krungthai.com/ktb-next/backend/channel-api/nextenv-go.git v0.1.2
	gitdev.devops.krungthai.com/ktb-next/backend/channel-api/nexterr-go.git/v2 v2.3.0
	gitdev.devops.krungthai.com/ktb-next/backend/channel-api/nextfnd-go.git/v2 v2.2.0
	gitdev.devops.krungthai.com/ktb-next/backend/channel-api/nextjwt-go.git/v2 v2.0.1
	gitdev.devops.krungthai.com/ktb-next/backend/channel-api/nextnet-go.git/v2 v2.0.0
	github.com/spf13/viper v1.15.0
	gorm.io/driver/mysql v1.5.0
	gorm.io/gorm v1.25.1
)

replace gitdev.devops.krungthai.com/hexagonal/kafka/v2 => /Users/p.thammachart/Desktop/finrevamp-lib/kafka

require (
	cloud.google.com/go v0.109.0 // indirect
	cloud.google.com/go/compute v1.15.1 // indirect
	cloud.google.com/go/compute/metadata v0.2.3 // indirect
	cloud.google.com/go/iam v0.10.0 // indirect
	cloud.google.com/go/storage v1.28.1 // indirect
	gitdev.devops.krungthai.com/hexagonal/gcs.git v0.1.2 // indirect
	gitdev.devops.krungthai.com/hexagonal/gormlog v0.2.7 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/gin-gonic/gin v1.8.1 // indirect
	github.com/go-playground/locales v0.14.0 // indirect
	github.com/go-playground/universal-translator v0.18.0 // indirect
	github.com/go-playground/validator/v10 v10.11.0 // indirect
	github.com/go-sql-driver/mysql v1.7.0 // indirect
	github.com/go-task/slim-sprig v0.0.0-20210107165309-348f09dbbbc0 // indirect
	github.com/goccy/go-json v0.9.11 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/mock v1.6.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/google/pprof v0.0.0-20230510003242-05f3f57da91a // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.2.1 // indirect
	github.com/googleapis/gax-go/v2 v2.7.0 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/imroc/req/v3 v3.32.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/compress v1.16.0 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/magiconair/properties v1.8.7 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.17 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/onsi/ginkgo/v2 v2.7.0 // indirect
	github.com/pelletier/go-toml/v2 v2.0.6 // indirect
	github.com/pierrec/lz4/v4 v4.1.17 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/quic-go/qpack v0.4.0 // indirect
	github.com/quic-go/qtls-go1-18 v0.2.0 // indirect
	github.com/quic-go/qtls-go1-19 v0.2.0 // indirect
	github.com/quic-go/qtls-go1-20 v0.1.0 // indirect
	github.com/quic-go/quic-go v0.32.0 // indirect
	github.com/rs/zerolog v1.29.0 // indirect
	github.com/segmentio/kafka-go v0.4.39 // indirect
	github.com/spf13/afero v1.9.4 // indirect
	github.com/spf13/cast v1.5.0 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/stretchr/objx v0.5.0 // indirect
	github.com/stretchr/testify v1.8.2 // indirect
	github.com/subosito/gotenv v1.4.2 // indirect
	github.com/ugorji/go/codec v1.2.7 // indirect
	github.com/xdg/scram v1.0.5 // indirect
	github.com/xdg/stringprep v1.0.3 // indirect
	go.opencensus.io v0.24.0 // indirect
	golang.org/x/crypto v0.6.0 // indirect
	golang.org/x/exp v0.0.0-20230118134722-a68e582fa157 // indirect
	golang.org/x/mod v0.7.0 // indirect
	golang.org/x/net v0.6.0 // indirect
	golang.org/x/oauth2 v0.4.0 // indirect
	golang.org/x/sys v0.5.0 // indirect
	golang.org/x/text v0.7.0 // indirect
	golang.org/x/tools v0.5.0 // indirect
	golang.org/x/xerrors v0.0.0-20220907171357-04be3eba64a2 // indirect
	google.golang.org/api v0.108.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto v0.0.0-20230117162540-28d6b9783ac4 // indirect
	google.golang.org/grpc v1.52.0 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
