package config

type ConfigStruct struct {
	Env      string   `mapstructure:"env"`
	Server   server   `mapstructure:"server"`
	Database database `mapstructure:"database"`
	Redis    redis    `mapstructure:"redis"`
	Jwt      jwt      `mapstructure:"jwt"`
	File     file     `mapstructure:"file"`
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

type jwt struct {
	Secret  string `mapstructure:"secret"`
	Issuer  string `mapstructure:"issuer"`
	Expires int    `mapstructure:"expires"`
}

type file struct {
	AutoPath   string `mapstructure:"auto_path"`
	ManualPath string `mapstructure:"manual_path"`
}

type redis struct {
	Address  string `mapstructure:"address"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}
