package response

import "fmt"

func Highlight(text string) string {
	return fmt.Sprintf("\033[33m%s\033[0m", text)
}
