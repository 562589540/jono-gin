package gbootstrap

import (
	"fmt"
	"github.com/spf13/viper"
	"time"
)

// SystemConfig 系统配置
type SystemConfig struct {
	NotCheckAuthAdminIds []uint `mapstructure:"notCheckAuthAdminIds"`
	NodeNumber           int64  `mapstructure:"nodeNumber"`
}

// CasbinConfig casbin权限配置
type CasbinConfig struct {
	ModelFile  string `mapstructure:"modelFile"`
	PolicyFile string `mapstructure:"policyFile"`
}

// PathConfig 本地路径配置
type PathConfig struct {
	Static       string `mapstructure:"static"`
	ResourcePath string `mapstructure:"resourcePath"`
	UploadsPath  string `mapstructure:"uploadsPath"`
	AvatarPath   string `mapstructure:"avatarPath"`
}

// ServerConfig 服务配置
type ServerConfig struct {
	Port int `mapstructure:"port"`
}

// LogConfig 日志配置
type LogConfig struct {
	Output     string `mapstructure:"output"`
	MaxSize    int    `mapstructure:"maxSize"`
	MaxBackups int    `mapstructure:"maxBackups"`
	MaxAge     int    `mapstructure:"maxAge"`
}

// ModeConfig 环境配置
type ModeConfig struct {
	Develop bool `mapstructure:"develop"`
}

// DBConfig mysql配置
type DBConfig struct {
	DSN         string `mapstructure:"dsn"`
	MaxIdleCons int    `mapstructure:"maxIdleCons"`
	MaxOpenCons int    `mapstructure:"maxOpenCons"`
}

// RedisConfig redis 配置
type RedisConfig struct {
	URL      string `mapstructure:"url"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

// JwtConfig jwt配置
type JwtConfig struct {
	TokenExpire        time.Duration `mapstructure:"tokenExpire"`
	RefreshTokenExpire time.Duration `mapstructure:"refreshTokenExpire"`
	SigningKey         string        `mapstructure:"signingKey"`
}

type Configuration struct {
	Server ServerConfig `mapstructure:"server"`
	Log    LogConfig    `mapstructure:"log"`
	Mode   ModeConfig   `mapstructure:"mode"`
	DB     DBConfig     `mapstructure:"db"`
	Redis  RedisConfig  `mapstructure:"redis"`
	Jwt    JwtConfig    `mapstructure:"jwt"`
	Path   PathConfig   `mapstructure:"path"`
	Casbin CasbinConfig `mapstructure:"casbin"`
	System SystemConfig `mapstructure:"system"`
}

var cfg *Configuration

func InitConfig() (*Configuration, error) {
	viper.SetConfigName("setting")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config/")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
		return nil, err
	}
	// 设置日志相关配置项的默认值
	viper.SetDefault("server.port", 8080)               // 默认服务端口号
	viper.SetDefault("log.maxSize", 10)                 // 默认日志文件最大尺寸(M)
	viper.SetDefault("log.maxBackups", 3)               // 默认保留旧文件的最大个数
	viper.SetDefault("log.maxAge", 7)                   // 默认保留旧文件的最大天数
	viper.SetDefault("log.output", "both")              // 默认日志输出方式
	viper.SetDefault("db.maxIdleCons", 10)              // 默认最多空闲连接数
	viper.SetDefault("db.maxOpenCons", 10)              // 默认最多打开链接数
	viper.SetDefault("path.resourcePath", "./resource") // 项目静态目录
	viper.SetDefault("path.uploadsPath", "uploads")     // 本地上传目录
	viper.SetDefault("path.avatarPath", "avatar")       // 头像保存目录
	viper.SetDefault("path.static", "static")           // 静态跟
	temp := Configuration{}
	err := viper.Unmarshal(&temp)
	if err != nil {
		return nil, err
	}
	cfg = &temp
	return cfg, nil
}
