package analyze

import (
	"fmt"
	"strings"

	"github.com/hashicorp/gaudit/config"
)

func Validate(options config.Options, rules []Rule) error {

	valid := true
	invalidRules := []string{}
	for _, rule := range rules {

		// validate Action
		action := strings.ToLower(rule.Action)
		if !(action == "exists" || action == "not_exists" || action == "contains") {
			valid = false
			invalidRules = append(invalidRules, rule.Name)
		}
	}

	if !valid {
		return fmt.Errorf("Invalid Rules file; invalid rules: %s", invalidRules)
	}

	return nil

}
