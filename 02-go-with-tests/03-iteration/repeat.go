package iteration

import "strings"

func Repeat(str string, n int) string {
	var repeated strings.Builder

	for i := 0; i < n; i++ {
		repeated.WriteString(str)
	}

	return repeated.String()
}
