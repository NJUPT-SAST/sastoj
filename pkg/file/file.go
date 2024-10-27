package file

import (
	"github.com/pelletier/go-toml/v2"
)

type BaseConfigManager[T any] struct {
	location string
}

func (m *BaseConfigManager[T]) GetConfig(problemId int64) (config *T, err error) {
	tomlFile, err := ReadTomlFile(problemId, m.location)
	if err != nil {
		return nil, err
	}
	err = toml.Unmarshal(tomlFile, &config)
	if err != nil {
		return nil, err
	}
	return
}

func (m *BaseConfigManager[T]) GetConfigString(problemId int64) (config string, err error) {
	tomlFile, err := ReadTomlFile(problemId, m.location)
	if err != nil {
		return "", err
	}
	return string(tomlFile), nil
}

func (m *BaseConfigManager[T]) SetConfig(problemId int64, config *T) error {
	tomlFile, err := toml.Marshal(config)
	if err != nil {
		return err
	}
	return WriteTomlFile(problemId, m.location, tomlFile)
}

func (m *BaseConfigManager[T]) SetConfigString(problemId int64, config string) error {
	tomlFile := []byte(config)
	return WriteTomlFile(problemId, m.location, tomlFile)
}

func NewBaseConfigManager(fileLocation string) *BaseConfigManager[any] {
	return &BaseConfigManager[any]{
		location: fileLocation,
	}
}
