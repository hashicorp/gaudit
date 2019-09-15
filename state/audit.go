package state

import (
	"time"
)

type Audit struct {
	Repos   map[string]Repo
	Index   []string
	State   map[string]string
	Diffs   map[string]Diff
	Results map[string]Result
	Start   int64
	End     int64
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
}

type Diff struct {
	State  string
	Fields []Field
}

type Field struct {
	Name string
	Old  string
	New  string
}

type Result struct {
	Rules []Rule
}

type Rule struct {
	Name   string
	Status string
}
