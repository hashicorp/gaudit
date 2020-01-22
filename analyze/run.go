package analyze

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/go-github/v25/github"
	"github.com/hashicorp/gaudit/config"
	"github.com/hashicorp/gaudit/state"
	"golang.org/x/oauth2"
)

// Run iterates over a list of repos and validates that the rule is followed.
func Run(options config.Options, audit state.Audit, rules []Rule) error {

	// github client
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: options.GithubToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	// check audit results
	if audit.Results == nil {
		audit.Results = make(map[string]state.Result)
	}

	// loop each repo, for each rule
	for _, k := range audit.Index {
		repo := audit.Repos[k]

		// check if already done
		if _, ok := audit.Results[repo.FullName]; !ok {

			fmt.Println(repo.FullName)
			result := state.Result{}

			for _, rule := range rules {

				if rule.Type == "" || (rule.Type == "public" && !repo.Private) || (rule.Type == "private" && repo.Private) {

					if rule.Action == "exists" {
						_, _, _, err := client.Repositories.GetContents(context.Background(), repo.Owner, repo.Name, rule.Resource, nil)
						if err != nil {
							result.Rules = append(result.Rules, state.Rule{
								Name:   rule.Name,
								Status: "error",
							})
						} else {
							result.Rules = append(result.Rules, state.Rule{
								Name:   rule.Name,
								Status: "success",
							})
						}
					}

					if rule.Action == "not_exists" {
						_, _, _, err := client.Repositories.GetContents(context.Background(), repo.Owner, repo.Name, rule.Resource, nil)
						if err != nil {
							result.Rules = append(result.Rules, state.Rule{
								Name:   rule.Name,
								Status: "success",
							})
						} else {
							result.Rules = append(result.Rules, state.Rule{
								Name:   rule.Name,
								Status: "error",
							})
						}
					}

					if rule.Action == "contains" {
						resp, _, _, err := client.Repositories.GetContents(context.Background(), repo.Owner, repo.Name, rule.Resource, nil)
						if err != nil {
							result.Rules = append(result.Rules, state.Rule{
								Name:   rule.Name,
								Status: "error",
							})
						} else {
							if strings.Contains(*resp.Content, rule.Match) {
								result.Rules = append(result.Rules, state.Rule{
									Name:   rule.Name,
									Status: "success",
								})
							} else {
								result.Rules = append(result.Rules, state.Rule{
									Name:   rule.Name,
									Status: "error",
								})
							}
						}
					}
				} else {
					result.Rules = append(result.Rules, state.Rule{
						Name:   rule.Name,
						Status: "na",
					})
				}
			}

			// save the result for this repo
			audit.Results[repo.FullName] = result
			err := state.Save(options, audit)
			if err != nil {
				fmt.Println("ERROR: " + err.Error())
			}

		}

	}

	return nil

}
