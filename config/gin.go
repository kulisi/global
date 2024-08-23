package config

type Gin struct {
	Use  bool   `mapstructure:"use" json:"use" yaml:"use"`
	Addr string `mapstructure:"addr" json:"addr" yaml:"addr"`
}
