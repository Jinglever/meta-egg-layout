/*
 * Generated by meta-egg.
 * WILL NOT be replace after re-generated. CAREFULLY EDIT.
 * Version: master-046024d-dirty
 * Author: meta-egg
 * Generated at: 2023-04-12 13:10
 */
package resource

import (
	"context"

	"meta-egg-layout/pkg/gormx"

	jgjwt "github.com/Jinglever/go-jwt"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Resource struct {
	DB  *gorm.DB
	JWT *jgjwt.JWT
}

type Config struct {
	DB  *gormx.Config `mapstructure:"db"`  // 数据库配置
	JWT *jgjwt.Config `mapstructure:"jwt"` // jwt配置
}

func InitResource(cfg *Config) (*Resource, context.CancelFunc, error) {
	db, err := gormx.ConnectDB(cfg.DB)
	if err != nil {
		logrus.Errorf("connect db failed, err: %v", err)
		return nil, nil, err
	}

	jwt, err := jgjwt.NewJWT(cfg.JWT)
	if err != nil {
		logrus.Errorf("init jwt failed, err: %v", err)
		return nil, nil, err
	}

	// your resource init here

	rsrc := &Resource{
		DB:  db,
		JWT: jwt,
	}
	cancel := func() {
		logrus.Info("release resource")
		sqlDB, err := rsrc.DB.DB()
		if err == nil {
			err = sqlDB.Close()
			if err != nil {
				logrus.Errorf("close db error: %v", err)
			}
		}
	}
	return rsrc, cancel, nil
}
