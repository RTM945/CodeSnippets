package switcher

import (
	"ares/pkg/logger"
)

var LOGGER = logger.GetLogger("switcher")

var linker *Linker
var provider *Provider

func GetLinker() *Linker {
	return linker
}

func GetProvider() *Provider {
	return provider
}
