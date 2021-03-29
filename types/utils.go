package types

import "strings"

func TrimComments(s string) string {
	var trimmed []string
	splitted := strings.Split(strings.TrimSpace(s), "\n")

	for _, s := range splitted {
		trimmed = append(trimmed, strings.TrimSpace(s))
	}

	return strings.Join(trimmed, "\n")
}
