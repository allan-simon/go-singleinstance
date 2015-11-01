// +build windows

package singleinstance

import (
	"os"
)

// CreateLockFile try to create a file with given name
// and acquire an exclusive lock on it
// if the file already exists AND is still locked, it will fail
func CreateLockFile(filename string) (*os.File, error) {
	// if the files exists
	if _, err := os.Stat(filename); err == nil {
		// we first try to remove it
		err = os.Remove(filename)
		if err != nil {
			return nil, err
		}

	}
	// and we try to acquire an exclusive "lock on it"
	return os.OpenFile(filename, os.O_CREATE|os.O_EXCL|os.O_RDWR, 0666)
}
