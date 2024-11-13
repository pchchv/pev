package pev

const (
	seqNums      = "0123456789"
	seqKeyboard0 = "qwertyuiop"
	seqKeyboard1 = "asdfghjkl"
	seqKeyboard2 = "zxcvbnm"
	seqAlphabet  = "abcdefghijklmnopqrstuvwxyz"
)

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

func removeMoreThanTwoRepeatingChars(s string) string {
	var prev rune
	var prevPrev rune
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		r := runes[i]
		if r == prev && r == prevPrev {
			runes = deleteRuneAt(runes, i)
			i--
		}
		prev, prevPrev = r, prev
	}

	return string(runes)
}

func getReversedString(s string) string {
	var n int
	rune := make([]rune, len(s))
	for _, r := range s {
		rune[n] = r
		n++
	}

	rune = rune[0:n]
	// Reverse
	for i := 0; i < n/2; i++ {
		rune[i], rune[n-1-i] = rune[n-1-i], rune[i]
	}

	// Convert back to UTF-8.
	return string(rune)
}

func getLength(password string) int {
	password = removeMoreThanTwoRepeatingChars(password)
	password = removeMoreThanTwoFromSequence(password, seqNums)
	password = removeMoreThanTwoFromSequence(password, seqKeyboard0)
	password = removeMoreThanTwoFromSequence(password, seqKeyboard1)
	password = removeMoreThanTwoFromSequence(password, seqKeyboard2)
	password = removeMoreThanTwoFromSequence(password, seqAlphabet)
	password = removeMoreThanTwoFromSequence(password, getReversedString(seqNums))
	password = removeMoreThanTwoFromSequence(password, getReversedString(seqKeyboard0))
	password = removeMoreThanTwoFromSequence(password, getReversedString(seqKeyboard1))
	password = removeMoreThanTwoFromSequence(password, getReversedString(seqKeyboard2))
	password = removeMoreThanTwoFromSequence(password, getReversedString(seqAlphabet))

	return len(password)
}
