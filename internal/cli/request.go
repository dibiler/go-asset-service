package cli

import (
	"errors"
	"flag"
	"fmt"
	"slices"
	"strings"
)

type Request struct {
	method string
	filter string
	value  string
}

var allowedFilterFields = []string{"name", "type", "location"}

var allowedMethods = []string{"list", "filter", "countbytype", "filterby"}

func ValidatedParams(method string, filter string, value string) []error {

	var errs []error

	if !slices.Contains(allowedMethods, method) {
		errs = append(errs, fmt.Errorf("method [%s] not found, allowed methods are: %s", method, strings.Join(allowedMethods, ", ")))
	}
	if (method == "countbytype" || method == "filter") && value == "" {
		errs = append(errs, fmt.Errorf("[%s] requires -value not empty", method))
	}
	if method != "filterby" {
		return errs
	}

	if value == "" || filter == "" {
		errs = append(errs, errors.New("[filterby] requires both -value and -filter not empty"))
	} else if !slices.Contains(allowedFilterFields, filter) {
		errs = append(errs, errors.New("[filterby] allows filter field to be one of: name, type, location"))
	}

	return errs
}

func ParseParams(args []string) (Request, []error) {
	fs := flag.NewFlagSet("cli", flag.ContinueOnError)

	cmdMethod := fs.String("method", "", "Method to execute on the assets, current available methods are: "+strings.Join(allowedMethods, ", ")+".")
	cmdFilter := fs.String("filter", "", "Filter to use in the filterBy method.")
	cmdValue := fs.String("value", "", "Value to use in the dynamic filter methods: filter, countByType and filterBy.")

	parseError := fs.Parse(args)

	filter := strings.ToLower(*cmdFilter)
	method := strings.ToLower(*cmdMethod)

	validated := ValidatedParams(method, filter, *cmdValue)

	if parseError != nil {
		validated = append(validated, parseError)
	}
	return Request{method: method, filter: filter, value: *cmdValue}, validated

}
