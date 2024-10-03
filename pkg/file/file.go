package file

import (
	"github.com/pelletier/go-toml/v2"
	u "sastoj/pkg/util"
)

type BaseConfigManager[T any] struct {
	location string
}

func (m *BaseConfigManager[T]) GetConfig(problemId int64) (config *T, err error) {
	tomlFile, err := u.ReadTomlFile(problemId, m.location)
	if err != nil {
		return nil, err
	}
	err = toml.Unmarshal(tomlFile, &config)
	if err != nil {
		return nil, err
	}
	return
}

func (m *BaseConfigManager[T]) SetConfig(problemId int64, config *T) error {
	tomlFile, err := toml.Marshal(config)
	if err != nil {
		return err
	}
	return u.WriteTomlFile(problemId, m.location, tomlFile)
}
