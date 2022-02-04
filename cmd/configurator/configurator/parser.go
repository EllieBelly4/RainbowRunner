package configurator

import (
	"RainbowRunner/cmd/configparser/configparser"
	"io/fs"
	"path/filepath"
	"strings"
)

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
		//[]string{"D:\\Work\\dungeon-runners\\666 dumps new\\avatar\\base\\Bank.txt"},
		[]string{"D:\\Work\\dungeon-runners\\666 dumps new\\items\\pal\\MageShieldPAL.txt", "D:\\Work\\dungeon-runners\\666 dumps new\\items\\modpal\\MageModPAL.txt"},
		//configsToload,
		rootDir,
	)

	all.MergeParents()

	return all, nil
}
