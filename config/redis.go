package config

import "time"

type Redis struct {
	Host               string        `mapstructure:"host"`
	Port               int           `mapstructure:"port"`
	Password           string        `mapstructure:"password"`
	DB                 int           `mapstructure:"DB,omitempty"`
	MaxRetries         int           `mapstructure:"max_retries,omitempty"`
	MinRetryBackoff    time.Duration `mapstructure:"min_retry_backoff,omitempty"`
	MaxRetryBackoff    time.Duration `mapstructure:"max_retry_backoff,omitempty"`
	DialTimeout        time.Duration `mapstructure:"dial_timeout,omitempty"`
	ReadTimeout        time.Duration `mapstructure:"read_timeout,omitempty"`
	WriteTimeout       time.Duration `mapstructure:"write_timeout,omitempty"`
	PoolSize           int           `mapstructure:"pool_size,omitempty"`
	MinIdleConns       int           `mapstructure:"min_idle_conns,omitempty"`
	MaxConnAge         time.Duration `mapstructure:"max_conn_age,omitempty"`
	IdleTimeout        time.Duration `mapstructure:"idle_timeout,omitempty"`
	IdleCheckFrequency time.Duration `mapstructure:"idle_check_frequency,omitempty"`
}
