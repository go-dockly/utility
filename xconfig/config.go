package xconfig

import (
	"io/ioutil"

	"github.com/BurntSushi/toml"
	"github.com/pkg/errors"
)

// LoadConfig reads in a toml file and inits the ServiceConfig
func LoadConfig(cfg interface{}, path string) error {
	bytes, err := ioutil.ReadFile(path)

	if err != nil {
		return errors.Wrap(err, "unable to read file")
	}

	err = toml.Unmarshal(bytes, cfg)

	if err != nil {
		return errors.Wrapf(err, "error while parsing config file %s", string(bytes))
	}

	return nil
}
