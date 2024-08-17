package config

import (
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Cfg *Config

func init() {
	MustLoad()
}

func Load() error {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	viper.AutomaticEnv()
	// viper.SetEnvPrefix("APP")
	// bindEnvs()

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	viper.WatchConfig()

	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Println("config file changed:", e.Name)
		if err := viper.Unmarshal(&Cfg); err != nil {
			log.Panicln(err)
		}
	})

	return viper.Unmarshal(&Cfg)
}

func MustLoad() {
	err := Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
}

// func bindEnvs() {
// 	viper.BindEnv("host", "APP_HOST")
// 	viper.BindEnv("port", "APP_PORT")
// 	viper.BindEnv("mode", "APP_MODE")

// 	viper.BindEnv("redis.host", "REDIS_HOST")
// 	viper.BindEnv("redis.port", "REDIS_PORT")
// 	viper.BindEnv("redis.password", "REDIS_PASSWORD")
// 	viper.BindEnv("redis.database", "REDIS_DATABASE")

// 	viper.BindEnv("mysql.host", "MYSQL_HOST")
// 	viper.BindEnv("mysql.port", "MYSQL_PORT")
// 	viper.BindEnv("mysql.db_name", "MYSQL_DB_NAME")
// 	viper.BindEnv("mysql.username", "MYSQL_USERNAME")
// 	viper.BindEnv("mysql.password", "MYSQL_PASSWORD")
// 	viper.BindEnv("mysql.prefix", "MYSQL_PREFIX")
// 	viper.BindEnv("mysql.singular", "MYSQL_SINGULAR")
// 	viper.BindEnv("mysql.engine", "MYSQL_ENGINE")
// 	viper.BindEnv("mysql.max_idle_conns", "MYSQL_MAX_IDLE_CONNS")
// 	viper.BindEnv("mysql.max_open_conns", "MYSQL_MAX_OPEN_CONNS")
// 	viper.BindEnv("mysql.log_mode", "MYSQL_LOG_MODE")
// 	viper.BindEnv("mysql.log_zap", "MYSQL_LOG_ZAP")

// 	viper.BindEnv("jwt.signing_key", "JWT_SIGNING_KEY")
// 	viper.BindEnv("jwt.expires_time", "JWT_EXPIRES_TIME")
// 	viper.BindEnv("jwt.buffer_time", "JWT_BUFFER_TIME")
// 	viper.BindEnv("jwt.issuer", "JWT_ISSUER")
// }

type Config struct {
	Host string `mapstructure:"host" json:"host" yaml:"host"`
	Port int    `mapstructure:"port" json:"port" yaml:"port"`
	Mode string `mapstructure:"mode" json:"mode" yaml:"mode"`

	JWT   JWT   `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Redis Redis `mapstructure:"redis" json:"redis" yaml:"redis"`
	MySQL MySQL `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
}
