package config

type AppConfig struct {
	Version    string          `yaml:"version"`
	Env        string          `yaml:"env"`
	YamlDir   string      `yaml:"yamlDir"`
	PoolNum    int             `yaml:"poolNum"`
	Httpd      *HttpdConfig   `yaml:"httpd"`
	Log        *LogConfig       `yaml:"log"`
	Mysql   *MysqlConfig  `yaml:"mysql"`
	Redis      *RedisConfig     `yaml:"redis"`
}

type HttpdConfig struct {
	Host string
	Port int
}

type LogConfig struct {
	Dir       string
	Name      string
	Format    string
	RetainDay int8
	Level     string
}

type MysqlConfig struct {
	Host        string
	Port        int
	DbName      string
	User        string
	Password    string
	MaxConn int
	MaxOpen int
}

type RedisConfig struct {
	Host        string
	Port        int
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout int
}

type MachineryConfig struct {
	Broker        string
	Backend        string
	DefaultQueue    string
}