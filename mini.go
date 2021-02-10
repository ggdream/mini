package mini

import "strings"


func New(tokens []string) *result {
	rest, flag := parse(tokens)
	return &result{
		rest: rest,
		flags: flag,
	}
}

func Raw(tokens []string) map[string]interface{} {
	results := make(map[string]interface{})

	rest, flag := parse(tokens)
	for k, v := range flag {
		results[k] = v
	}
	results["_"] = rest

	return results
}

func parse(tokens []string) ([]string, map[string]string) {
	rest := make([]string, 0)
	flag := make(map[string]string)

	for i := 0; i < len(tokens); i++ {
		if strings.HasPrefix(tokens[i], "-") || strings.HasPrefix(tokens[i], "--") {
			if i == len(tokens) - 1 {
				flag[tokens[i]] = ""
			} else {
				flag[tokens[i]] = tokens[i+1]
				i += 1
				continue
			}
		} else {
			rest = append(rest, tokens[i])
		}
	}

	return rest, flag
}
