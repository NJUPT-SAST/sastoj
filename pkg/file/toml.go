package file

import (
	"os"
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
	err := os.WriteFile(baseLocation+"/"+strconv.FormatInt(problemId, 10)+"/testdata/config.toml", tomlFile, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}
