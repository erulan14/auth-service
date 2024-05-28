package config

type Server struct {
	Path    string `mapstructure:"PATH"`
	Secret  string `mapstructure:"SECRET"`
	Port    string `mapstructure:"PORT"`
	GinMode string `mapstructure:"GIN_MODE"`
}

type Db struct {
	Host string `mapstructure:"HOST"`
	Port string `mapstructure:"PORT"`
	User string `mapstructure:"USER"`
	Pass string `mapstructure:"PASS"`
	Name string `mapstructure:"NAME"`
}

type Env struct {
	AppEnv string `mapstructure:"APP_ENV"`
	Server Server `mapstructure:"SERVER"`
	Db     Db     `mapstructure:"DB"`
}
