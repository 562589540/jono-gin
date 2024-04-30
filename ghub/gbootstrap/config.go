package gbootstrap

import (
	"fmt"
	"github.com/spf13/viper"
	"time"
)

type SystemConfig struct {
	NotCheckAuthAdminIds []uint `mapstructure:"notCheckAuthAdminIds"`
}

type CasbinConfig struct {
	ModelFile  string `mapstructure:"modelFile"`
	PolicyFile string `mapstructure:"policyFile"`
}

type PathConfig struct {
	Static       string `mapstructure:"static"`
	ResourcePath string `mapstructure:"resourcePath"`
	UploadsPath  string `mapstructure:"uploadsPath"`
	AvatarPath   string `mapstructure:"avatarPath"`
}

type ServerConfig struct {
	Port int `mapstructure:"port"`
}

type LogConfig struct {
	Output     string `mapstructure:"output"`
	MaxSize    int    `mapstructure:"maxSize"`
	MaxBackups int    `mapstructure:"maxBackups"`
	MaxAge     int    `mapstructure:"maxAge"`
}

type ModeConfig struct {
	Develop bool `mapstructure:"develop"`
}

type DBConfig struct {
	DSN         string `mapstructure:"dsn"`
	MaxIdleCons int    `mapstructure:"maxIdleCons"`
	MaxOpenCons int    `mapstructure:"maxOpenCons"`
}

type RedisConfig struct {
	URL      string `mapstructure:"url"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

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

var Cfg *Configuration

func InitConfig() {
	viper.SetConfigName("setting")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config/")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
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

	var cfg Configuration
	err := viper.Unmarshal(&cfg)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v\n", err)
		return
	}
	Cfg = &cfg
}
