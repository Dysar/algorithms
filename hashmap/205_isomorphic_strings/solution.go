package main

func isIsomorphic(s, t string) bool {
	mapST := make(map[string]string, 128) // Stores frequency of s
	mapTS := make(map[string]string, 128) // Stores frequency of t

	for i := 0; i < len(s); i++ {
		sCh, tCh := string(s[i]), string(t[i])

		// check if TS mapping contains the key. If yes, return false
		if val, ok := mapTS[tCh]; ok && val != sCh {
			return false
		} else if !ok {
			// map T to S
			mapTS[tCh] = sCh
		}

		// check if ST mapping contains the key S. If yer, return false. else, fill
		if val, ok := mapST[sCh]; ok && val != tCh {
			return false
		} else if !ok {
			// map S to T
			mapST[sCh] = tCh
		}
	}

	return true
}
