package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func READ(s string) string {
	return s
}

func EVAL(s string) string {
	return s
}

func PRINT(s string) string {
	return s
}

func rep(s string) string {
	return PRINT(EVAL(READ(s)))
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("user> ")
		text, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Print(rep(text))
	}
}
