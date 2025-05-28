package phantom

import (
	"ares/configs"
	"ares/internel/discovery"
	"ares/logger"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strconv"
)

var LOGGER = logger.GetLogger("phantom http")

func Start() {
	cfg, err := configs.Load("ares/configs/config.yaml")
	if err != nil {
		panic(err)
	}
	_, res, err := discovery.NewRegistryAndResolver(cfg.Discovery)
	if err != nil {
		panic(err)
	}
	http.Handle("/Index", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		type rsp struct {
			Switchers []string `json:"switchers"`
		}
		list, err := res.List(r.Context(), fmt.Sprintf(cfg.Discovery.KeyPrefix, "switcher"))
		if err != nil || len(list) == 0 {
			LOGGER.Errorf("get list err: %v", err)
			bytes, _ := json.Marshal(rsp{Switchers: make([]string, 0)})
			_, _ = w.Write(bytes)
		} else {
			// 先随便搞一个
			switcher := list[0]
			bytes, _ := json.Marshal(rsp{Switchers: []string{switcher.Address}})
			_, _ = w.Write(bytes)
		}
	}))

	err = http.ListenAndServe(net.JoinHostPort(cfg.Service.Host, strconv.Itoa(cfg.Service.Port)), nil)
	if err != nil {
		panic(err)
	}
}
