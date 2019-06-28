package util

import "os"

func DefaultGetEnvString(k string, v string) string {
	s := os.Getenv(k)
	if s == "" {
		return v
	}
	return s
}
