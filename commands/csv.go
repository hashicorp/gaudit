package commands

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/mmcquillan/gaudit/analyze"
	"github.com/mmcquillan/gaudit/config"
	"github.com/mmcquillan/gaudit/state"
)

func CSV(options config.Options) {

	// get latest audit
	audit, err := state.Load(options.Storage)
	if err != nil {
		fmt.Println("ERROR: " + err.Error())
	}

	// get rules
	rules, err := analyze.Load(options)
	if err != nil {
		fmt.Println("ERROR: " + err.Error())
	}
	var rulesList []string
	for _, rule := range rules {
		rulesList = append(rulesList, rule.Name)
	}
	sort.Strings(rulesList)

	// header
	fmt.Print("\"ID\"")
	fmt.Print(",\"Name\"")
	fmt.Print(",\"Description\"")
	fmt.Print(",\"Language\"")
	fmt.Print(",\"Topics\"")
	fmt.Print(",\"Private\"")
	fmt.Print(",\"Archived\"")
	fmt.Print(",\"Disabled\"")
	fmt.Print(",\"License\"")
	fmt.Print(",\"Default Branch\"")
	fmt.Print(",\"Stars\"")
	fmt.Print(",\"Forks\"")
	fmt.Print(",\"Watchers\"")
	fmt.Print(",\"Size\"")
	fmt.Print(",\"Updated\"")
	for _, r := range rulesList {
		fmt.Print(",\"" + r + "\"")
	}
	fmt.Print("\n")

	// output
	for _, r := range audit.Index {
		repo := audit.Repos[r]
		fmt.Print("\"" + strconv.FormatInt(repo.ID, 10) + "\"")
		fmt.Print(",\"" + repo.FullName + "\"")
		fmt.Print(",\"" + strings.ReplaceAll(repo.Description, "\"", "\\\"") + "\"")
		fmt.Print(",\"" + repo.Language + "\"")
		fmt.Print(",\"" + strings.Join(repo.Topics, ",") + "\"")
		fmt.Print(",\"" + strconv.FormatBool(repo.Private) + "\"")
		fmt.Print(",\"" + strconv.FormatBool(repo.Archived) + "\"")
		fmt.Print(",\"" + strconv.FormatBool(repo.Disabled) + "\"")
		fmt.Print(",\"" + repo.License + "\"")
		fmt.Print(",\"" + repo.DefaultBranch + "\"")
		fmt.Print("," + strconv.Itoa(repo.Stargazers))
		fmt.Print("," + strconv.Itoa(repo.Forks))
		fmt.Print("," + strconv.Itoa(repo.Watchers))
		fmt.Print("," + strconv.Itoa(repo.Size))
		fmt.Print(",\"" + repo.Updated.Format("2006-01-02 15:04:05 MST") + "\"")
		for _, rule := range audit.Results[r].Rules {
			for _, r := range rulesList {
				if rule.Name == r {
					fmt.Print(",\"" + rule.Status + "\"")
				}
			}
		}
		fmt.Print("\n")
	}

}
