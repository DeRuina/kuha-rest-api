package env

import (
	"os"
	"strconv"
)

func GetString(key, defOption string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return defOption
	}
	return val
}

func GetInt(key string, defOption int) int {
	val, ok := os.LookupEnv(key)
	if !ok {
		return defOption
	}

	valInt, err := strconv.Atoi(val)
	if err != nil {
		return defOption
	}

	return valInt
}

func GetBool(key string, fallback bool) bool {
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	boolVal, err := strconv.ParseBool(val)
	if err != nil {
		return fallback
	}

	return boolVal
}
