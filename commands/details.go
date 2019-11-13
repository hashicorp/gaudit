package commands

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/mmcquillan/gaudit/analyze"
	"github.com/mmcquillan/gaudit/appends"
	"github.com/mmcquillan/gaudit/config"
	"github.com/mmcquillan/gaudit/state"
)

func Details(options config.Options) {

	// get last audit
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

	// filter
	filter := ""
	if options.Args["filter"] != "" {
		filter = options.Args["filter"]
	}

	// output
	for _, k := range audit.Index {
		repo := audit.Repos[k]
		if filter == "" || strings.Contains(repo.FullName, filter) {
			fmt.Println(repo.FullName)
			fmt.Println("  id:          " + strconv.FormatInt(repo.ID, 10))
			fmt.Println("  description: " + repo.Description)
			fmt.Println("  language:    " + repo.Language)
			fmt.Println("  topics:      " + strings.Join(repo.Topics, ","))
			fmt.Println("  private:     " + strconv.FormatBool(repo.Private))
			fmt.Println("  archived:    " + strconv.FormatBool(repo.Archived))
			fmt.Println("  disabled:    " + strconv.FormatBool(repo.Disabled))
			fmt.Println("  license:     " + repo.License)
			fmt.Println("  def branch:  " + repo.DefaultBranch)
			fmt.Println("  stargazers:  " + strconv.Itoa(repo.Stargazers))
			fmt.Println("  forks:       " + strconv.Itoa(repo.Forks))
			fmt.Println("  watchers:    " + strconv.Itoa(repo.Watchers))
			fmt.Println("  size:        " + strconv.Itoa(repo.Size))
			fmt.Println("  updated:     " + repo.Updated.Format("2006-01-02 15:04:05 MST"))
			fmt.Print("  teams:       ")
			for _, t := range repo.Teams {
				fmt.Print(t.Name + " (" + t.Permission + ")  ")
			}
			fmt.Println("")
			for _, rule := range audit.Results[k].Rules {
				fmt.Println("  rule:        " + rule.Name + "=" + rule.Status)
			}
			for _, a := range appendList {
				if a.Name == repo.FullName {
					fmt.Println("  owner:       " + a.Owner)
					fmt.Println("  category:    " + a.Category)
					fmt.Println("  notes:       " + a.Notes)
				}
			}
			fmt.Print("\n")
		}
	}

}
