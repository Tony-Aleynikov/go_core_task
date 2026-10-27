package main

func intersection(s1, s2 []int) (bool, []int) {
	m1 := make(map[int]struct{}, len(s1))
	for _, v := range s1 {
		m1[v] = struct{}{}
	}

	var s3 []int
	for _, v := range s2 {
		if _, ok := m1[v]; ok {
			s3 = append(s3, v)
		}
	}

	return len(s3) > 0, s3
}
