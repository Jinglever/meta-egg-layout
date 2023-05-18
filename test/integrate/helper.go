package integrate

import (
	"meta-egg-layout/internal/common/resource"
	"meta-egg-layout/internal/config"
)

func GetResource() *resource.Resource {
	// load config
	if err := config.LoadConfig("../../configs/conf-local.yml"); err != nil {
		panic(err)
	}
	rsrc, err := resource.InitResource(config.GetResourceConfig())
	if err != nil {
		panic(err)
	}
	return rsrc
}
