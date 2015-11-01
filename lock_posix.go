// +build !windows

package singleinstance

import (
	"io/ioutil"
	"os"
	"strconv"
	"syscall"
)

// CreateLockFile try to create a file with given name
// and acquire an exclusive lock on it
// if the file already exists AND is still locked, it will fail
func CreateLockFile(filename string) (*os.File, error) {
	var (
		file *os.File
		err  error
	)

	if _, err = os.Stat(filename); os.IsNotExist(err) {
		// file doesnt exist, create
		file, err = os.Create(filename)
	} else {
		// file does exist, open
		file, err = os.OpenFile(filename, os.O_WRONLY, 0666)
	}

	if err != nil {
		return nil, err
	}

	err = syscall.Flock(int(file.Fd()), syscall.LOCK_EX|syscall.LOCK_NB)
	if err != nil {
		return nil, err
	}

	_, err = file.WriteString(strconv.Itoa(os.Getpid()))
	if err != nil {
		return nil, err
	}

	return file, nil
}

func GetLockFilePid(filename string) (pid int, err error) {
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}

	pid, err = strconv.Atoi(string(contents))
	return
}
