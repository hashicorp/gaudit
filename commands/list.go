package commands

import (
	"fmt"

	"github.com/hashicorp/gaudit/config"
	"github.com/hashicorp/gaudit/state"
)

func List(options config.Options) {

	// get list of repos
	audit, err := state.Load(options.Storage)
	if err != nil {
		fmt.Println("ERROR: " + err.Error())
	}

	// output
	for _, k := range audit.Index {
		fmt.Println(k)
	}

}
