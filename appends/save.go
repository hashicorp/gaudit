package appends

import (
	"io/ioutil"

	"github.com/mmcquillan/gaudit/config"
	"gopkg.in/yaml.v2"
)

func Save(options config.Options, appends []Append) error {

	b, err := yaml.Marshal(appends)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(options.Append, b, 0600)
	if err != nil {
		return err
	}

	return nil

}
