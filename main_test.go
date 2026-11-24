package main

import (
	"reflect"
	"testing"
)

func TestLPS(t *testing.T) {
	cases := []struct {
		in   string
		want []int
	}{
		{"aabaaab", []int{0, 1, 0, 1, 2, 2, 3}},
		{"abcabc", []int{0, 0, 0, 1, 2, 3}},
		{"aaaa", []int{0, 1, 2, 3}},
		{"abcd", []int{0, 0, 0, 0}},
		{"", []int{}},
	}
	for _, c := range cases {
		if got := LPS(c.in); !reflect.DeepEqual(got, c.want) {
			t.Errorf("LPS(%q) = %v, want %v", c.in, got, c.want)
		}
	}
}

func TestKMPSearch(t *testing.T) {
	cases := []struct {
		text, pattern string
		want          []int
	}{
		{"aabaacaadaabaaba", "aaba", []int{0, 9, 12}},
		{"aaaa", "aa", []int{0, 1, 2}},
		{"abc", "d", []int{}},
		{"abc", "", []int{}},
		{"abababab", "abab", []int{0, 2, 4}},
	}
	for _, c := range cases {
		if got := Search(c.text, c.pattern); !reflect.DeepEqual(got, c.want) {
			t.Errorf("Search(%q,%q) = %v, want %v", c.text, c.pattern, got, c.want)
		}
	}
}
