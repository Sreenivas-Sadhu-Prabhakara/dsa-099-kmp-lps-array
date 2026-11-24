# KMP / LPS (pi) Array

> Striver SDE Sheet 099 · Day 16 String II — the failure function: on mismatch, fall back along the pattern's borders; the text pointer never retreats. Go, stdlib only, tested.

## Problem

Build the LPS array (longest proper prefix that is also a suffix, per pattern prefix) and use it for Knuth–Morris–Pratt search: find all pattern occurrences scanning the text exactly once.

## Approach

The insight: when a match of length k breaks, the text just read ends with the pattern's first k characters — so the longest *border* of that prefix (a prefix that's also a suffix) tells exactly how much of a new match is already in hand. `lps[k−1]` is that border's length; falling back along borders (`k = lps[k−1]`, possibly repeatedly) re-aligns the pattern without ever re-reading text. Building LPS uses the very same fallback on the pattern against itself — the function bootstraps itself, which is the part worth staring at until it clicks. After a full match the same fallback continues the scan, so overlapping occurrences come free.

**Complexity:** O(n+m) worst case — the adversarial inputs that break Rabin–Karp don't exist here. O(m) space.

## Run

    go test ./...   # verify
    go run .        # demo

---

One of 191 repos solving Striver's SDE Sheet in Go — stdlib only, one problem per repo, every solution tested. Browse the set: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
