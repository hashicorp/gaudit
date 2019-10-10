package commands

import (
	"fmt"

	"github.com/mmcquillan/gaudit/appends"
	"github.com/mmcquillan/gaudit/config"
	"github.com/mmcquillan/gaudit/state"
)

func Append(options config.Options) {

	// get existing append file
	fmt.Print("Loading append... ")
	appendList, err := appends.Load(options)
	if err != nil {
		fmt.Println("ERROR: " + err.Error())
	}
	fmt.Println("Loaded")

	// get existing audit
	fmt.Print("Loading state... ")
	audit, err := state.Refresh(options)
	if err != nil {
		fmt.Println("ERROR: " + err.Error())
		return
	}
	fmt.Println("Loaded")

	// run through audit
	fmt.Print("Updating appends... ")
	for _, repo := range audit.Repos {
		match := false
		for _, a := range appendList {
			if a.Name == repo.FullName {
				match = true
			}
		}
		if !match {
			appendList = append(appendList, appends.Append{
				Name: repo.FullName,
			})
		}
	}
	fmt.Println("Updaed")

	// save/overwrite append file
	fmt.Print("Saving appends... ")
	err = appends.Save(options, appendList)
	if err != nil {
		fmt.Println("ERROR: " + err.Error())
		return
	}
	fmt.Println("Saved")

}
