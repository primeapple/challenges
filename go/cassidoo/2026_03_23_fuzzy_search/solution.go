package fuzzySearch

import "sort"

func FuzzySearch(text, pattern string, k int) []Result {
	m := len(pattern)
	if m == 0 {
		return []Result{}
	}

	// best maps start position → minimum errors found so far.
	best := make(map[int]int)

	// R[e][j] = true means: first j chars of pattern match text ending here with e errors.
	// D[e][j] = set of distinct deletion counts reachable at this state.
	//   Each distinct d gives a different start position: pos = i - (m - d) + 1.
	R := make([][]bool, k+1)
	D := make([][][]int, k+1)
	for e := range R {
		R[e] = make([]bool, m+1)
		D[e] = make([][]int, m+1)
		R[e][0] = true
		D[e][0] = []int{0}
	}

	// mergeD returns the union of two deletion-count slices (no duplicates).
	mergeD := func(a, b []int) []int {
		m := make(map[int]bool, len(a)+len(b))
		for _, v := range a {
			m[v] = true
		}
		for _, v := range b {
			m[v] = true
		}
		out := make([]int, 0, len(m))
		for v := range m {
			out = append(out, v)
		}
		return out
	}

	// addD returns ds with v added (if not already present).
	addD := func(ds []int, v int) []int {
		for _, x := range ds {
			if x == v {
				return ds
			}
		}
		return append(ds, v)
	}

	for i := 0; i < len(text); i++ {
		c := text[i]

		// Snapshot D before the main loop so we inherit the deletion pass
		// states from the previous iteration.
		prevD := make([][][]int, k+1)
		for e := range prevD {
			prevD[e] = make([][]int, m+1)
			for j := range prevD[e] {
				prevD[e][j] = append([]int(nil), D[e][j]...)
			}
		}

		// Main loop: right-to-left to avoid using this iteration's updates.
		for j := m; j >= 1; j-- {
			R[0][j] = false
			D[0][j] = nil
			if R[0][j-1] && c == pattern[j-1] {
				R[0][j] = true
				D[0][j] = append([]int(nil), prevD[0][j-1]...)
			}

			for e := 1; e <= k; e++ {
				R[e][j] = false
				D[e][j] = nil

				// Exact match inheriting from same error level
				if R[e][j-1] && c == pattern[j-1] {
					R[e][j] = true
					D[e][j] = mergeD(D[e][j], prevD[e][j-1])
				}
				// Substitution from e-1
				if R[e-1][j-1] {
					R[e][j] = true
					D[e][j] = mergeD(D[e][j], prevD[e-1][j-1])
				}
				// Insertion (text char consumed, pattern position unchanged) from e-1
				if R[e-1][j] {
					R[e][j] = true
					D[e][j] = mergeD(D[e][j], prevD[e-1][j])
				}
			}
		}

		// Deletion pass: left-to-right so deletions can chain.
		// A deletion skips one pattern char without consuming a text char (costs 1 error).
		// Repeat up to k times to allow up to k consecutive deletions.
		for range k {
			for e := 1; e <= k; e++ {
				for j := 1; j <= m; j++ {
					if R[e-1][j-1] {
						// Each d in D[e-1][j-1] produces d+1 here.
						newDs := []int(nil)
						for _, d := range D[e-1][j-1] {
							newDs = addD(newDs, d+1)
						}
						if len(newDs) > 0 {
							R[e][j] = true
							D[e][j] = mergeD(D[e][j], newDs)
						}
					}
				}
			}
		}

		// Record results: for each (e, d) at R[e][m], compute pos and keep minimum errors.
		for e := 0; e <= k; e++ {
			if R[e][m] {
				for _, d := range D[e][m] {
					pos := i - (m - d) + 1
					if pos >= 0 {
						if prev, exists := best[pos]; !exists || e < prev {
							best[pos] = e
						}
					}
				}
			}
		}
	}

	// Collect and sort results by position.
	results := make([]Result, 0, len(best))
	for pos, errors := range best {
		results = append(results, Result{position: pos, errors: errors})
	}
	sort.Slice(results, func(i, j int) bool {
		return results[i].position < results[j].position
	})

	return results
}

func ExactMatch(text, pattern string) int {
	m := len(pattern)
	if m == 0 {
		return -1
	}

	R := make([]bool, m+1)
	R[0] = true

	for i := 0; i < len(text); i++ {
		// Update the bit array
		for j := m; j >= 1; j-- {
			R[j] = R[j-1] && (text[i] == pattern[j-1])
		}

		if R[m] {
			return i - m + 1
		}
	}

	return -1
}
