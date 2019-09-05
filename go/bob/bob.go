// Package bob ...
// Bob.Hey accepts string and outputs Bob's reponse
package bob

import (
	"regexp"
	"strings"
)

// Hey func: sloppy but works
// replace periods, commas, and whitespace
// check for empty string
// control flow checks for all possible combinations
func Hey(remark string) string {
	remark = strings.Replace(remark, ".", "", -1)
	remark = strings.Replace(remark, ",", "", -1)
	remark = regexp.MustCompile(`[[:space:]]`).ReplaceAllString(remark, "")

	if remark == "" {
		return "Fine. Be that way!"
	}

	finalRune := remark[len(remark)-1]

	if finalRune == '?' && !regexp.MustCompile(`[[:alpha:]]`).MatchString(remark) {
		return "Sure."
	}

	if finalRune == '?' && strings.ToUpper(remark) == remark {
		return "Calm down, I know what I'm doing!"
	}

	if finalRune == '?' {
		return "Sure."
	}

	if finalRune == '!' && strings.HasPrefix(remark, "Let's") {
		return "Whatever."
	}

	if !regexp.MustCompile(`[[:alpha:]]`).MatchString(remark) {
		return "Whatever."
	}

	if finalRune == '!' || strings.ToUpper(remark) == remark {
		return "Whoa, chill out!"
	}

	if remark == "" {
		return "Fine. Be that way!"
	}

	return "Whatever."
}
