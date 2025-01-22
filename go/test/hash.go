package test

import (
	"fmt"
	"github.com/cespare/xxhash/v2"
	"github.com/davecgh/go-spew/spew"
	"sync"
)

var hasherPool = sync.Pool{
	New: func() interface{} {
		return xxhash.New()
	},
}

var prettyPrintConfigForHash = &spew.ConfigState{
	Indent:                  " ",
	SortKeys:                true,
	DisableMethods:          true,
	SpewKeys:                true,
	DisablePointerAddresses: true,
	DisableCapacities:       true,
}

func Hash(key interface{}) int {
	hasher := hasherPool.Get().(*xxhash.Digest)
	defer hasherPool.Put(hasher)
	hasher.Reset()
	str := prettyPrintConfigForHash.Sprintf("%#v", key)
	fmt.Fprint(hasher, str)
	return int(hasher.Sum64())
}
