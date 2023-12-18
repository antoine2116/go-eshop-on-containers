package utils

import (
	"strconv"
	"strings"
)

// StringToIntArray
// From : https://stackoverflow.com/questions/24972950/go-convert-strings-in-array-to-integer
func StringToIntArray(s string) ([]int, error) {
	sa := strings.Split(s, ",")

	si := make([]int, 0, len(sa))
	for _, a := range sa {
		i, err := strconv.Atoi(a)
		if err != nil {
			return si, err
		}
		si = append(si, i)
	}
	return si, nil
}
