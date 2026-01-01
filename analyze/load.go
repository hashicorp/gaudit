// Copyright IBM Corp. 2019, 2020
// SPDX-License-Identifier: MIT

package analyze

import (
	"io/ioutil"
	"os"

	"github.com/hashicorp/gaudit/config"
	"gopkg.in/yaml.v2"
)

func Load(options config.Options) (rules []Rule, err error) {

	if _, err = os.Stat(options.Rules); err != nil {
		if os.IsNotExist(err) {
			return rules, nil
		}
	}

	b, err := ioutil.ReadFile(options.Rules)
	if err != nil {
		return rules, err
	}

	err = yaml.Unmarshal(b, &rules)
	if err != nil {
		return rules, err
	}

	return rules, nil

}
