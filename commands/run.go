package commands

import (
	"fmt"

	"github.com/hashicorp/gaudit/config"
	"github.com/hashicorp/gaudit/state"
)

func Run(options config.Options) {

	// get previous state

	// refresh state
	newAudit, err := state.Refresh(options)
	if err != nil {
		fmt.Println("ERROR: " + err.Error())
		return
	}

	// update policy

	// run rules

	// diff

	// report changes

	// issue check

	// backup previous state

	// record new state

	//////////////
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
