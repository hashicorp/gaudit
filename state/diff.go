// Copyright IBM Corp. 2019, 2020
// SPDX-License-Identifier: MIT

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
