package utils

func Unique(s []string) []string {
	if len(s) == 0 {
		return s
	}
	lastUniq := 0
	for i := 1; i < len(s); i++ {
		if inArray(s[0:lastUniq+1], s[i]) {
			continue
		}
		lastUniq++
		if i == lastUniq {
			continue
		}
		s[lastUniq] = s[i]
	}
	return s[0 : lastUniq+1]
}

func inArray(haystack []string, needle string) bool {
	for _, s := range haystack {
		if s == needle {
			return true
		}
	}
	return false
}
