package configurator

import (
	"RainbowRunner/cmd/configparser/configparser"
	"RainbowRunner/internal/types/drconfigtypes"
	"compress/zlib"
	"errors"
	"fmt"
	"github.com/goccy/go-json"
	"golang.org/x/exp/slices"
	"io"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func LoadFromCategoryConfigFile(path string) (map[string]*drconfigtypes.DRCategory, error) {
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

	drConfig := map[string]*drconfigtypes.DRCategory{}

	err = json.Unmarshal(data, &drConfig)

	if err != nil {
		return nil, err
	}

	return drConfig, nil
}

func LoadFromDumpedConfigFile(path string) (*drconfigtypes.DRConfig, error) {
	stat, err := os.Stat(path)

	if err != nil {
		return nil, err
	}

	if stat.IsDir() {
		return nil, errors.New(fmt.Sprintf("Config file path %s is a directory", path))
	}

	fp, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer fp.Close()

	zlibReader, err := zlib.NewReader(fp)

	if err != nil {
		return nil, err
	}

	defer zlibReader.Close()

	data, err := io.ReadAll(zlibReader)

	if err != nil {
		return nil, err
	}

	drConfig := drconfigtypes.NewDRConfig()

	err = json.Unmarshal(data, drConfig)

	if err != nil {
		return nil, err
	}

	return drConfig, nil
}

func LoadAllConfigurationFiles(rootDir string, extensions []string) (*drconfigtypes.DRConfig, error) {
	configsToload := make([]string, 0, 1024)

	err := filepath.WalkDir(rootDir, func(path string, d fs.DirEntry, err error) error {
		if filepath.Base(path) == "GCDictionary.txt" ||
			filepath.Base(filepath.Dir(path)) == "Migrators" ||
			strings.Contains(filepath.Dir(path), "\\help") ||
			strings.Contains(filepath.Dir(path), "\\fonts") ||
			strings.Contains(filepath.Dir(path), "\\effects\\2.0") {
			return nil
		}

		ext := filepath.Ext(path)
		if slices.Contains(extensions, ext) {
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
