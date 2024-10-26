package file

import (
	"os"
	"path/filepath"
	"strconv"
)

func ReadTomlFile(problemId int64, baseLocation string) ([]byte, error) {
	tomlFile, err := os.ReadFile(baseLocation + "/" + strconv.FormatInt(problemId, 10) + "/testdata/config.toml")
	if err != nil {
		return nil, err
	}
	return tomlFile, nil
}

func WriteTomlFile(problemId int64, baseLocation string, tomlFile []byte) error {
	filePath := baseLocation + "/" + strconv.FormatInt(problemId, 10) + "/testdata/config.toml"
	dirPath := filepath.Dir(filePath)
	err := os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		return err
	}
	err = os.WriteFile(filePath, tomlFile, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}
