// Copyright IBM Corp. 2019, 2020
// SPDX-License-Identifier: MIT

package appends

import ()

type Append struct {
	Name     string `yaml:"name"`
	Owner    string `yaml:"owner"`
	Category string `yaml:"category"`
	Notes    string `yaml:"notes"`
}
