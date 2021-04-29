package config

import "kubernetes-go-demo/config"

var conf *config.Config

func SetConfig(c *config.Config){
	conf = c
}

func GetConfig() *config.Config {
	return conf
}