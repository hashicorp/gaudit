// Copyright IBM Corp. 2019, 2020
// SPDX-License-Identifier: MIT

package commands

import (
	"fmt"

	"github.com/hashicorp/gaudit/analyze"
	"github.com/hashicorp/gaudit/config"
	"github.com/hashicorp/gaudit/state"
)

func Analyze(options config.Options) {

	// get last audit
	fmt.Print("Loading state... ")
	audit, err := state.Load(options.Storage)
	if err != nil {
		fmt.Println("ERROR: " + err.Error())
	}
	fmt.Println("Loaded")

	// loading rules
	fmt.Print("Loading rules... ")
	rules, err := analyze.Load(options)
	if err != nil {
		fmt.Println("ERROR: " + err.Error())
	}
	fmt.Println("Loaded")

	// validating rules
	fmt.Print("Validating rules... ")
	err = analyze.Validate(options, rules)
	if err != nil {
		fmt.Println("ERROR: " + err.Error())
	}
	fmt.Println("Validated")

	// analyizing rules
	fmt.Println("Analyzing rules... ")
	err = analyze.Run(options, audit, rules)
	if err != nil {
		fmt.Println("ERROR: " + err.Error())
	}
	fmt.Println("Complete")

}
