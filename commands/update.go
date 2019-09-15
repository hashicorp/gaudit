package commands

import (
	"fmt"

	"github.com/mmcquillan/gaudit/config"
	"github.com/mmcquillan/gaudit/state"
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

	// get previous audit
	fmt.Print("Reloading prior state... ")
	oldAudit, err := state.Load(options)
	if err != nil {
		fmt.Println("ERROR: " + err.Error())
		return
	}
	fmt.Println("Loaded")

	// loop over new for diff
	fmt.Print("Comparing States...")
	diffs := state.Compare(oldAudit, newAudit)
	newAudit.Diffs = diffs
	fmt.Println("Compared")

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
