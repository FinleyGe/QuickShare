package config

type ConfigStruct struct {
	Dev      bool     `mapstructure:"dev"`
	Server   server   `mapstructure:"server"`
	Database database `mapstructure:"database"`
}

type server struct {
	Port         string   `mapstructure:"port"`
	AllowOrigins []string `mapstructure:"allow_origins"`
}

type database struct {
	DBName   string `mapstructure:"db_name"`
	Address  string `mapstructure:"address"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}
