package config

import (
	appconfig "anis/config/app_config"
	dbconfig "anis/config/db_config"
)

func InitConfig(){
	appconfig.InitAppConfig()
	dbconfig.InitDBConfig()
}