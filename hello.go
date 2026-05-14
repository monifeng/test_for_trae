package main

func main() {
}

func mergeList(l1, l2 []int) []int {
	i1, i2 := 0, 0
	len1, len2 := len(l1), len(l2)
	maxLen := max(len1, len2)

	res := make([]int, 0, len1+len2)
	for i1 < maxLen && i2 < maxLen {
		if l1[i1] < l2[i2] {
			res = append(res, l1[i1])
			if i1 < len1 {
				i1++
			}
		} else if l1[i1] > l2[i2] {
			res = append(res, l2[i2])
			if i2 < len2 {
				i2++
			}
		} else {
			res = append(res, l1[i1])
			res = append(res, l2[i2])
			if i1 < len1 {
				i1++
			}
			if i2 < len2 {
				i2++
			}
		}
	}

	return res
}
