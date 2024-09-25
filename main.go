package easyconversion

import (
	"strconv"
	"strings"
)

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

func Int64ToString(input int64) string {
	return strconv.FormatInt(input, 10)
}

func IntToFloat64(i int) float64 {
	return float64(i)
}

func Float64ToInt(f float64) int {
	return int(f)
}

func Int64ToInt(i64 int64) int {
	return int(i64)
}

func IntToInt64(i int) int64 {
	return int64(i)
}

func ParseBool(value string) bool {
	value = strings.ToLower(value)
	return value == "true" || value == "1"
}

func StringToInt32(input string) int32 {
	intValue, err := strconv.ParseInt(input, 10, 32)
	if err != nil {
		return 0
	}
	return int32(intValue)
}

func Int32ToString(i int32) string {
	return strconv.Itoa(int(i))
}
