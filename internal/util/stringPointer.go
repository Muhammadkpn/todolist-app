package util

import "regexp"

func StringJson(s string) (jsonString string) {
	re := regexp.MustCompile(`\r?\n`)
	jsonString = re.ReplaceAllString(s, " ")
	return
}
