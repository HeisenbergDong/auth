package config

type Config struct {
	System  System  `mapstructure:"system" json:"system" yaml:"system"`
	Zap     Zap     `mapstructure:"zap" json:"zap" yaml:"zap"`
	Mysql   Mysql   `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Redis   Redis   `mapstructure:"redis" json:"redis" yaml:"redis"`
	Auth    Auth    `mapstructure:"auth" json:"auth" yaml:"auth"`
	Grpc    Grpc    `mapstructure:"grpc" json:"grpc" yaml:"grpc"`
}
