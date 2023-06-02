package core

import (
	"0049-server-go/configs"
	"0049-server-go/global"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

const ConfigFile = "settings.yaml"

// InitConf is the function to initialize the configuration in yaml file.
func InitConf() {
	c := &configs.Config{}
	yamlConf, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		panic(fmt.Errorf("get yamlConf err: %s", err))
	}

	err = yaml.Unmarshal(yamlConf, c)
	if err != nil {
		log.Fatalf("config Init Unmarshal: %v", err)
	}

	log.Println("config yamlFile load Init success.")
	// Assigning configuration files to global variables
	global.Config = c
}
