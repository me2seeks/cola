package config

type JWT struct {
	SigningKey  string   `mapstructure:"signing_key" json:"signing_key" yaml:"signing_key"`    // jwt签名
	ExpiresTime int64    `mapstructure:"expires_time" json:"expires_time" yaml:"expires_time"` // 过期时间
	NotBefore   int64    `mapstructure:"not_before" json:"not_before" yaml:"not_before"`       // 生效时间
	Issuer      string   `mapstructure:"issuer" json:"issuer" yaml:"issuer"`                   // 签发者
	Audience    []string `mapstructure:"audience" json:"audience" yaml:"audience"`             // 接收者
}
