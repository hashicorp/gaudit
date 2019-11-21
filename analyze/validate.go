package analyze

import (
	"errors"
	"strings"

	"github.com/hashicorp/gaudit/config"
)

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
