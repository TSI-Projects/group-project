package utils

import (
	"fmt"
	"strconv"
)

func StringToUint(s string) (uint, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("failed to convert string to uint type")
	}
	return uint(i), nil
}
