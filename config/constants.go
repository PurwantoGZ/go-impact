package config

import (
	"fmt"

	"github.com/spf13/viper"
)

//JSONTab json tabulasi
var JSONTab = "	"

//JSONConstants JSONConstants(configname string) *viper.Viper
func JSONConstants(configname string) *viper.Viper {
	var jsonconst = viper.New()
	jsonconst.SetConfigName(configname)
	jsonconst.AddConfigPath("constants")
	errjsonconst := jsonconst.ReadInConfig()
	if errjsonconst != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", errjsonconst))
	}
	return jsonconst
}
