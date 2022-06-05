package adr

import (
	"io/ioutil"
	"os"
	"sync"

	"gopkg.in/yaml.v3"
)

const (
	configFilename = ".adr.yaml"
)

var (
	once   sync.Once
	config Config
)

type Config struct {
	RootDir string `yaml:"rootDir"`
}

func createConfig(rootDir string) error {
	cfg := Config{
		RootDir: rootDir,
	}
	d, err := yaml.Marshal(&cfg)
	if err != nil {
		return err
	}

	f, err := os.Create(configFilename)
	defer f.Close()
	if err != nil {
		return err
	}

	_, err = f.Write(d)
	if err != nil {
		return err
	}

	return nil
}

func getConfig() (Config, error) {
	var (
		err error
	)
	once.Do(func() {
		config, err = loadConfig()
	})
	if err != nil {
		return Config{}, err
	}

	return config, nil
}

func loadConfig() (Config, error) {
	body, err := ioutil.ReadFile(configFilename)
	if err != nil {
		return Config{}, err
	}

	c := Config{}
	if err := yaml.Unmarshal(body, &c); err != nil {
		return Config{}, err
	}

	return c, nil
}
