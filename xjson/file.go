package xjson

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
)

// Load reads and verifies the contents of a json file
func Load(path string, v interface{}) error {

	afile, err := os.Open(path)
	if err != nil {
		return errors.Wrapf(err, "open %s failed", path)
	}
	defer afile.Close()

	abytes, err := ioutil.ReadAll(afile)
	if err != nil {
		return errors.Wrapf(err, "reading %s failed", path)
	}

	err = json.Unmarshal(abytes, &v)
	if err != nil {
		return errors.Wrapf(err, "unmarshalling %s failed", path)
	}
	return nil
}

// Write saves contents to a json file
func Write(path string, toWrite interface{}) error {

	file, err := json.MarshalIndent(toWrite, "", " ")
	if err != nil {
		return errors.Wrapf(err, "marshalling %s failed", path)
	}

	err = ioutil.WriteFile(path, file, 0644)
	if err != nil {
		return errors.Wrapf(err, "writing %s failed", path)
	}

	return nil
}
