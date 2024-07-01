package files

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func WriteFloatToFile(floatValue float64, fileName string) {
	floatValueStr := fmt.Sprint(floatValue)
	os.WriteFile(fileName, []byte(floatValueStr), 0644)
}

func ReadFloatFromFile(fileName string) (floatValue float64, err error) {
	data, err := os.ReadFile(fileName)

	if err != nil {
		return floatValue, errors.New("failed to read file")
	}

	floatValueStr := string(data)
	floatValue, err = strconv.ParseFloat(floatValueStr, 64)

	if err != nil {
		return floatValue, errors.New("failed to convert float")
	}
	return
}
