package utils

import "os"

func Int32Ptr(i int32) *int32 { return &i }

func Getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
