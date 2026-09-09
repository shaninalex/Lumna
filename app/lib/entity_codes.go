package utils

import (
	"fmt"
	"strings"
)

// TaskCode - return sequential unique code like: TSK-1, TSK-2 and so on.
func TaskCode(prTitle string, taskNumber int) string {
	if len(prTitle) < 3 {
		return fmt.Sprintf("%d", taskNumber)
	}
	prefix := prTitle[0:3]
	prefix = strings.ToUpper(prefix)
	return fmt.Sprintf("%s-%d", prefix, taskNumber)
}

func ProjectKey(prTitle string) string {
	prTitle = strings.ToUpper(prTitle)
	words := strings.Split(prTitle, " ")

	// if prTitle has more then 1 words - use uppercased first letters each of the words
	if len(words) > 1 {
		var output []string
		for _, w := range words {
			output = append(output, w[0:1])
		}
		return strings.Join(output, "")
	}

	// if prTitle has only 1 word - use first 3 uppercased letters
	if len(prTitle) < 2 {
		return prTitle
	}
	return prTitle[0:2]
}
