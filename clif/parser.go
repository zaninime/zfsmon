package clif

import (
	"bytes"
	"strings"
)

// ParseListOutput parses the output of commands like "zpool list ..."
func ParseListOutput(output *bytes.Buffer) []string {
	str := output.String()
	str = strings.Trim(str, "\t\r\n ")

	if len(str) == 0 {
		return []string{}
	}

	return strings.Split(str, "\n")
}
