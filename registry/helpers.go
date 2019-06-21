package registry

func isIn(s string, in ...string) bool {
	for _, x := range in {
		if s == x {
			return true
		}
	}
	return false
}
