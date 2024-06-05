package coder

import "strings"

type Caesar struct {
	Alphabet []string
}

func NewCaesar(alphabet string) *Caesar {
	return &Caesar{Alphabet: strings.Split(alphabet, "")}
}

// EnglishAlphabet это модернизированый английский алфавит
// EnglishAlphabet is a modern English language alphabet
func (c *Caesar) doShiftedAlphabet(shift int) (shiftedAlphabet []string) {
	shiftedAlphabet = make([]string, len(c.Alphabet))
	if shift <= 0 {
		for shift < len(c.Alphabet) {
			shift = len(c.Alphabet) + shift
		}
	}
	for shift > len(c.Alphabet) {
		for shift > len(c.Alphabet) {
			shift = shift - len(c.Alphabet)
		}

	}
	for ak := range c.Alphabet {
		if (ak + shift) < len(c.Alphabet) {
			shiftedAlphabet[ak] = c.Alphabet[ak+shift]
		} else {
			shiftedAlphabet[ak] = c.Alphabet[ak+shift-len(c.Alphabet)]
		}
	}
	return
}
func (c *Caesar) findInAlphabet(alphabet []string, input string) int {
	for k, v := range alphabet {
		if v == input {
			return k
		}
	}
	return -1
}

func (c *Caesar) Decrypt(input string, shift uint) (ret string) {
	d := int(shift)
	shiftedAlphabet := c.doShiftedAlphabet(d)
	for _, v := range strings.Split(input, "") {
		position := c.findInAlphabet(shiftedAlphabet, v)
		if position != -1 {
			ret = ret + c.Alphabet[position]
		} else {
			ret = ret + v
		}
	}
	return
}
func (c *Caesar) Encrypt(input string, shift uint) (ret string) {
	d := int(shift)
	shiftedAlphabet := c.doShiftedAlphabet(d)
	for _, v := range strings.Split(input, "") {
		position := c.findInAlphabet(c.Alphabet, v)
		if position != -1 {
			ret = ret + shiftedAlphabet[position]
		} else {
			ret = ret + v
		}

	}
	return
}
