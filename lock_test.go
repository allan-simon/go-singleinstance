package singleinstance

import (
	"io/ioutil"
	"os"
	"testing"
)

func createTestLock(t *testing.T) *os.File {
	tmpfile, err := ioutil.TempFile("", "singleinstance")
	if err != nil {
		t.Fatal("Cannot create temporary file:", err)
	}
	tmpfile.Close()

	f, err := CreateLockFile(tmpfile.Name())
	if err != nil {
		t.Fatal("Expected no error while creating lock, got:", err)
	}

	return f
}

func TestCreateLockFile(t *testing.T) {
	f := createTestLock(t)
	defer os.Remove(f.Name())

	_, err := CreateLockFile(f.Name())
	if err == nil {
		t.Fatal("Expected an error while trying to lock, got:", err)
	}

	f.Close() // Remove the lock

	f, err = CreateLockFile(f.Name())
	if err != nil {
		t.Fatal("Expected no error while trying to lock, got:", err)
	}

	f.Close()
}

func TestGetLockFilePid(t *testing.T) {
	f := createTestLock(t)
	defer os.Remove(f.Name())
	defer f.Close()

	pid, err := GetLockFilePid(f.Name())
	if err != nil {
		t.Fatal("Expected no error while getting PID, got:", err)
	}
	if pid != os.Getpid() {
		t.Errorf("Invalid PID: expected %v but got %v", os.Getpid(), pid)
	}
}
