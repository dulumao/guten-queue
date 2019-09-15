package env

import (
	"github.com/BurntSushi/toml"
	"path/filepath"
)

type tomlConfig struct {
	Queue queue
}

type queue struct {
	Dir                 string `toml:"dir"`
	NoQueueSleepMinutes int    `toml:"no_queue_sleep_minutes"`

	Redis struct {
		Addr     string `toml:"addr"`
		DbNumber int    `toml:"db_number"`
		Password string `toml:"password"`
	}
}

var Value *tomlConfig

func New() (error) {
	filePath, err := filepath.Abs("env.toml")

	if err != nil {
		return err
	}

	if _, err := toml.DecodeFile(filePath, &Value); err != nil {
		return err
	}

	return nil
}
