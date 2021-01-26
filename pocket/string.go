package pocket

func InStringArray(v string, vs []string) bool {
	for _, s := range vs {
		if v == s {
			return true
		}
	}
	return false
}
