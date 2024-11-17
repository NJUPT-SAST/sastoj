package file

import "github.com/pelletier/go-toml/v2"

type FcConfigManager struct {
	BaseConfigManager
}

type FcConfig struct {
	ReferenceAnswer string
	PartialScore    int16
}

// NewFcConfigManager create a new file manager
func NewFcConfigManager(fileLocation string) *FcConfigManager {
	return &FcConfigManager{
		BaseConfigManager: BaseConfigManager{location: fileLocation},
	}
}

func (m *FcConfigManager) GetConfig(problemId int64) (*FcConfig, error) {
	tomlFile, err := m.ReadFile(problemId)
	if err != nil {
		return nil, err
	}

	var config FcConfig
	err = toml.Unmarshal(tomlFile, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func (m *FcConfigManager) SetConfig(problemId int64, config FcConfig) error {
	tomlFile, err := toml.Marshal(config)
	if err != nil {
		return err
	}
	return m.WriteFile(problemId, tomlFile)
}
