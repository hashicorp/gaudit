package commands

import (
	"fmt"

	"github.com/mmcquillan/gaudit/config"
	"github.com/mmcquillan/gaudit/state"
)

func List(options config.Options) {

	// get list of repos
	audit, err := state.Load(options)
	if err != nil {
		fmt.Println("ERROR: " + err.Error())
	}

	// output
	for _, k := range audit.Index {
		fmt.Println(k)
	}

}
