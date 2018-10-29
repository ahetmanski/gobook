// Echo simulates `echo` command.
// Contains 2 variants for comparison.
package echo

import (
	"strings"
)

func EchoSlow(args []string) string {
	var s, sep string
	for _, v := range args {
		s += sep + v
		sep = " "
	}
	return s
}

func EchoFast(args []string) string {
	return strings.Join(args, " ")
}
