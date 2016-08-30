// +build !windows

package singleinstance

import (
	"os"
	"strconv"
	"syscall"
)

// CreateLockFile tries to create a file with given name and acquire an
// exclusive lock on it. If the file already exists AND is still locked, it will
// fail.
func CreateLockFile(filename string) (*os.File, error) {
	var (
		file *os.File
		err  error
	)

	if _, err = os.Stat(filename); os.IsNotExist(err) {
		// File doesn't exist, create it
		file, err = os.Create(filename)
	} else {
		// File does exist, open it
		file, err = os.OpenFile(filename, os.O_WRONLY, 0600)
	}

	if err != nil {
		return nil, err
	}

	err = syscall.Flock(int(file.Fd()), syscall.LOCK_SH|syscall.LOCK_NB)
	if err != nil {
		return nil, err
	}

	// Write PID to lock file
	_, err = file.WriteString(strconv.Itoa(os.Getpid()))
	if err != nil {
		return nil, err
	}

	return file, nil
}
