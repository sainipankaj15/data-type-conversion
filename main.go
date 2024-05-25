package easyconversion

import "strconv"

// Note : First letter of every funciton will be captial as we are exporting the functions
func StringToFloat64(input string) float64 {
	floatValue, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0.0
	}
	return floatValue
}

func Float64ToString(floatValue float64) string {
	return strconv.FormatFloat(floatValue, 'f', -1, 64)
}

func StringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return i
}

func IntToString(i int) string {
	return strconv.Itoa(i)
}

func StringToInt64(input string) int64 {
	intValue, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		return 0
	}
	return intValue
}

func IntToFloat64(i int) float64 {
	return float64(i)
}
