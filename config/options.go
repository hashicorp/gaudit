package config

import ()

type Options struct {
	Version      string
	GithubToken  string
	Organization string
	Storage      string
	Rules        string
	Append       string
	Debug        bool
	Args         map[string]string
}
