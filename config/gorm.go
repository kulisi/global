package config

type Gorm struct {
	Use   string `mapstructure:"Use" json:"Use" yaml:"Use"`
	Mssql Mssql  `json:"mssql" mapstructure:"mssql" yaml:"mssql"`
	Mysql Mysql  `json:"mysql" mapstructure:"mysql" yaml:"mysql"`
}
