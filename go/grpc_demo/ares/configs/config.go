package configs

import (
	"gopkg.in/yaml.v3"
	"os"
	"time"
)

/*
service:
  name: app
  host: "0.0.0.0"
  port: 8080
discovery:
  type: etcd
  etcd:
    endpoints:
      - "localhost:2379"
    dialTimeout: 5s
    leaseTTL: 10
*/

type Config struct {
	Service   *Service   `yaml:"service"`
	Discovery *Discovery `yaml:"discovery"`
}

type Service struct {
	Name string `yaml:"name"`
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type Discovery struct {
	KeyPrefix string `yaml:"keyPrefix"`
	KeyFormat string `yaml:"keyFormat"`
	Type      string `yaml:"type"`
	Etcd      *Etcd  `yaml:"etcd"`
}

type Etcd struct {
	Endpoints   []string      `yaml:"endpoints"`
	DialTimeout time.Duration `yaml:"dialTimeout"`
	LeaseTTL    int64         `yaml:"leaseTTL"`
}

func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
