package helpers

import "github.com/purwantogz/go-impact/config"

var responseConfig = config.JSONConstants("response")

//ResponseDict ResponseDict(thekey string) interface{}
func ResponseDict(thekey string) interface{} {
	var result = responseConfig.Get(thekey)
	return result
}
