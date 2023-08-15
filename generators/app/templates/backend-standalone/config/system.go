package config

type System struct {
	Env               string `mapstructure:"env" json:"env" yaml:"env"`
	Addr              int    `mapstructure:"addr" json:"addr" yaml:"addr"`
	DbType            string `mapstructure:"db-type" json:"dbType" yaml:"db-type"`
	StaticPath        string `mapstructure:"static-path" json:"staticPath" yaml:"static-path"`
	WebUrl            string `mapstructure:"web-url" json:"web-url" yaml:"web-url"`
	ExceptionReceiver string `mapstructure:"exception-receiver" json:"exception-receiver" yaml:"exception-receiver"`
	Region            string `mapstructure:"region" json:"region" yaml:"region"`
}
