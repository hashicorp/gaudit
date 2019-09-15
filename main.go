package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/mmcquillan/gaudit/commands"
	"github.com/mmcquillan/gaudit/config"
	"github.com/mmcquillan/matcher"
)

var version string

func main() {

	// options
	var options config.Options
	options.Version = version

	// Github Token
	if os.Getenv("GAUDIT_GITHUB_TOKEN") == "" {
		fmt.Println("ERROR: No GAUDIT_GITHUB_TOKEN defined")
		os.Exit(1)
	}
	options.GithubToken = os.Getenv("GAUDIT_GITHUB_TOKEN")

	// Organization
	if os.Getenv("GAUDIT_ORGANIZATION") == "" {
		fmt.Println("ERROR: No GAUDIT_ORGANIZATION defined")
		os.Exit(1)
	}
	options.Organization = os.Getenv("GAUDIT_ORGANIZATION")

	// Storage
	options.Storage = "gaudit.state"
	if os.Getenv("GAUDIT_STORAGE") != "" {
		options.Storage = os.Getenv("GAUDIT_STORAGE")
	}

	// Rules
	options.Rules = "rules.yml"
	if os.Getenv("GAUDIT_RULES") != "" {
		options.Rules = os.Getenv("GAUDIT_RULES")
	}

	// Debug
	options.Debug = false

	// interpret input
	match, _, arguments := matcher.Matcher("<bin> <command> [filter] [--]", strings.Join(os.Args, " "))
	if match {

		// debug
		if arguments["debug"] == "true" {
			options.Debug = true
		}

		// args
		options.Args = arguments

		// command
		switch arguments["command"] {
		case "list":
			commands.List(options)
		case "details":
			commands.Details(options)
		case "update":
			commands.Update(options)
		case "diff":
			commands.Diff(options)
		case "analyze":
			commands.Analyze(options)
		case "results":
			commands.Results(options)
		default:
			commands.Help(options)
		}
	} else {
		commands.Help(options)
	}

}
