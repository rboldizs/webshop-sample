package config

import (
	"encoding/json"
	"os"
	"strconv"

	"github.com/golang/glog"
)

//Configuration interface provides configured values
type Configuration interface {
	InitConf(file string) error
	GetServerPort() uint16
	GetServerToken() string
}

type config struct {
	Server struct {
		Port  uint16 `json:"port"`
		Cert  string `json:"cert"`
		Token string `json:"token"`
	} `json:"server"`
}

//GetConfig returns factory configuration object
func GetConfig() Configuration {
	return &config{}
}

func (cfg *config) InitConf(file string) error {

	err := cfg.loadConfiguration(file)
	if err != nil {
		return err
	}

	return nil
}

func (cfg *config) loadConfiguration(file string) error {

	glog.Info("Processing config file: ", file)
	configFile, err := os.Open(file)
	defer configFile.Close()

	if err != nil {
		glog.Error("Failed to open config file", err)
		return err
	}

	jsonParser := json.NewDecoder(configFile)

	err = jsonParser.Decode(&cfg)
	if err != nil {
		glog.Error("Failed to process conf json: ", err)
		return err
	}

	token, exists := os.LookupEnv("SERVER_TOKEN")

	if exists {
		glog.Warningln("Warning env variable SERVER_TOKEN was set, overriding config file values.")
		cfg.Server.Token = token
	}

	port, exists := os.LookupEnv("SERVER_PORT")

	if exists {
		glog.Warningln("Warning env variable SERVER_PORT was set, overriding config file values.")
		cfg.Server.Port = 1323
		p, err := strconv.ParseUint(port, 10, 16)
		if err == nil {
			cfg.Server.Port = uint16(p)
		}
	}

	return nil
}

func (cfg *config) GetServerPort() uint16 {
	return cfg.Server.Port
}

func (cfg *config) GetServerToken() string {
	return cfg.Server.Token
}
