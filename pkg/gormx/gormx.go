// Code generated by meta-egg. CAREFULLY EDIT.
// WILL NOT BE replace after re-generated, unless you confirm it!
// CAREFULLY EDIT.
// Version: v1.0.15-EE
// Author: meta-egg
// Generated at: 2023-04-25 22:58
package gormx

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Config struct {
	Type string `mapstructure:"type"` // 数据库类型
	DSN  string `mapstructure:"dsn"`  // 数据库连接字符串
}

// ConnectDB 连接数据库
func ConnectDB(cfg *Config) (*gorm.DB, error) {
	var (
		db  *gorm.DB
		err error
	)
	switch cfg.Type {
	case "mysql":
		db, err = connectMySQL(cfg.DSN)
	case "postgresql":
		db, err = connectPostgreSQL(cfg.DSN)
	default:
		log.Fatalf("unsupported database type: %s", cfg.Type)
		return nil, fmt.Errorf("unsupported database type: %s", cfg.Type)
	}
	return db, err
}

func connectMySQL(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:                 logger.Default.LogMode(logger.Info),
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func connectPostgreSQL(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger:                 logger.Default.LogMode(logger.Info),
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})
	if err != nil {
		return nil, err
	}
	return db, nil
}
