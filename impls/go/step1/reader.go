package main

import (
	"regexp"
)

type Reader struct {
	tokens   []string
	position int
}

func (r Reader) next() []string {
	return nil
}

func (r Reader) peek() []string {
	return nil
}

func read_str(s string) Reader {
	tokens := tokenize(s)
	return Reader{tokens, 0}
}

func tokenize(s string) []string {
	p := `[\\s,]*(~@|[\[\]{}()'` + "`" + `~^@]|"(?:\\.|[^\\"])*"?|;.*|[^\s\[\]{}('"` + "`" + `,;)]*)`
	re := regexp.MustCompile(p)
	return re.FindAllString(s, -1)
}