package state

import (
	"time"
)

type Audit struct {
	Repos     map[string]Repo
	Index     []string
	State     map[string]string
	Results   map[string]Result
	Timestamp int64
}

type Repo struct {
	ID            int64
	FullName      string
	Owner         string
	Name          string
	Description   string
	Language      string
	Topics        []string
	DefaultBranch string
	Private       bool
	Archived      bool
	Disabled      bool
	License       string
	Stargazers    int
	Forks         int
	Watchers      int
	Size          int
	Updated       time.Time
	Teams         []Team
	Policy        map[string]string
}

type Team struct {
	Name       string
	Permission string
}

type Result struct {
	Rules []Rule
}

type Rule struct {
	Name   string
	Status string
}
