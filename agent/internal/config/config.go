package config

import (
	"errors"
	"flag"
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Env        string           `yaml:"env"`
	CountCalcs int              `yaml:"count_calculators"`
	GRPCClient GRPCClientConfig `yaml:"grpc_client"`
	Durations  DurationsConfig  `yaml:"durations"`
}

type GRPCClientConfig struct {
	Addr         string `yaml:"orch_addr"`
	RetriesCount int    `yaml:"retries_count"`
}

type DurationsConfig struct {
	Plus  time.Duration `yaml:"plus"`
	Minus time.Duration `yaml:"minus"`
	Mult  time.Duration `yaml:"mult"`
	Del   time.Duration `yaml:"del"`
	Pow   time.Duration `yaml:"pow"`
}

func MustLoad() (*Config, error) {
	configPath := fetchConfigPath()
	if configPath == "" {
		return nil, errors.New("config path is empty")
	}

	return MustLoadPath(configPath)
}

func MustLoadPath(configPath string) (*Config, error) {
	// check if file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return nil, errors.New("config file does not exist: " + configPath)
	}

	var cfg Config

	file, err := os.Open(configPath)
	if err != nil {
		return nil, errors.New("cannot open config file: " + err.Error())
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		return nil, errors.New("cannot decode config: " + err.Error())
	}

	return &cfg, nil
}

// fetchConfigPath fetches config path from command line flag or environment variable.
// Priority: flag > env > default.
// Default value is empty string.
func fetchConfigPath() string {
	var res string

	flag.StringVar(&res, "config", "", "path to config file")
	// Note: flag.Parse() should be called in main() to avoid conflicts.

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res
}
