package state

import (
	"context"
	"log"
	"sort"
	"time"

	"github.com/google/go-github/v25/github"
	"github.com/hashicorp/gaudit/config"
	"golang.org/x/oauth2"
)

func Refresh(options config.Options) (audit Audit, err error) {

	// init empty repos
	audit.Timestamp = time.Now().Unix()
	audit.Repos = make(map[string]Repo)

	// github client
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: options.GithubToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	// list options
	opt := &github.RepositoryListByOrgOptions{
		ListOptions: github.ListOptions{Page: 1},
	}

	// list all repositories for the authenticated user
	loop := true
	for loop {

		repos, resp, err := client.Repositories.ListByOrg(ctx, options.Organization, opt)
		if err != nil {
			return audit, err
		}

		for _, r := range repos {

			// debug output
			if options.Debug {
				log.Printf("\n\n%+v\n\n", r)
			}

			// cleanup some values
			description := ""
			if r.Description != nil {
				description = *r.Description
			}

			language := ""
			if r.Language != nil {
				language = *r.Language
			}

			license := ""
			if r.License != nil {
				license = *r.License.Name
			}

			// get collaborators
			/*
				users, _, err := client.Repositories.ListCollaborators(ctx, options.Organization, *r.Name, nil)
				if err != nil {
					return audit, err
				}
				if options.Debug {
					log.Printf("USERS: %+v", users)
				}
			*/

			// get teams
			teams, _, err := client.Repositories.ListTeams(ctx, options.Organization, *r.Name, nil)
			if err != nil {
				log.Print("ERROR: " + err.Error())
			}
			if options.Debug {
				log.Printf("\n\nTEAMS: %+v", teams)
			}
			var teamList []Team
			for _, t := range teams {
				teamList = append(teamList, Team{
					Name:       *t.Name,
					Permission: *t.Permission,
				})
			}

			// save record
			audit.Repos[*r.FullName] = Repo{
				ID:            *r.ID,
				FullName:      *r.FullName,
				Owner:         *r.Owner.Login,
				Name:          *r.Name,
				Description:   description,
				Language:      language,
				Topics:        r.Topics,
				DefaultBranch: *r.DefaultBranch,
				Private:       *r.Private,
				Archived:      *r.Archived,
				Disabled:      *r.Disabled,
				License:       license,
				Stargazers:    *r.StargazersCount,
				Forks:         *r.ForksCount,
				Watchers:      *r.WatchersCount,
				Size:          *r.Size,
				Updated:       r.UpdatedAt.Time,
				Teams:         teamList,
			}

		}

		// end of list
		if resp.NextPage == 0 {
			loop = false
		}

		// debug
		if options.Debug {
			log.Printf("%+v", resp)
		}

		// next list
		opt.ListOptions.Page = resp.NextPage

	}

	// sort list
	for k, _ := range audit.Repos {
		audit.Index = append(audit.Index, k)
	}
	sort.Strings(audit.Index)

	return audit, nil

}
