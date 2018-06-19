package clif

import (
	"bytes"
	"strings"
)

func ParseListOutput(output *bytes.Buffer) []string {
	str := output.String()
	str = strings.Trim(str, "\t\r\n ")

	if len(str) == 0 {
		return []string{}
	}

	return strings.Split(str, "\n")
}
