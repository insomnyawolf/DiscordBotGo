package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

var (
	//config Target Platform Values
	config *myConf
)

type myConf struct {
	//AppConfig
	Token string `description:"Token"`
}

// ConfigPath Defines where the configuration will be saved/written
const ConfigPath = "./AppConfig.json"

func loadConf() {
	defaultConf := myConf{}

	config = &myConf{}

	load(config, defaultConf)

	if defaultConf == *config {
		log.Fatalln("Please configure the app and start it again.")
	}
}

func load(output interface{}, defaultValues interface{}) error {
	_, err := os.Stat(ConfigPath)
	if err == nil {
		file, err := ioutil.ReadFile(ConfigPath)
		if err != nil {
			log.Fatalln(err)
		}
		err = json.Unmarshal(file, &output)
		if err != nil {
			log.Fatalln(err)
		}
		return nil
	}
	if os.IsNotExist(err) {
		config, err := json.MarshalIndent(defaultValues, "", "    ")
		if err != nil {
			return err
		}
		err = ioutil.WriteFile(ConfigPath, config, 777)
		if err != nil {
			return err
		}
	}
	return err
}
