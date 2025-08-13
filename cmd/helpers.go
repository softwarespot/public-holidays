package cmd

import (
	"flag"
	"fmt"
	"math"
	"strconv"
)

// The naming has been taken from "pflag" i.e. used in "cobra". URL: https://pkg.go.dev/github.com/spf13/pflag#BoolVarP
func flagBoolVarP(p *bool, name, shorthand string, value bool, usage string) {
	flag.BoolVar(p, name, value, usage)
	flag.BoolVar(p, shorthand, value, usage)
}

func parseMaxConcurrency(s string) (int32, error) {
	v, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("invalid max concurrency of %q. Expected an integer of 32 bits", s)
	}
	if v < math.MinInt32 || v > math.MaxInt32 {
		return 0, fmt.Errorf("invalid max concurrency of %q. Expected an integer of 32 bits", s)
	}
	return int32(v), nil
}

func parseYear(s string) (int, error) {
	if len(s) != 4 {
		return 0, fmt.Errorf("invalid year of %q. Expected a valid 4 digit year", s)
	}

	year, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("invalid year of %q. Expected a valid 4 digit year", s)
	}
	if year <= 0 {
		return 0, fmt.Errorf("invalid year of %q. Expected a valid 4 digit year", s)
	}
	return year, nil
}
