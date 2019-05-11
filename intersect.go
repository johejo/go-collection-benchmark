package main

func Intersect(s1 []interface{}, s2 []interface{}) []interface{} {
	var capacity int
	if len(s1) <= len(s2) {
		capacity = len(s1)
	} else {
		capacity = len(s2)
	}

	intersect := make([]interface{}, 0, capacity)
	for i := range s1 {
		for j := range s2 {
			if s1[i] == s2[j] {
				intersect = append(intersect, s1[i])
				break
			}
		}
	}

	return intersect
}
