package str

import "strings"

func Trim(s string) string {
	return strings.TrimSpace(
		strings.ReplaceAll(s, "\n", ""),
	)
}
