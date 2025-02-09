package utils

import (
	"strings"

	"github.com/containers/common/pkg/secrets"
	"github.com/hanks177/podman/v4/pkg/util"
	"github.com/pkg/errors"
)

func IfPassesSecretsFilter(s secrets.Secret, filters map[string][]string) (bool, error) {
	result := true
	for key, filterValues := range filters {
		switch strings.ToLower(key) {
		case "name":
			result = util.StringMatchRegexSlice(s.Name, filterValues)
		case "id":
			result = util.StringMatchRegexSlice(s.ID, filterValues)
		default:
			return false, errors.Errorf("invalid filter %q", key)
		}
	}
	return result, nil
}
