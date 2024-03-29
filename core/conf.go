package core

import (
	"0049-server-go/configs"
	"0049-server-go/global"
	"embed"
	"gopkg.in/yaml.v2"
	"io/fs"
	"log"
)

const ConfigFile = "settings.yaml"

// InitConf is the function to initialize the configuration in yaml file.
func InitConf(embedConfigFiles embed.FS) {
	c := &configs.Config{}
	//yamlConf, err := os.ReadFile(ConfigFile)
	//if err != nil {
	//	panic(fmt.Errorf("get yamlConf err: %s", err))
	//}

	yamlConf, err := fs.ReadFile(embedConfigFiles, ConfigFile)
	if err != nil {
		log.Fatalf("get yamlConf err: %v", err)
	}

	err = yaml.Unmarshal(yamlConf, c)
	if err != nil {
		log.Fatalf("config Init Unmarshal: %v", err)
	}

	log.Println("config yamlFile load Init success.")
	// Assigning configuration files to global variables
	global.Config = c
}
