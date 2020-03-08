package config

//JWTConfig jwt config
type JWTConfig struct {
	Key    string `json:"key"`
	Issuer string `json:"issuer"`
}

//JwtConf global jwt key
var JwtConf *JWTConfig
