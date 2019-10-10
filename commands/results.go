package commands

import (
	"fmt"

	"github.com/mmcquillan/gaudit/config"
	"github.com/mmcquillan/gaudit/state"
)

func Results(options config.Options) {

	// get latest audit
	audit, err := state.Load(options.Storage)
	if err != nil {
		fmt.Println("ERROR: " + err.Error())
	}

	// output
	for _, r := range audit.Index {
		repo := audit.Repos[r]
		fmt.Println(repo.FullName)
		for _, rule := range audit.Results[r].Rules {
			if rule.Status == "error" || options.Args["verbose"] == "true" {
				fmt.Println("  [" + rule.Status + "] " + rule.Name)
			}
		}
	}

}
