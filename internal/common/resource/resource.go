// Code generated by meta-egg. CAREFULLY EDIT.
// WILL NOT BE replace after re-generated, unless you confirm it!
// CAREFULLY EDIT.
// Version: v1.5.1-EE
// Author: meta-egg
// Generated at: 2023-05-22 16:28

package resource

import (
	"context"

	"meta-egg-layout/pkg/gormx"
	log "meta-egg-layout/pkg/log"

	jgjwt "github.com/Jinglever/go-jwt"
)

type Resource struct {
	cancelFuncs []context.CancelFunc
	DB          gormx.DB
	JWT         *jgjwt.JWT
}

type Config struct {
	DB  *gormx.Config `mapstructure:"db"`  // 数据库配置
	JWT *jgjwt.Config `mapstructure:"jwt"` // jwt配置
}

func InitResource(cfg *Config) (*Resource, error) {
	rsrc := &Resource{
		cancelFuncs: make([]context.CancelFunc, 0),
	}

	if cfg.DB != nil {
		db, err := gormx.ConnectDB(cfg.DB)
		if err != nil {
			log.WithError(err).Error("connect db failed, err")
			return nil, err
		}
		rsrc.DB = db
		rsrc.cancelFuncs = append(rsrc.cancelFuncs, func() {
			if err = rsrc.DB.Close(); err != nil {
				log.WithError(err).Errorf("close db error")
			}
			log.Info("close db success")
		})
	}

	if cfg.JWT != nil {
		jwt, err := jgjwt.NewJWT(cfg.JWT)
		if err != nil {
			log.WithError(err).Error("init jwt failed")
			return nil, err
		}
		rsrc.JWT = jwt
	}

	// your resource init here

	return rsrc, nil
}

func (r *Resource) Release() {
	for _, cancel := range r.cancelFuncs {
		cancel()
	}
}
