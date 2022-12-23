package util

import "os"

func GetHostName() string {
	if h, err := os.Hostname(); err == nil {
		return h
	}
	return ""
}
