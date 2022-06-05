package adr

import (
	"fmt"
	"os"
	"time"
)

const (
	timeFormatInFilename = "20060102150405"
)

func CreateNewRecord(name string) error {
	config, err := getConfig()
	if err != nil {
		return err
	}

	f, err := os.Create(generateFilename(config.RootDir, name))
	defer f.Close()
	if err != nil {
		return err
	}

	_, err = f.Write(template)
	if err != nil {
		return err
	}

	return nil
}

func CreateConfigFile(rootDir string) error {
	if err := os.Mkdir(rootDir, 0777); err != nil {
		return err
	}
	return createConfig(rootDir)
}

func generateFilename(dir, name string) string {
	now := time.Now().Format(timeFormatInFilename)
	return fmt.Sprintf("%s/%s_%s.md", dir, now, name)
}
