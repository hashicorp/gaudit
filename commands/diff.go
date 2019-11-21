package commands

import (
	"fmt"

	"github.com/hashicorp/gaudit/config"
	"github.com/hashicorp/gaudit/state"
)

func Diff(options config.Options) {

	oldAudit, err := state.Load(options.Args["old"])
	if err != nil {
		fmt.Println("ERROR: " + err.Error())
	}

	newAudit, err := state.Load(options.Args["new"])
	if err != nil {
		fmt.Println("ERROR: " + err.Error())
	}

	diffs := state.Compare(oldAudit, newAudit)

	for k, diff := range diffs {
		if diff.State != "=" || options.Args["verbose"] == "true" {
			fmt.Println(diff.State + " " + k)
			for _, field := range diff.Fields {
				fmt.Println("  " + field.Name + ": " + field.Old + " => " + field.New)
			}
		}
	}

}
