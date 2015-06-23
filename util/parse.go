package util

import "regexp"

func Parse(src interface{}, vars map[string]string) interface{} {
	switch src := src.(type) {
	default:
		return src
	case string:
		return ParseString(src, vars)
	case []string:
		return ParseArray(src, vars)
	case map[string]interface{}:
		return ParseMap(src, vars)
	}
}

func ParseString(src string, vars map[string]string) string {
	r, _ := regexp.Compile("{{([^{}]+)}}")
	return r.ReplaceAllStringFunc(src, func(s string) string {
		k := s[2 : len(s)-2]
		v, ok := vars[k]
		if !ok {
			return s
		}
		return v
	})
}

func ParseArray(src []string, vars map[string]string) []string {
	var results []string
	for _, e := range src {
		results = append(results, ParseString(e, vars))
	}
	return results
}

func ParseMap(src map[string]interface{}, vars map[string]string) map[string]interface{} {
	results := make(map[string]interface{})

	for k, v := range src {
		results[k] = Parse(v, vars)
	}
	return results
}