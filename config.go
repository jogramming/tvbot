package main

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Token           string `json:"token"`
	Guild           string `json:"guild"`
	ReorderStartPos int    `json:"reorder_start_pos"`
	ReorderDeadZone int    `json:"reorder_dead_zone"`
}

func LoadConfig(path string) (conf *Config, err error) {
	var data []byte
	data, err = ioutil.ReadFile(path)
	if err != nil {
		return
	}

	conf = &Config{}
	err = json.Unmarshal(data, conf)
	return
}
