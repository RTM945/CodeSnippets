package switcher

import (
	"ares/switcher/linker"
	"ares/switcher/provider"
)

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
