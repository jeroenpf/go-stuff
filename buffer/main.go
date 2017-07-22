package main

import (
	"bytes"
	"fmt"
)

func main() {
	//	fmt.Println(intsToString([]int{1, 2, 3, 4, 5}))
	fmt.Println(len("123743479394833") % 3)
	fmt.Println(comma("1237434793943833"))
}

func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}

// 8823747734

func comma(s string) string {

	n := len(s)

	if n <= 3 {
		return s
	}

	var buf bytes.Buffer
	offset := n % 3
	if offset > 0 {
		buf.WriteString(s[0:offset] + ",")
	}
	for i := offset; i < n; i += 3 {
		buf.WriteString(s[i : i+3])
		if i+3 != n {
			buf.WriteString(",")
		}
	}

	return buf.String()

}
