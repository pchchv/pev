package pev

func deleteRuneAt(runes []rune, i int) []rune {
	if i >= len(runes) ||
		i < 0 {
		return runes
	}

	copy(runes[i:], runes[i+1:])
	runes[len(runes)-1] = 0
	runes = runes[:len(runes)-1]

	return runes
}

func removeMoreThanTwoFromSequence(s, seq string) string {
	var matches int
	runes := []rune(s)
	seqRunes := []rune(seq)
	for i := 0; i < len(runes); i++ {
		for j := 0; j < len(seqRunes); j++ {
			if i >= len(runes) {
				break
			}

			r, r2 := runes[i], seqRunes[j]
			if r != r2 {
				matches = 0
				continue
			}

			// found a match, advance the counter
			matches++
			if matches > 2 {
				runes = deleteRuneAt(runes, i)
			} else {
				i++
			}
		}
	}

	return string(runes)
}
