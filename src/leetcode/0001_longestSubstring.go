package main

import "fmt"

func exists(s string, r rune) int {
	// fmt.Printf("exists? %c in %v\n", r, s)
	i := 0
	for ; i < len(s); i++ {
		if rune(s[i]) == r {
			break
		}
	}
	// if we reached end and no match,
	// then we set i to -1
	if i == len(s) {
		i = -1
	}
	return i
}

func longestSubString(s string) string {
	i := 0
	j := i + 1
	mi := 0
	mj := 0
	for ; j < len(s); j++ {
		// check new char in the existing substring (SS) so far
		// fmt.Printf("SS: %v, i:%d, j:%d, s[i]:%c s[j]:%c\n", s[i:j], i, j, s[i], s[j])
		k := exists(s[i:j], rune(s[j]))
		// fmt.Println(">>>>k: ", k)
		// if we found a match
		if k != -1 {
			// print debug
			// if we found a long SS, update max indexes
			if (mj - mi) < (j - i) {
				mj, mi = j, i
			}
			// bring i (prev SS start) to match(new SS start)
			i += (k + 1)
		}
	}
	// // reached the end of string
	// // check to see if last stretch is longest
	if len(s) == j {
		// fmt.Printf("SS: %v, i:%d, j:%d, s[i]:%c s[j]:%c\n", s[i:j], i, j, s[i], s[j])
		if (mj - mi) < (j - i) {
			mj, mi = j, i
		}
	}

	return s[mi:mj]
}

func lengthOfLongestSubstring(s string) int {
	return len(longestSubString(s))
}

func tests() {
	arr_s := []string{
		"aaaaa",
		"dvdf",
		"",
		" ",
		"  ",
		"a",
		"abcd",
		"aaaabc",
		"aaaabcdefaa",
		"dvdf",
		"dvvvdddf",
		"dddaasdfasdfff1qszcd34edcv",
	}

	for i := 0; i < len(arr_s); i++ {
		// if i == 1 {
		// 	break
		// }
		s := arr_s[i]
		fmt.Printf("%v|%v|%d\n", s, longestSubString(s), lengthOfLongestSubstring(s))
	}
}

func main() {
	tests()
}
