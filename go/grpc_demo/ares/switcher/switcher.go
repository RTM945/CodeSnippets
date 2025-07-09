package switcher

import (
	"ares/logger"
	"ares/switcher/linker"
	"ares/switcher/provider"
)

var (
	linkerRateMin int
	linkerRateMax int
	maxSession    int
)

var LOGGER = logger.GetLogger("switcher")

type Switcher struct {
	linker   *linker.Linker
	provider *provider.Provider
}

func GetInstance() *Switcher {
	return &Switcher{}
}

func (s *Switcher) GetLinker() *linker.Linker {
	return s.linker
}

func (s *Switcher) GetProvider() *provider.Provider {
	return s.provider
}
