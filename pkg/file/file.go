package file

type BaseConfigManager struct {
	location string
}

func (m *BaseConfigManager) ReadFile(problemId int64) (config []byte, err error) {
	return ReadTomlFile(problemId, m.location)
}

func (m *BaseConfigManager) ReadString(problemId int64) (config string, err error) {
	tomlFile, err := ReadTomlFile(problemId, m.location)
	if err != nil {
		return "", err
	}
	return string(tomlFile), nil
}

func (m *BaseConfigManager) WriteFile(problemId int64, config []byte) error {
	return WriteTomlFile(problemId, m.location, config)
}

func (m *BaseConfigManager) WriteString(problemId int64, config string) error {
	//tomlFile := []byte(config)
	return WriteTomlFile(problemId, m.location, []byte(config))
}
