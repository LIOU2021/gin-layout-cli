package helpers

import "os"

// check file or directory exists ?
func OsExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}
