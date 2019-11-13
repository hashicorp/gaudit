package commands

import (
	"fmt"

	"github.com/mmcquillan/gaudit/config"
	"github.com/mmcquillan/gaudit/state"
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
