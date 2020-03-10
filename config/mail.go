package config

//MailConfig mail config
type MailConfig struct {
	Driver   string `json:"driver"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}

//MailConf global config mail
var MailConf *MailConfig
