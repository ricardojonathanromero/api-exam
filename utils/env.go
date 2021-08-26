package utils

import "os"

func GetEnv(key, val string) string {
	v, ok := os.LookupEnv(key)
	if ok {
		return v
	}

	return val
}
