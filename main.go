package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	src := strings.Join(os.Args[1:], "")
	wc := countWords(src)
	fmt.Printf("%d\n", wc)
}

// returns count of words separated by any separator from wordSep slice
// consequitive separators counts as one
func countWords(src string, wordSep ...string) int {
	ws := map[string]struct{}{}
	for _, sep := range  wordSep {
		ws[sep] = struct{}{}
	}
	if len(wordSep) > 0 {
	        for _, sep := range  wordSep {
		        ws[sep] = struct{}{}
	        }
	}else {
		ws[" "] = struct{}{}
	}
	count := 0
	word := 0  // state of parser = number of non separators chars
	for _, r := range []rune(src) {
		if _, issep := ws[string(r)]; issep {
			if word > 0 {
				count ++
			}
			word = 0
		} else {
			word ++
		}
	}
	if word > 0 {
		count ++
	}

	return count
}
