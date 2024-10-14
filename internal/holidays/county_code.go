package holidays

import (
	"fmt"
	"strings"
)

type CountryCode string

func NewCountryCode(code string) (CountryCode, error) {
	code = strings.ToUpper(code)
	if len(code) != 2 {
		return "", fmt.Errorf("invalid county code of %q. Expected an ISO 3166-1 alpha-2 country code", code)
	}
	return CountryCode(code), nil
}
