package unit

import (
	"meta-egg-layout/internal/common/resource"
	"meta-egg-layout/internal/config"
	gormxmock "meta-egg-layout/pkg/gormx/mock"

	"github.com/golang/mock/gomock"
)

func GetResource(ctrl *gomock.Controller) *resource.Resource {
	// load config
	if err := config.LoadConfig("../../configs/conf-local.yml"); err != nil {
		panic(err)
	}
	cfg := config.GetResourceConfig()
	cfg.DB = nil
	rsrc, err := resource.InitResource(cfg)
	if err != nil {
		panic(err)
	}
	rsrc.DB = gormxmock.NewMockDB(ctrl)
	return rsrc
}
