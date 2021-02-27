package mini

import "strings"



func New(tokens []string) *result {
	rest, flag := parse(tokens)
	return &result{
		rest: rest,
		flags: flag,
	}
}

// Deprecated: Will be removed when the package is published the release version.
func Raw(tokens []string) map[string]interface{} {
	results := make(map[string]interface{})

	rest, flag := parse(tokens)
	for k, v := range flag {
		results[k] = v
	}
	results["_"] = rest

	return results
}


func includePrefix(token string) bool { return strings.HasPrefix(token, "-") || strings.HasPrefix(token, "--") }

func parse(tokens []string) ([]string, map[string]interface{}) {
	rest := make([]string, 0)
	flag := make(map[string]interface{})

	for i := 0; i < len(tokens); i++ {
		if includePrefix(tokens[i]) {
			if i == len(tokens) - 1 || includePrefix(tokens[i + 1]) { // the arg that is the last or includes prefix `-`, `--`.
				flag[tokens[i]] = true
			} else {
				flag[tokens[i]] = tokens[i+1]
				i++
			}
		} else {
			rest = append(rest, tokens[i])
		}
	}

	return rest, flag
}
