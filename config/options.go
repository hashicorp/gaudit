// Copyright IBM Corp. 2019, 2026
// SPDX-License-Identifier: MIT

package config

import ()

type Options struct {
	Version      string
	GithubToken  string
	Organization string
	Storage      string
	Policy       string
	Rules        string
	Append       string
	Debug        bool
	Args         map[string]string
}
