package state

type Diff struct {
	State  string
	Fields []Field
}

type Field struct {
	Name string
	Old  string
	New  string
}
