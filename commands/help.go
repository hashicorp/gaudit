// Copyright IBM Corp. 2019, 2020
// SPDX-License-Identifier: MIT

package commands

import (
	"fmt"

	"github.com/hashicorp/gaudit/config"
)

func Help(options config.Options) {

	fmt.Print("gaudit")
	if options.Version != "" {
		fmt.Print(" (v " + options.Version + ")")
	}
	fmt.Println("")
	fmt.Println("")
	fmt.Println("  gaudit update - Update working list of github repos")
	fmt.Println("  gaudit list - List of working github repos")
	fmt.Println("  gaudit team <team> - List of repos available to a team")
	fmt.Println("  gaudit diff <old> <new> - Difference since last update")
	fmt.Println("  gaudit details [filter] - Detail list of github repos")
	fmt.Println("  gaudit analyze - Analyze rules against working repos")
	fmt.Println("  gaudit results [--verbose] - Show results of analysis")
	fmt.Println("  gaudit append - Creates or updates the append file")
	fmt.Println("  gaudit stats - Summarizes statistics data on the audit")
	fmt.Println("  gaudit csv - Outputs to csv format")
}
