package registry

import "golang.org/x/exp/slices"

// CamelCase converts a snake_cased (_) identifier to an CamelCased identifier (hello_world -> HelloWorld)
// From the golang/protobuf
func CamelCase(s string) string {
	if s == "" {
		return ""
	}

	t := make([]byte, 0, 32)
	i := 0
	if s[0] == '_' {
		t = append(t, 'X')
		i++
	}
	for ; i < len(s); i++ {
		c := s[i]
		if c == '_' && i+1 < len(s) && isASCIILower(s[i+1]) {
			continue // Skip the underscore in s.
		}

		if isASCIIDigit(c) {
			t = append(t, c)
			continue
		}

		if isASCIILower(c) {
			c ^= ' ' // Make it a capital letter.
		}

		t = append(t, c) // Guaranteed not lower case.
		for i+1 < len(s) && isASCIILower(s[i+1]) {
			i++
			t = append(t, s[i])
		}
	}

	return string(t)
}

// ReservedNames are reserved by Protobuf and therefore cannot be field names (Protobuf generates methods with those  names).
var ReservedNames = []string{"Reset", "String", "ProtoMessage", "Descriptor"}

// Sanitize will apply CamelCase and rename Fields from the ReservedNames list following the scheme: Reset -> Reset_ and so on
func Sanitize(s string) string {
	s = CamelCase(s)
	if slices.Contains(ReservedNames, s) {
		return s + "_"
	}

	return s
}

func isASCIILower(c byte) bool {
	return 'a' <= c && c <= 'z'
}

func isASCIIDigit(c byte) bool {
	return '0' <= c && c <= '9'
}
