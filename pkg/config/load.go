package config

import (
	"fmt"
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

// Load returns configuration struct
func Load(path string) (*Configuration, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error reading configuration from file, %s", err)
	}

	var cfg = new(Configuration)
	if err := yaml.Unmarshal(bytes, cfg); err != nil {
		return nil, fmt.Errorf("unable to decode into struct, %v", err)
	}

	return cfg, nil
}
