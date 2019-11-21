package commands

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/hashicorp/gaudit/analyze"
	"github.com/hashicorp/gaudit/appends"
	"github.com/hashicorp/gaudit/config"
	"github.com/hashicorp/gaudit/state"
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

	// get appends
	appendList, err := appends.Load(options)
	if err != nil {
		fmt.Println("ERROR: " + err.Error())
		return
	}

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
	fmt.Print(",\"Teams Admins\"")
	fmt.Print(",\"Teams Push\"")
	fmt.Print(",\"Teams Pull\"")
	for _, r := range rulesList {
		fmt.Print(",\"" + r + "\"")
	}
	fmt.Print(",\"Owner\"")
	fmt.Print(",\"Category\"")
	fmt.Print(",\"Notes\"")
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

		// teams
		fmt.Print(",\"")
		for _, t := range repo.Teams {
			if t.Permission == "admin" {
				fmt.Print(t.Name)
			}
		}
		fmt.Print("\"")
		fmt.Print(",\"")
		for _, t := range repo.Teams {
			if t.Permission == "push" {
				fmt.Print(t.Name)
			}
		}
		fmt.Print("\"")
		fmt.Print(",\"")
		for _, t := range repo.Teams {
			if t.Permission == "pull" {
				fmt.Print(t.Name)
			}
		}
		fmt.Print("\"")

		// rules
		for _, rule := range audit.Results[r].Rules {
			for _, r := range rulesList {
				if rule.Name == r {
					fmt.Print(",\"" + rule.Status + "\"")
				}
			}
		}
		match := false
		for _, a := range appendList {
			if a.Name == repo.FullName {
				match = true
				fmt.Print(",\"" + a.Owner + "\"")
				fmt.Print(",\"" + a.Category + "\"")
				fmt.Print(",\"" + a.Notes + "\"")
			}
		}
		if !match {
			fmt.Print(",\"\",\"\",\"\"")
		}
		fmt.Print("\n")
	}

}
