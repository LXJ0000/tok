package utils

import "regexp"

func IsValidWithRegex(pattern, str string) bool { // TODO move to util
	regex := regexp.MustCompile(pattern)
	matched := regex.MatchString(str)
	return matched
}
