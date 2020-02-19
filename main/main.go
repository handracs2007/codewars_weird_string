package main

import (
	"bytes"
	"fmt"
)

func toWeirdCase(str string) string {
	// Your code here and happy coding!
	var buffer bytes.Buffer
	upper := true

	for _, ch := range str {
		if ch == ' ' {
			upper = true
		}

		if upper {
			if ch >= 'a' && ch <= 'z' {
				ch = ch - 32
			}
		} else {
			if ch >= 'A' && ch <= 'Z' {
				ch = ch + 32
			}
		}

		upper = !upper
		buffer.WriteByte(byte(ch))
	}

	return buffer.String()
}

func main() {
	fmt.Println(toWeirdCase("Weird string case"))
}
