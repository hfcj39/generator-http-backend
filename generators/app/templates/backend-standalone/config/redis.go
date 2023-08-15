package config

type Redis struct {
	DB       int    `mapstructure:"db" json:"db" yaml:"db"`
	Host     string `mapstructure:"host" json:"host" yaml:"host"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
}

type RedisKey struct {
	PersistentTimeFromKey string
}
