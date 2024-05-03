package utils

import "os"

func DbPath() string {

	dir := os.Getenv("HOME") + "/task-program/"
	return dir
}