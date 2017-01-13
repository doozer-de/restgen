package registry

import "strings"

// ParsePath parses the Path format for the router and extract the parameters in the path (from /City/:postal/Street/:name we extract postal and name).
func ParsePath(p string) []string {
	var parts []string
	for _, value := range strings.Split(p, "/") {
		if strings.HasPrefix(value, ":") {
			parts = append(parts, value[1:])
		}
	}
	return parts
}

func harmonizePathVars(p string) string {
	parts := strings.Split(p, "/")

	for i := range parts {
		if strings.HasPrefix(parts[i], ":") {
			parts[i] = ":id"
		}
	}

	return strings.Join(parts, "/")
}
