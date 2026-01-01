// Copyright IBM Corp. 2019, 2020
// SPDX-License-Identifier: MIT

package state

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func Load(state string) (audit Audit, err error) {

	if _, err = os.Stat(state); err != nil {
		if os.IsNotExist(err) {
			return audit, nil
		}
	}

	b, err := ioutil.ReadFile(state)
	if err != nil {
		return audit, err
	}

	err = json.Unmarshal(b, &audit)
	if err != nil {
		return audit, err
	}

	return audit, nil

}
