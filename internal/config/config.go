// Code generated by meta-egg. CAREFULLY EDIT.
// WILL NOT BE replace after re-generated, unless you confirm it!
// CAREFULLY EDIT.
// Version: v1.0.0-EE
// Author: meta-egg
// Generated at: 2023-04-16 18:27
package config

import (
	"meta-egg-layout/internal/common/resource"
	grpcserver "meta-egg-layout/internal/server/grpc"
	httpserver "meta-egg-layout/internal/server/http"

	jgconf "github.com/Jinglever/go-config"
	"github.com/Jinglever/go-config/option"
)

type Config struct {
	Resource   *resource.Config   `mapstructure:"resource"`    // 资源配置
	HTTPServer *httpserver.Config `mapstructure:"http_server"` // http服务配置
	GRPCServer *grpcserver.Config `mapstructure:"grpc_server"` // grpc服务配置
}

var cfg *Config

const EnvPrefix = "MEL"

func LoadConfig(path string) error {
	var c Config
	if err := jgconf.LoadYamlConfig(path, &c, option.WithEnvPrefix(EnvPrefix)); err != nil {
		return err
	}
	cfg = &c
	return nil
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
