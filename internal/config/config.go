// Code generated by meta-egg. CAREFULLY EDIT.
// WILL NOT BE replace after re-generated, unless you confirm it!
// CAREFULLY EDIT.
// Version: v1.5.12-EE
// Author: meta-egg
// Generated at: 2023-06-09 16:24

package config

import (
	"meta-egg-layout/internal/common/constraint"
	"meta-egg-layout/internal/common/resource"
	grpcserver "meta-egg-layout/internal/server/grpc"
	httpserver "meta-egg-layout/internal/server/http"
	"meta-egg-layout/internal/server/monitor"
	log "meta-egg-layout/pkg/log"
	"time"

	jgconf "github.com/Jinglever/go-config"
	"github.com/Jinglever/go-config/option"
)

type Config struct {
	Log           *LogConfig         `mapstructure:"log"`            // 系统日志配置
	Constraint    *constraint.Config `mapstructure:"constraint"`     // 业务约束配置
	Resource      *resource.Config   `mapstructure:"resource"`       // 资源配置
	HTTPServer    *httpserver.Config `mapstructure:"http_server"`    // http服务配置
	GRPCServer    *grpcserver.Config `mapstructure:"grpc_server"`    // grpc服务配置
	MonitorServer *monitor.Config    `mapstructure:"monitor_server"` // 监控服务配置
}

type LogConfig struct {
	Level string `mapstructure:"level"` // 日志级别 debug, info, warn, error, fatal
}

var cfg *Config

const EnvPrefix = "MEL"

func LoadConfig(path string) error {
	var c Config
	if err := jgconf.LoadYamlConfig(path, &c, option.WithEnvPrefix(EnvPrefix)); err != nil {
		return err
	}
	cfg = &c

	// 业务约束配置
	if cfg.Constraint != nil {
		constraint.InjectConfig(*cfg.Constraint)

		// 设定时区
		loc, err := time.LoadLocation(constraint.GetConfig().Timezone)
		if err != nil {
			log.WithError(err).Error("load location failed")
			return err
		}
		time.Local = loc
	}
	return nil
}

func GetLogConfig() *LogConfig {
	return cfg.Log
}

func GetResourceConfig() *resource.Config {
	return cfg.Resource
}

func GetHTTPServerConfig() *httpserver.Config {
	return cfg.HTTPServer
}

func GetGRPCServerConfig() *grpcserver.Config {
	return cfg.GRPCServer
}

func GetMonitorServerConfig() *monitor.Config {
	return cfg.MonitorServer
}
