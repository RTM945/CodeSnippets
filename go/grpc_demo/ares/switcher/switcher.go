package switcher

import (
	"ares/logger"
	ares "ares/pkg/io"
)

var LOGGER = logger.GetLogger("switcher")

var linker ares.INode
var provider ares.INode
