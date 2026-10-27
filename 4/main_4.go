package main

func diff(s1, s2 []string) []string {
	m2 := make(map[string]struct{}, len(s2))
	for _, s := range s2 {
		m2[s] = struct{}{}
	}

	var s3 []string
	for _, s := range s1 {
		if _, ok := m2[s]; !ok {
			s3 = append(s3, s)
		}
	}

	return s3
}
