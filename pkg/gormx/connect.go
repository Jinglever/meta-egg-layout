// Code generated by meta-egg. CAREFULLY EDIT.
// WILL NOT BE replace after re-generated, unless you confirm it!
// CAREFULLY EDIT.
// Version: master-b9138a4-IE
// Author: meta-egg
// Generated at: 2023-08-20 19:41

package gormx

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Config struct {
	Type     string `mapstructure:"type"`      // 数据库类型
	DSN      string `mapstructure:"dsn"`       // 数据库连接字符串
	LogLevel string `mapstructure:"log_level"` // 日志级别 info|warn|error|silent
}

// ConnectDB 连接数据库
func ConnectDB(cfg *Config) (DB, error) {
	var (
		db  *gorm.DB
		err error
	)
	logLevel := convertLogLevel(cfg.LogLevel)
	gormOpt := &gorm.Config{
		Logger:                 logger.Default.LogMode(logLevel),
		SkipDefaultTransaction: true,  // 根据官方文档，可以提高30%性能
		PrepareStmt:            false, // 关闭预编译，因为gorm的预编译会导致不支持savepoint
		TranslateError:         true,  // 自动翻译错误
	}
	switch cfg.Type {
	case "mysql":
		db, err = connectMySQL(cfg.DSN, gormOpt)
	case "postgres":
		db, err = connectPostgreSQL(cfg.DSN, gormOpt)
	default:
		return nil, fmt.Errorf("unsupported database type: %s", cfg.Type)
	}
	return &DBImpl{DB: db}, err
}

// convert log level string to gorm logger level
func convertLogLevel(level string) logger.LogLevel {
	switch level {
	case "info":
		return logger.Info
	case "warn":
		return logger.Warn
	case "error":
		return logger.Error
	case "silent":
		return logger.Silent
	default:
		return logger.Info
	}
}

func connectMySQL(dsn string, opts ...gorm.Option) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), opts...)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func connectPostgreSQL(dsn string, opts ...gorm.Option) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), opts...)
	if err != nil {
		return nil, err
	}
	return db, nil
}
