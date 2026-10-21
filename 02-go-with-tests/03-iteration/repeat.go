package iteration

import "strings"

func Repeat(str string) string {
	var repeated strings.Builder

	for i := 0; i < 5; i++ {
		repeated.WriteString(str)
	}

	return repeated.String()
}
