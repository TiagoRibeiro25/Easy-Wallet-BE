package utils

import "strconv"

func ConvertStringToUint(str string) (uint, error) {
	nUint64, err := strconv.ParseUint(str, 10, 32)
	if err != nil {
		return 0, err
	}

	return uint(nUint64), nil
}
