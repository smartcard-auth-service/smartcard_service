package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"strings"
)

var (
	Cfg *GlobalConfig
)

type GlobalConfig struct {
	BASEDIR          string `json:"BASEDIR" binding:"required-non-empty"`
	MONGO_URI        string `json:"MONGO_URI" default:"mongodb://localhost:27017"`
	GRPC_LISTEN_HOST string `json:"GRPC_LISTEN_HOST" default:"8082"`
	TLS_CLIENT_CRT   string `json:"TLS_CLIENT_CRT" default:"{BASEDIR}/internal/tls/tls_client/client.crt"`
	TLS_CLIENT_KEY   string `json:"TLS_CLIENT_KEY" default:"{BASEDIR}/internal/tls/tls_client/client.key"`
	TLS_SERVER_CRT   string `json:"TLS_SERVER_CRT" default:"{BASEDIR}/internal/tls/tls_server/server.crt"`
	TLS_SERVER_KEY   string `json:"TLS_SERVER_KEY" default:"{BASEDIR}/internal/tls/tls_server/server.key"`
	TLS_URL          string `json:"TLS_URL" default:"https://localhost:1443"`
	TLS_LISTEN_HOST  string `json:"TLS_LISTEN_HOST" default:"1443"`
	LOG_LEVEL        string `json:"LOG_LEVEL" default:"TRACE"`
}

func InitGlobalConfig() {
	configPath := getConfigPath()
	_, err := os.Stat(configPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Config file cannot be load; err:%v\n", err)
		os.Exit(1)
	}
	Cfg = newCfg(configPath)
}

func getConfigPath() string {
	configPath, err := filepath.Abs(path.Join("config/main.json"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot execute filepath.Abs for config path; err:%v\n", err)
		os.Exit(1)
	}
	return configPath
}

func newCfg(filepath string) *GlobalConfig {
	conf := GlobalConfig{}
	err := conf.read(filepath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read config file %s; err:%v\n", filepath, err)
		os.Exit(1)
	}
	return &conf
}

func (globalConfig *GlobalConfig) read(filepath string) error {
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot read file %s; err: %v\n", filepath, err)
		return err
	}
	var bufferMap map[string]interface{}
	err = json.Unmarshal(content, &bufferMap)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to decode config; err: %v\n", err)
		return err
	}
	err = json.Unmarshal(content, &globalConfig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to decode config; err: %v\n", err)
		return err
	}
	globalConfig.setDefaultValues()
	v := reflect.ValueOf(&*globalConfig).Elem()
	replaceCfgVar(strings.NewReplacer("{BASEDIR}", globalConfig.BASEDIR), v)
	return nil
}

func replaceCfgVar(r *strings.Replacer, v reflect.Value) {
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		switch field.Kind() {
		case reflect.String:
			field.SetString(r.Replace(field.String()))
		case reflect.Ptr:
			if !field.IsNil() {
				replaceCfgVar(r, field.Elem())
			}
		case reflect.Struct:
			replaceCfgVar(r, field)
		}
	}
}

func (gF *GlobalConfig) setDefaultValues() {
	if gF.MONGO_URI == "" {
		gF.MONGO_URI = "mongodb://mongo:27017"
	}
	if gF.GRPC_LISTEN_HOST == "" {
		gF.GRPC_LISTEN_HOST = ":8082"
	}
	if gF.TLS_LISTEN_HOST == "" {
		gF.TLS_LISTEN_HOST = ":1443"
	}
	if gF.TLS_CLIENT_CRT == "" {
		gF.TLS_CLIENT_CRT = gF.BASEDIR + "/internal/tls/tls_server/client.crt"
	}
	if gF.TLS_CLIENT_KEY == "" {
		gF.TLS_CLIENT_KEY = gF.BASEDIR + "/internal/tls/tls_server/client.key"
	}
	if gF.TLS_SERVER_CRT == "" {
		gF.TLS_SERVER_CRT = gF.BASEDIR + "/internal/tls/tls_server/server.crt"
	}
	if gF.TLS_SERVER_KEY == "" {
		gF.TLS_SERVER_KEY = gF.BASEDIR + "/internal/tls/tls_server/server.key"
	}
	if gF.TLS_URL == "" {
		gF.TLS_URL = "https://localhost:1443"
	}
	if gF.LOG_LEVEL == "" {
		gF.LOG_LEVEL = "TRACE"
	}
}
