package validparenthesis

var parenMp = map[byte]byte{
	')': '(',
	']': '[',
	'}': '{',
}

// IsValid will return whether the pair of parenthesis given in s is valid.
func IsValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	var i int
	parenQueue := make([]byte, 0, len(s))

	parenQueue = append(parenQueue, s[i])
	i++

	for i < len(s) {

		if openingMatch, isClosing := parenMp[s[i]]; len(parenQueue) > 0 && isClosing && openingMatch == parenQueue[len(parenQueue)-1] {
			parenQueue = parenQueue[:len(parenQueue)-1]
			i++
		} else {
			parenQueue = append(parenQueue, s[i])
			i++
		}

	}

	return len(parenQueue) == 0
}
