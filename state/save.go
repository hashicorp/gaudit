package state

import (
	"encoding/json"
	"io/ioutil"

	"github.com/mmcquillan/gaudit/config"
)

func Save(options config.Options, audit Audit) error {

	b, err := json.Marshal(audit)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(options.Storage, b, 0600)
	if err != nil {
		return err
	}

	return nil

}
