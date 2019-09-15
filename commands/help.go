package commands

import (
	"fmt"

	"github.com/mmcquillan/gaudit/config"
)

func Help(options config.Options) {

	fmt.Print("gaudit")
	if options.Version != "" {
		fmt.Print(" (v " + options.Version + ")")
	}
	fmt.Println("")
	fmt.Println("")
	fmt.Println("  gaudit update [--debug] - Update working list of github repos")
	fmt.Println("  gaudit list - List of working github repos")
	fmt.Println("  gaudit diff - Difference since last update")
	fmt.Println("  gaudit details [filter] - Detail list of github repos")
	fmt.Println("  gaudit analyze - Analyze rules against working repos")
	fmt.Println("  gaudit results [--verbose] - Show results of analysis")

}
