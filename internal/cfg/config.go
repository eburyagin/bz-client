package cfg

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Config struct {
	Datastore DatastoreConfig `json:"datastore"`
	Bus       BusConfig       `json:"bus"`
}

type DatastoreConfig struct {
	Type     string `json:"type"`
	Addr     string `json:"addr"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
}

type BusConfig struct {
	Type string `json:"type"`
	Urls string `json:"urls"`
}

func Load(cf string) (*Config, error) {
	var config Config
	raw, err := ioutil.ReadFile(cf)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	err = json.Unmarshal(raw, &config)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return &config, nil
}
