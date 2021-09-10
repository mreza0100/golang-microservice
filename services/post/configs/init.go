package configs

import (
	"freeSociety/configs"
)

const serviceName = "post"

var Configs *configs.ServiceConfigs

func init() {
	Configs = new(configs.ServiceConfigs)
	Configs.Name = serviceName
	Configs.SetConfigFile()
}
