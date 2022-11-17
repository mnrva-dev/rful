package rful

import (
	"os"
	"sync"
	"time"
)

// RotateWriter implements io.Writer. RotateWriter will write to one file until it reaches a
// specified max size. At which point, it will move to a new file for subsequent writes.
type RotateWriter struct {
	mu               sync.Mutex
	filename         string
	fp               *os.File
	size             int64
	generateFilename func(old string) string
}

// Creates a new RotateWriter.
//
// Filename is the name of the file that will be written to.
//
// Size is the max size of each file in bytes. Pass 0 to write to only
// one file. The max size is not strict and may be surpassed by up to one write.
func NewRotateWriter(filename string, size int64) *RotateWriter {
	r := &RotateWriter{
		filename: filename,
		size:     size,
		generateFilename: func(old string) string {
			return old + "." + time.Now().Format("2006-01-02T15-04-05")
		},
	}
	r.Move()
	return r
}

// Checks if the current file has surpassed the max size. If not, it writes to the
// file. If yes, it moves to a new file and writes to it.
func (r *RotateWriter) Write(p []byte) (n int, err error) {

	r.mu.Lock()
	defer r.mu.Unlock()
	i, err := os.Stat(r.filename)
	if err != nil {
		return 0, err
	}
	if i.Size() >= r.size && r.size != 0 {
		r.mu.Unlock()
		err = r.Move()
		if err != nil {
			return 0, err
		}
		r.mu.Lock()
	}
	return r.fp.Write(p)

}

// Moves to a new file
func (r *RotateWriter) Move() error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// close current logging file
	if r.fp != nil {
		err := r.fp.Close()
		r.fp = nil
		if err != nil {
			return err
		}
	}

	// if file already exists, rename it
	_, err := os.Stat(r.filename)
	if err == nil {
		newfile := r.generateFilename(r.filename)
		err = os.Rename(r.filename, newfile)
		if err != nil {
			return err
		}
	}

	r.fp, err = os.Create(r.filename)
	if err != nil {
		return err
	}
	return nil
}

// Set a new function to generate file names for old files that have been rotated out.
// The function should take the old file name as an argument and return a unique new name.
func (r *RotateWriter) SetFilenameGenerator(fn func(string) string) {
	r.generateFilename = fn
}
