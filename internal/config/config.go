package config

import (
	"os"
	"path/filepath"
	kdl "github.com/sblinch/kdl-go"
)

type Config struct {
	BaseIndex int `kdl:"base-index"`
}

func GetConfig() (Config, error) {
	base, err := GetBaseDir();
	if err != nil {
		return Config{}, err;
	}

	config := Config{
		BaseIndex: 1, 
	}

	configPath := filepath.Join(base, "config.kdl");
	fileConfig, err := parseConfig(configPath);
	if err != nil {
		return Config{}, err;
	}

	if fileConfig.BaseIndex != 0 {
		config.BaseIndex = fileConfig.BaseIndex;
	}

	return config, nil;
}

func parseConfig(path string) (Config, error) {
	var c Config;
	dat, err := os.ReadFile(path);
	if err != nil {
		if os.IsNotExist(err) {
			return Config{}, nil;
		}
		return Config{}, err;
	}

	if err := kdl.Unmarshal(dat, &c); err != nil {
		return Config{}, err;
	}

	return c, nil;
}

func GetBaseDir() (string, error) {
	xdg := os.Getenv("XDG_CONFIG_HOME");

	if xdg == "" {
		home, err := os.UserHomeDir();
		if err != nil {
			return "", err;
		}
		xdg = filepath.Join(home, ".config");
	}
	path := filepath.Join(xdg, "muxt")
	if err := os.MkdirAll(path, 0755); err != nil {
		return "", err;
	}

	return path, nil;
}

func GetLayoutsDir() (string, error) {
	base, err := GetBaseDir();
	if err != nil {
		return "", err;
	}
	p := filepath.Join(base, "layouts/");
	return p, nil;
}

