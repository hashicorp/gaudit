// Copyright IBM Corp. 2019, 2020
// SPDX-License-Identifier: MIT

package commands

import (
	"fmt"

	"github.com/hashicorp/gaudit/config"
	"github.com/hashicorp/gaudit/state"
)

func Team(options config.Options) {

	// filter
	team := options.Args["team"]
	permission := options.Args["permission"]

	// get list of repos
	audit, err := state.Load(options.Storage)
	if err != nil {
		fmt.Println("ERROR: " + err.Error())
	}

	// output
	for _, k := range audit.Index {
		repo := audit.Repos[k]
		for _, t := range repo.Teams {
			if t.Name == team && (t.Permission == permission || permission == "") {
				fmt.Println(k + " (" + t.Permission + ")")
			}
		}
	}

}
