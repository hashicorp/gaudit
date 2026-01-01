// Copyright IBM Corp. 2019, 2020
// SPDX-License-Identifier: MIT

package analyze

import ()

type Rule struct {
	Name     string   `yaml:"name"`
	Action   string   `yaml:"action"`
	Resource string   `yaml:"resource"`
	Match    string   `yaml:"match"`
	Exclude  []string `yaml:"exclude"`
	Type     string   `yaml:"type"`
}
