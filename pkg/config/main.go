package config

import (
	"io/ioutil"
	"log"

	"github.com/performl/hibernate/pkg/utils"

	"gopkg.in/yaml.v2"
)

func InitConfig() *Configs {
	conf := &(Configs{})
	conf.initEnvs()
	conf.initSpecs()

	return conf
}

func (conf *Configs) init() {
	conf.initEnvs()
	conf.initSpecs()
}

func (conf *Configs) initEnvs() {
	envs := EnvMap{
		ConfigPath: utils.Getenv("KH_CONFIG_PATH", "config.yaml"),
	}
	conf.Envs = envs
}

func (conf *Configs) initSpecs() {
	var cfgMap ConfigMap

	data, err := ioutil.ReadFile(conf.Envs.ConfigPath)
	if err != nil {
		log.Fatal("File reading error", err)
	}

	err = yaml.Unmarshal(data, &cfgMap)
	if err != nil {
		log.Fatal("File reading error", err)
	}
	conf.Specs = cfgMap.Specs
}
