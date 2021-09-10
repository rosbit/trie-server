package trie

import (
	"unicode"
	"bytes"
)

const (
	_letter = iota
	_digit
	_hz
	_ignore
)

func parse(s string) <-chan string {
	res := make(chan string)

	go func(res chan string) {
		w := &bytes.Buffer{}
		lastCate := _ignore

		var dumpWord = func(cate int) {
			if lastCate != cate {
				if w.Len() > 0 {
					res <- w.String()
					w.Reset()
				}
				lastCate = cate
			}
		}

		for _, ch := range s {
			switch {
			case ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z'):
				dumpWord(_letter)
				w.WriteRune(ch)
			case unicode.IsDigit(ch):
				dumpWord(_digit)
				w.WriteRune(ch)
			case unicode.In(ch, unicode.Han):
				dumpWord(_ignore)
				w.WriteRune(ch)
				res <- w.String()
				w.Reset()
			default:
				dumpWord(_ignore)
			}
		}

		dumpWord(_ignore)
		close(res)
	}(res)

	return res
}
