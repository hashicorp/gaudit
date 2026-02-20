// Copyright IBM Corp. 2019, 2020
// SPDX-License-Identifier: MIT

package appends

import (
	"io/ioutil"
	"os"

	"github.com/hashicorp/gaudit/config"
	"gopkg.in/yaml.v2"
)

func Load(options config.Options) (appends []Append, err error) {

	if _, err = os.Stat(options.Append); err != nil {
		if os.IsNotExist(err) {
			return appends, nil
		}
	}

	b, err := ioutil.ReadFile(options.Append)
	if err != nil {
		return appends, err
	}

	err = yaml.Unmarshal(b, &appends)
	if err != nil {
		return appends, err
	}

	return appends, nil

}
