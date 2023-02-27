module github.com/degalaxy/wegalaxy-service/wegalaxy-foundation

go 1.17

require (
	github.com/degalaxy/gp-common v1.1.9
	github.com/degalaxy/helper/dahelper v1.0.4-alpha.6
	github.com/degalaxy/wegalaxy-service/wegalaxy-model/app_error v0.0.0-00010101000000-000000000000
	github.com/degalaxy/wegalaxy-service/wegalaxy-model/model v0.0.0-00010101000000-000000000000
	github.com/fsnotify/fsnotify v1.5.4
	github.com/swaggo/swag v1.8.6
	go.uber.org/zap v1.21.0
	gorm.io/gorm v1.23.5
)

// replace github.com/degalaxy/gp-common => ../../gp-common

replace github.com/degalaxy/wegalaxy-service/wegalaxy-model/model => ./wegalaxy-model/model

replace github.com/degalaxy/wegalaxy-service/wegalaxy-model/app_error => ./wegalaxy-model/app_error

// replace github.com/degalaxy/helper/dahelper => ../../helper/dahelper

// replace github.com/degalaxy/helper/coshelper => ../helper/coshelper

require (
	github.com/KyleBanks/depth v1.2.1 // indirect
	github.com/aws/aws-sdk-go v1.44.174 // indirect
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-openapi/jsonpointer v0.19.5 // indirect
	github.com/go-openapi/jsonreference v0.20.0 // indirect
	github.com/go-openapi/spec v0.20.7 // indirect
	github.com/go-openapi/swag v0.22.3 // indirect
	github.com/go-playground/locales v0.14.0 // indirect
	github.com/go-playground/universal-translator v0.18.0 // indirect
	github.com/go-playground/validator/v10 v10.10.1 // indirect
	github.com/go-redis/redis/v8 v8.11.5 // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/golang-jwt/jwt v3.2.2+incompatible // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/go-cmp v0.5.8 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/magiconair/properties v1.8.6 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/mattn/go-sqlite3 v1.14.12 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pelletier/go-toml v1.9.5 // indirect
	github.com/pelletier/go-toml/v2 v2.0.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/spf13/afero v1.8.2 // indirect
	github.com/spf13/cast v1.4.1 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/subosito/gotenv v1.2.0 // indirect
	github.com/ugorji/go/codec v1.2.7 // indirect
	golang.org/x/crypto v0.0.0-20220427172511-eb4f295cb31f // indirect
	golang.org/x/sys v0.1.0 // indirect
	golang.org/x/text v0.4.0 // indirect
	golang.org/x/tools v0.1.12 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/ini.v1 v1.66.4 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	gorm.io/driver/mysql v1.3.3 // indirect
)

require (
	github.com/gin-gonic/gin v1.7.7
	github.com/petermattis/goid v0.0.0-20220331194723-8ee3e6ded87a // indirect
	github.com/spf13/viper v1.11.0
	github.com/stretchr/testify v1.8.0
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.8.0 // indirect
	gorm.io/driver/sqlite v1.3.6
)
