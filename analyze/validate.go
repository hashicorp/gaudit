package analyze

import (
	"errors"
	"strings"

	"github.com/hashicorp/gaudit/config"
)

// Validate iterates over a []Rule and validates rule actions.
// It returns an error if the rules file contains an invalid action.
func Validate(options config.Options, rules []Rule) error {

	valid := true
	for _, rule := range rules {

		// validate Action
		action := strings.ToLower(rule.Action)
		if !(action == "exists" || action == "not_exists" || action == "contains") {
			valid = false
		}

	}

	if !valid {
		return errors.New("Invalid Rules file")
	}

	return nil

}
