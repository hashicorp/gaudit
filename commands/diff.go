package commands

import (
	"fmt"

	"github.com/mmcquillan/gaudit/config"
	"github.com/mmcquillan/gaudit/state"
)

func Diff(options config.Options) {

	// get latest audit
	audit, err := state.Load(options)
	if err != nil {
		fmt.Println("ERROR: " + err.Error())
	}

	// output
	for k, diff := range audit.Diffs {
		fmt.Println(diff.State + " " + k)
		for _, field := range diff.Fields {
			fmt.Println("  " + field.Name + ": " + field.Old + " => " + field.New)
		}
	}

}
