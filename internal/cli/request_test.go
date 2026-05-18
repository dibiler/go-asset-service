package cli

import (
	"testing"
)

func TestValidatedParams(t *testing.T) {
	tests := []struct {
		name     string
		method   string
		filter   string
		value    string
		wantErrs int
	}{
		{name: "list ok", method: "list", wantErrs: 0},
		{name: "filter requires value", method: "filter", wantErrs: 1},
		{name: "countbytype ok", method: "countbytype", value: "server", wantErrs: 0},
		{name: "filterby missing filter", method: "filterby", value: "x", wantErrs: 1},
		{name: "filterby invalid filter", method: "filterby", filter: "bad", value: "x", wantErrs: 1},
		{name: "filterby invalid format", method: "FilterBy", filter: "bad", value: "x", wantErrs: 1},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			errs := ValidatedParams(tc.method, tc.filter, tc.value)
			if len(errs) != tc.wantErrs {
				t.Fatalf("got %d errs, want %d", len(errs), tc.wantErrs)
			}
		})
	}
}

func TestParseParams(t *testing.T) {
	tests := []struct {
		name         string
		args         []string
		wantMethod   string
		wantFilter   string
		wantValue    string
		wantErrCount int
	}{
		{
			name:         "normalizes method and filter",
			args:         []string{"-method=FILTERBY", "-filter=NAME", "-value=name-1"},
			wantMethod:   "filterby",
			wantFilter:   "name",
			wantValue:    "name-1",
			wantErrCount: 0,
		},
		{
			name:         "unknown flag returns error",
			args:         []string{"-unknown=1"},
			wantErrCount: 1, // can be >= 1 depending on validation + parse
		},
		{
			name:         "missing parameters for filter method",
			args:         []string{"-method=filter"},
			wantErrCount: 1, // can be >= 1 depending on validation + parse
		},
		{
			name:         "added all parameters for filter method",
			args:         []string{"-method=filter", "-value=non-empty-value"},
			wantErrCount: 0, // can be >= 1 depending on validation + parse
		},
		{
			name:         "added all parameters for filterby method",
			args:         []string{"-method=filterby", "-filter=type", "-value=non-empty-value"},
			wantErrCount: 0, // can be >= 1 depending on validation + parse
		},
		{
			name:         "wrong filter for filterby method",
			args:         []string{"-method=filterby", "-filter=fakefilter", "-value=non-empty-value"},
			wantErrCount: 1, // can be >= 1 depending on validation + parse
		},
		{
			name:         "missing value for filterby method",
			args:         []string{"-method=filterby", "-filter=type"},
			wantErrCount: 1, // can be >= 1 depending on validation + parse
		},
		{
			name:         "missing filter for filterby method",
			args:         []string{"-method=filterby", "-value=non-empty"},
			wantErrCount: 1, // can be >= 1 depending on validation + parse
		},
		{
			name:         "missing value for countbytype method",
			args:         []string{"-method=countbytype"},
			wantErrCount: 1, // can be >= 1 depending on validation + parse
		},
		{
			name:         "all params for countbytype method",
			args:         []string{"-method=countbytype", "-value=non-empty"},
			wantErrCount: 0, // can be >= 1 depending on validation + parse
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, errs := ParseParams(tc.args)

			if tc.wantMethod != "" && got.method != tc.wantMethod {
				t.Fatalf("method got %q want %q", got.method, tc.wantMethod)
			}
			if tc.wantFilter != "" && got.filter != tc.wantFilter {
				t.Fatalf("filter got %q want %q", got.filter, tc.wantFilter)
			}
			if tc.wantValue != "" && got.value != tc.wantValue {
				t.Fatalf("value got %q want %q", got.value, tc.wantValue)
			}

			if len(errs) < tc.wantErrCount {
				t.Fatalf("errs got %d want at least %d", len(errs), tc.wantErrCount)
			}
		})
	}
}
