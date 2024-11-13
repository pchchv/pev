package pev

import "strings"

const (
	sepChars          = `_-., `
	replaceChars      = `!@$&*`
	digitsChars       = `0123456789`
	lowerChars        = `abcdefghijklmnopqrstuvwxyz`
	upperChars        = `ABCDEFGHIJKLMNOPQRSTUVWXYZ`
	otherSpecialChars = `"#%'()+/:;<=>?[\]^{|}~`
)

func getBase(password string) (base int) {
	chars := map[rune]struct{}{}
	for _, c := range password {
		chars[c] = struct{}{}
	}

	hasSep := false
	hasLower := false
	hasUpper := false
	hasDigits := false
	hasReplace := false
	hasOtherSpecial := false
	for c := range chars {
		switch {
		case strings.ContainsRune(replaceChars, c):
			hasReplace = true
		case strings.ContainsRune(sepChars, c):
			hasSep = true
		case strings.ContainsRune(otherSpecialChars, c):
			hasOtherSpecial = true
		case strings.ContainsRune(lowerChars, c):
			hasLower = true
		case strings.ContainsRune(upperChars, c):
			hasUpper = true
		case strings.ContainsRune(digitsChars, c):
			hasDigits = true
		default:
			base++
		}
	}

	if hasReplace {
		base += len(replaceChars)
	}

	if hasSep {
		base += len(sepChars)
	}

	if hasOtherSpecial {
		base += len(otherSpecialChars)
	}

	if hasLower {
		base += len(lowerChars)
	}

	if hasUpper {
		base += len(upperChars)
	}

	if hasDigits {
		base += len(digitsChars)
	}

	return
}
