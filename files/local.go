package files

import "io"

// Storage defines the behavior for file operations
// Implementation may be of time local disk, or cloud storage, etc
type Storage interface {
	Save(path string, file io.Reader) error
}
