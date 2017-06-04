package utils

import (
	"strconv"
	"strings"
)

func CutCommaAndTrimStrings(str string) ([]int, error) {
	var ids []int
	strs := strings.Split(str, ",")
	for _, str := range strs {
		idStr := strings.Trim(str, " ")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return ids, err
		}
		ids = append(ids, id)
	}
	return ids, nil
}
