package configurator

import (
	"RainbowRunner/cmd/configparser/configparser"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func LoadFromCategoryConfigFile(path string) (map[string]*configparser.DRCategory, error) {
	stat, err := os.Stat(path)

	if err != nil {
		return nil, err
	}

	if stat.IsDir() {
		return nil, errors.New(fmt.Sprintf("Config file path %s is a directory", path))
	}

	data, err := ioutil.ReadFile(path)

	if err != nil {
		return nil, err
	}

	drConfig := map[string]*configparser.DRCategory{}

	err = json.Unmarshal(data, &drConfig)

	if err != nil {
		return nil, err
	}

	return drConfig, nil
}

func LoadFromDumpedConfigFile(path string) (*configparser.DRConfig, error) {
	stat, err := os.Stat(path)

	if err != nil {
		return nil, err
	}

	if stat.IsDir() {
		return nil, errors.New(fmt.Sprintf("Config file path %s is a directory", path))
	}

	data, err := ioutil.ReadFile(path)

	if err != nil {
		return nil, err
	}

	drConfig := configparser.NewDRConfig()

	err = json.Unmarshal(data, drConfig)

	if err != nil {
		return nil, err
	}

	return drConfig, nil
}

func LoadAllConfigurationFiles(rootDir string) (*configparser.DRConfig, error) {
	configsToload := make([]string, 0, 1024)

	err := filepath.WalkDir(rootDir, func(path string, d fs.DirEntry, err error) error {
		if filepath.Base(path) == "GCDictionary.txt" ||
			filepath.Base(filepath.Dir(path)) == "Migrators" ||
			strings.Contains(filepath.Dir(path), "\\help") ||
			strings.Contains(filepath.Dir(path), "\\fonts") ||
			strings.Contains(filepath.Dir(path), "\\effects\\2.0") {
			return nil
		}

		if filepath.Ext(path) == ".txt" {
			configsToload = append(configsToload, path)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	all, err := configparser.ParseAllFilesToDRConfig(
		configsToload,
		rootDir,
	)

	return all, nil
}
