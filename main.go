package main

import "fmt"

// LPS returns the longest-proper-prefix-which-is-also-suffix array: lps[i]
// is the length of the longest proper prefix of p[:i+1] that is also its
// suffix. This is KMP's failure function.
func LPS(p string) []int {
	lps := make([]int, len(p))
	k := 0 // length of the current matched prefix
	for i := 1; i < len(p); i++ {
		for k > 0 && p[i] != p[k] {
			k = lps[k-1] // fall back to the next-shorter border
		}
		if p[i] == p[k] {
			k++
		}
		lps[i] = k
	}
	return lps
}

// Search returns all indexes where pattern occurs in text, scanning the text
// once and never moving backwards: on a mismatch the pattern pointer falls
// back along the LPS borders instead.
func Search(text, pattern string) []int {
	out := []int{}
	m := len(pattern)
	if m == 0 || m > len(text) {
		return out
	}
	lps := LPS(pattern)
	k := 0
	for i := 0; i < len(text); i++ {
		for k > 0 && text[i] != pattern[k] {
			k = lps[k-1]
		}
		if text[i] == pattern[k] {
			k++
		}
		if k == m {
			out = append(out, i-m+1)
			k = lps[k-1] // keep going: overlapping matches allowed
		}
	}
	return out
}

func main() {
	fmt.Println(LPS("aabaaab"))
	fmt.Println(Search("aabaacaadaabaaba", "aaba"))
}
