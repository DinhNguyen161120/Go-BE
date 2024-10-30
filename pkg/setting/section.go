package setting

type Config struct {
	Mysql MYSQLSetting `mapstructure:"mysql"`
}

type MYSQLSetting struct {
	Host             string `mapstructure:"host"`
	Port             int    `mapstructure:"port"`
	Username         string `mapstructure:"username"`
	Password         string `mapstructure:"password"`
	Dbname           string `mapstructure:"dbname"`
	MaxIdleConns     string `mapstructure:"maxIdleConns"`
	MaxOpenConns     string `mapstructure:"maxOpenConns"`
	ConnsMaxLifeTime string `mapstructure:"connMaxLifeTime"`
}
