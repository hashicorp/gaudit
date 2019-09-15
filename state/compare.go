package state

import (
	"strconv"
	"strings"
)

func Compare(oldAudit Audit, newAudit Audit) (diffs map[string]Diff) {

	// init diffs
	diffs = make(map[string]Diff)

	// loop over new for additions and changes
	for _, newRepo := range newAudit.Repos {

		// initialize new diff
		diff := Diff{
			State: "",
		}

		// find existing old repos by ID (in case name changed)
		oldRepoName := ""
		for _, oldRepo := range oldAudit.Repos {
			if oldRepo.ID == newRepo.ID {
				oldRepoName = oldRepo.FullName
			}
		}

		// compare old repo
		if oldRepoName != "" {

			// grab previous repo entry
			oldRepo := oldAudit.Repos[oldRepoName]

			if oldRepo.FullName != newRepo.FullName {
				diff.Fields = append(diff.Fields, Field{
					Name: "name",
					Old:  oldRepo.FullName,
					New:  newRepo.FullName,
				})
			}

			if oldRepo.Description != newRepo.Description {
				diff.Fields = append(diff.Fields, Field{
					Name: "description",
					Old:  oldRepo.Description,
					New:  newRepo.Description,
				})
			}

			if oldRepo.Language != newRepo.Language {
				diff.Fields = append(diff.Fields, Field{
					Name: "language",
					Old:  oldRepo.Language,
					New:  newRepo.Language,
				})
			}

			if strings.Join(oldRepo.Topics, ",") != strings.Join(newRepo.Topics, ",") {
				diff.Fields = append(diff.Fields, Field{
					Name: "topics",
					Old:  strings.Join(oldRepo.Topics, ","),
					New:  strings.Join(newRepo.Topics, ","),
				})
			}

			if oldRepo.Private != newRepo.Private {
				diff.Fields = append(diff.Fields, Field{
					Name: "private",
					Old:  strconv.FormatBool(oldRepo.Private),
					New:  strconv.FormatBool(newRepo.Private),
				})
			}

			if oldRepo.Archived != newRepo.Archived {
				diff.Fields = append(diff.Fields, Field{
					Name: "archived",
					Old:  strconv.FormatBool(oldRepo.Archived),
					New:  strconv.FormatBool(newRepo.Archived),
				})
			}

			if oldRepo.Disabled != newRepo.Disabled {
				diff.Fields = append(diff.Fields, Field{
					Name: "disabled",
					Old:  strconv.FormatBool(oldRepo.Disabled),
					New:  strconv.FormatBool(newRepo.Disabled),
				})
			}

			if oldRepo.License != newRepo.License {
				diff.Fields = append(diff.Fields, Field{
					Name: "license",
					Old:  oldRepo.License,
					New:  newRepo.License,
				})
			}

			if oldRepo.DefaultBranch != newRepo.DefaultBranch {
				diff.Fields = append(diff.Fields, Field{
					Name: "def branch",
					Old:  oldRepo.DefaultBranch,
					New:  newRepo.DefaultBranch,
				})
			}

			if oldRepo.Stargazers != newRepo.Stargazers {
				diff.Fields = append(diff.Fields, Field{
					Name: "stargazers",
					Old:  strconv.Itoa(oldRepo.Stargazers),
					New:  strconv.Itoa(newRepo.Stargazers),
				})
			}

			if oldRepo.Forks != newRepo.Forks {
				diff.Fields = append(diff.Fields, Field{
					Name: "forks",
					Old:  strconv.Itoa(oldRepo.Forks),
					New:  strconv.Itoa(newRepo.Forks),
				})
			}

			if oldRepo.Watchers != newRepo.Watchers {
				diff.Fields = append(diff.Fields, Field{
					Name: "watchers",
					Old:  strconv.Itoa(oldRepo.Watchers),
					New:  strconv.Itoa(newRepo.Watchers),
				})
			}

			if oldRepo.Size != newRepo.Size {
				diff.Fields = append(diff.Fields, Field{
					Name: "size",
					Old:  strconv.Itoa(oldRepo.Size),
					New:  strconv.Itoa(newRepo.Size),
				})
			}

			if oldRepo.Updated != newRepo.Updated {
				diff.Fields = append(diff.Fields, Field{
					Name: "updated",
					Old:  oldRepo.Updated.Format("2006-01-02 15:04:05 MST"),
					New:  newRepo.Updated.Format("2006-01-02 15:04:05 MST"),
				})
			}

			if len(diff.Fields) == 0 {
				diff.State = "="
			} else {
				diff.State = "*"
			}

		} else {
			diff.State = "+"
		}

		diffs[newRepo.FullName] = diff

	}

	// loop over old for deleted
	for _, oldRepo := range oldAudit.Repos {
		match := false
		for _, newRepo := range newAudit.Repos {
			if oldRepo.ID == newRepo.ID {
				match = true
			}
		}
		if !match {
			diff := Diff{
				State: "-",
			}
			diffs[oldRepo.FullName] = diff
		}
	}

	return diffs

}
