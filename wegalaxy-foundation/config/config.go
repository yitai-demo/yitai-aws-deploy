package config

import (
	"crypto/rsa"
	"log"

	gp_config "github.com/degalaxy/gp-common/config"
	gp_log "github.com/degalaxy/gp-common/log"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type ServerConfig struct {
	Port           string          `json:"port" yaml:"port"`
	PrivateKeyPath string          `json:"privateKeyPath" yaml:"privateKeyPath" mapstructure:"private-key-path"`
	PrivateKey     *rsa.PrivateKey `json:"-" yaml:"-"`
}

type DatabaseConfig struct {
	Login             string `json:"login" yaml:"login"`
	Password          string `json:"password" yaml:"password"`
	Host              string `json:"host" yaml:"host"`
	Port              string `json:"port" yaml:"port"`
	DatabaseName      string `json:"databaseName" yaml:"databaseName" mapstructure:"database-name"`
	ConnectionMaxIdle int    `json:"connectionMaxIdle" yaml:"connectionMaxIdle" mapstructure:"connection-max-idle"`
	ConnectionMaxOpen int    `json:"connectionMaxOpen" yaml:"connectionMaxOpen" mapstructure:"connection-max-open"`
}

type CookieConfig struct {
	Domain   string `json:"domain" yaml:"domain"`
	Secure   bool   `json:"secure" yaml:"secure"`
	HttpOnly bool   `json:"httpOnly" yaml:"httpOnly" mapstructure:"http-only"`
}

type ApplicationConfig struct {
	Server  ServerConfig            `json:"server" yaml:"server"`
	DB      DatabaseConfig          `json:"database" yaml:"database" mapstructure:"database"`
	Log     gp_log.LogConfig        `json:"log" yaml:"log"`
	AppAuth gp_config.AppAuthConfig `json:"appAuth" yaml:"appAuth" mapstructure:"app-auth"`
	Cookie  CookieConfig            `json:"cookie" yaml:"cookie"`
	Redis   RedisConfig             `json:"redis" yaml:"redis" mapstructure:"redis"`
	App     AppConfig               `json:"appConfig" yaml:"appConfig" mapstructure:"app-config"`
}

type AppConfig struct {
	GamicoreServiceUrl  string `json:"gamicoreServiceUrl" yaml:"gamicoreServiceUrl" mapstructure:"gamicore-service-url"`
	LocalhostServiceUrl string `json:"localhostServiceUrl" yaml:"localhostServiceUrl" mapstructure:"localhost-service-url"`
	GamiRankCoreUrl     string `json:"gamiRankCoreUrl" yaml:"gamiRankCoreUrl" mapstructure:"gami-rank-core-url"`

	AppId                 string `json:"appId" yaml:"appId" mapstructure:"appId"`
	UserServiceUrl        string `json:"userServiceUrl" yaml:"userServiceUrl" mapstructure:"user-service-url"`
	UserServiceSourceAuth bool   `json:"userServiceSourceAuth" yaml:"userServiceSourceAuth" mapstructure:"user-service-source-auth"`
	TradeServiceUrl       string `json:"tradeServiceUrl" yaml:"tradeServiceUrl" mapstructure:"trade-service-url"`

	SecretId  string `json:"secretId" yaml:"secretId" mapstructure:"secret-id"`
	SecretKey string `json:"secretKey" yaml:"secretKey" mapstructure:"secret-key"`
	CosName   string `json:"cosName" yaml:"cosName" mapstructure:"cos-name"`
}

type RedisConfig struct {
	Enabled  bool   `json:"enabled" yaml:"enabled" mapstructure:"enabled"`
	Address  string `json:"address" yaml:"address" mapstructure:"address"`
	User     string `json:"user" yaml:"user" mapstructure:"user"`
	Password string `json:"password" yaml:"password" mapstructure:"password"`
	UseTLS   bool   `json:"useTLS" yaml:"useTLS" mapstructure:"use-tls"`
}

var (
	GlobalConfig *ApplicationConfig
)

func InitConfigAndWatchChanges(configFilePath string, configChangeHandler func(fsnotify.Event)) {
	viper.SetConfigFile(configFilePath)
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	GlobalConfig = &ApplicationConfig{}
	err = viper.Unmarshal(GlobalConfig)
	if err != nil {
		log.Fatalf("Unable to decode into config struct: %s", err)
	}

	if configChangeHandler != nil {
		viper.WatchConfig()
		viper.OnConfigChange(configChangeHandler)
	}
}
