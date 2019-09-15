package state

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/mmcquillan/gaudit/config"
)

func Load(options config.Options) (audit Audit, err error) {

	if _, err = os.Stat(options.Storage); err != nil {
		if os.IsNotExist(err) {
			return audit, nil
		}
	}

	b, err := ioutil.ReadFile(options.Storage)
	if err != nil {
		return audit, err
	}

	err = json.Unmarshal(b, &audit)
	if err != nil {
		return audit, err
	}

	return audit, nil

}
