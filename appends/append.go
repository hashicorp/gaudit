package appends

import ()

type Append struct {
	Name     string `yaml:"name"`
	Owner    string `yaml:"owner"`
	Category string `yaml:"category"`
	Notes    string `yaml:"notes"`
}
