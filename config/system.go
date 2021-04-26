package config

type System struct {
	Addr          int    `mapstructure:"addr" json:"addr" yaml:"addr"`
	DbType        string `mapstructure:"db-type" json:"dbType" yaml:"db-type"`
	GrpcAddress   string `mapstructure:"grpc-address" json:"grpcAddress" yaml:"grpc-address"`
}
