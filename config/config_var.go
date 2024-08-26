package config

import "os"

var ZLM_SERVICE_NAME = "zlmediakit"
var ZLM_SECRET = "c8026b69-f2ce-41e3-b7d3-1acdefa016b2" //zlmediakit/config.ini 中配置的admin_params secret,固定不变，可同时修改

func init() {
	if val, ok := os.LookupEnv("ZLM_SERVICE_NAME"); ok {
		ZLM_SERVICE_NAME = val
	}
	if val, ok := os.LookupEnv("ZLM_SECRET"); ok {
		ZLM_SECRET = val
	}
}
