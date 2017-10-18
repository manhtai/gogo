package mumbling

import "strings"

// Accum is solution for http://www.codewars.com/kata/5667e8f4e3f572a8f2000039/
func Accum(s string) string {
	var result []string
	for index, r := range s {
		str := string(r)
		result = append(result, strings.ToUpper(str)+strings.Repeat(strings.ToLower(str), index))
	}
	return strings.Join(result, "-")
}
