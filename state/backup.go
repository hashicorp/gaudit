package state

import (
	"io/ioutil"
	"os"

	"github.com/hashicorp/gaudit/config"
)

func Backup(options config.Options) error {

	if _, err := os.Stat(options.Storage); err != nil {
		if os.IsNotExist(err) {
			return nil
		}
	}

	b, err := ioutil.ReadFile(options.Storage)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(options.Storage+".bak", b, 0600)
	if err != nil {
		return err
	}

	return nil

}
