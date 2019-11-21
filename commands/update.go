package commands

import (
	"fmt"

	"github.com/hashicorp/gaudit/config"
	"github.com/hashicorp/gaudit/state"
)

func Update(options config.Options) {

	// get new audit
	fmt.Print("Refreshing state... ")
	newAudit, err := state.Refresh(options)
	if err != nil {
		fmt.Println("ERROR: " + err.Error())
		return
	}
	fmt.Println("Refreshed")

	// backup
	err = state.Backup(options)
	if err != nil {
		fmt.Println("ERROR: " + err.Error())
		return
	}

	// save
	err = state.Save(options, newAudit)
	if err != nil {
		fmt.Println("ERROR: " + err.Error())
		return
	}

}
