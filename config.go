package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type config struct {
	Directory        string `json:"directory"`
	ConnectionString string `json:"connection_string"`
	Readonly         bool   `json:"readonly"`
}

func LoadConfig() (*config, error) {
	configFile, err := os.Open("./justatftpd.json")
	if err != nil {
		return nil, err
	}
	defer configFile.Close()

	bytes, _ := ioutil.ReadAll(configFile)
	var c config
	json.Unmarshal(bytes, &c)

	return &c, nil
}
