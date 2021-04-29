package config

import "kubernetes-go-demo/config"

var conf *config.AppConfig

func SetAppConfig(c *config.AppConfig){
	conf = c
}

func GetAppConfig() *config.AppConfig {
	return conf
}