package file

type FcConfigManager struct {
	BaseConfigManager[FcConfig]
}

type FcConfig struct {
	ReferenceAnswer string
	PartialScore    int16
}

// NewFcConfigManager create a new file manager
func NewFcConfigManager(fileLocation string) *FcConfigManager {
	return &FcConfigManager{
		BaseConfigManager: BaseConfigManager[FcConfig]{location: fileLocation},
	}
}
