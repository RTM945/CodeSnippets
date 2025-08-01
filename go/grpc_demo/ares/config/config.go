package config

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
	PvId int    `yaml:"pvId"`
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

// EtcdServiceDiscoveryConfig Etcd service discovery config
type EtcdServiceDiscoveryConfig struct {
	Endpoints   []string      `mapstructure:"endpoints"`
	DialTimeout time.Duration `mapstructure:"dialtimeout"`
	Heartbeat   struct {
		TTL time.Duration `mapstructure:"ttl"`
		Log bool          `mapstructure:"log"`
	} `mapstructure:"heartbeat"`
	SyncServers struct {
		Interval time.Duration `mapstructure:"interval"`
	} `mapstructure:"syncservers"`
	Revoke struct {
		Timeout time.Duration `mapstructure:"timeout"`
	} `mapstructure:"revoke"`
	GrantLease struct {
		Timeout       time.Duration `mapstructure:"timeout"`
		MaxRetries    int           `mapstructure:"maxretries"`
		RetryInterval time.Duration `mapstructure:"retryinterval"`
	} `mapstructure:"grantlease"`
	Shutdown struct {
		Delay time.Duration `mapstructure:"delay"`
	} `mapstructure:"shutdown"`
}

// newDefaultEtcdServiceDiscoveryConfig Etcd service discovery default config
func newDefaultEtcdServiceDiscoveryConfig() *EtcdServiceDiscoveryConfig {
	return &EtcdServiceDiscoveryConfig{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
		Heartbeat: struct {
			TTL time.Duration `mapstructure:"ttl"`
			Log bool          `mapstructure:"log"`
		}{
			TTL: 60 * time.Second,
			Log: false,
		},
		SyncServers: struct {
			Interval time.Duration `mapstructure:"interval"`
		}{
			Interval: 120 * time.Second,
		},
		Revoke: struct {
			Timeout time.Duration `mapstructure:"timeout"`
		}{
			Timeout: 5 * time.Second,
		},
		GrantLease: struct {
			Timeout       time.Duration `mapstructure:"timeout"`
			MaxRetries    int           `mapstructure:"maxretries"`
			RetryInterval time.Duration `mapstructure:"retryinterval"`
		}{
			Timeout:       60 * time.Second,
			MaxRetries:    15,
			RetryInterval: 5 * time.Second,
		},
		Shutdown: struct {
			Delay time.Duration `mapstructure:"delay"`
		}{
			Delay: 300 * time.Millisecond,
		},
	}
}
