package hermes

import (
	"os"
	"path/filepath"
)
import "os/user"

// HasReadAccess checks if a file can be opened for reading.
// Returns true if the file can be opened, false otherwise.
func HasReadAccess(filename string) bool {
	f, err := os.Open(filename)
	if err != nil {
		return false
	}
	_ = f.Close()
	return true
}

// MessageDBFilename returns the full path to the message database.
func MessageDBFilename() (string, error) {
	u, err := user.Current()
	if err != nil {
		return "", err
	}
	homeDir := u.HomeDir

	// Default path in MacOS 12
	return filepath.Join(homeDir, "Library", "Messages", "chat.db"), nil
}
