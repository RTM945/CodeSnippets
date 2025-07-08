package switcher

import (
	"ares/logger"
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
