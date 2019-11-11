package env

import (
	"os"
	"strconv"
	"time"
)

func Int(key string, defaultValue int, desc string) int {
	v, err := strconv.Atoi(os.Getenv(key))
	if err != nil {
		return v
	}
	return defaultValue
}

func String(key string, defaultValue string, desc string) string {
	v := os.Getenv(key)
	if v != "" {
		return v
	}
	return defaultValue
}

func Bool(key string, defaultValue bool, desc string) bool {
	v, err := strconv.ParseBool(os.Getenv(key))
	if err != nil {
		return v
	}
	return defaultValue
}

func Duration(key string, defaultValue time.Duration, desc string) time.Duration {
	value := os.Getenv(key)
	if value != "" {
		v, err := strconv.ParseInt(value, 10, 64)
		if err == nil {
			return time.Duration(v)
		}
	}
	return defaultValue
}
