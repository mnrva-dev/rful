package rful

import (
	"os"
	"sync"
	"time"
)

// RotateWriter implements io.Writer
type RotateWriter struct {
	mu         sync.Mutex
	filename   string
	fp         *os.File
	size       int64
	timeformat string
}

// Creates a new RotateWriter. Timeformat is the format in which
// time will be written to the filename.
// Ensure that the specified time format does not contain any
// disallowed characters for your operating system
//
// Size is the max size of the file in bytes. The max size is not strict
// and may be surpassed by up to one log entry.
func NewRotateWriter(filename, time string, size int64) RotateWriter {
	return RotateWriter{
		filename:   filename,
		timeformat: time,
		size:       size,
	}
}

func (r *RotateWriter) Write(p []byte) (n int, err error) {

	r.mu.Lock()
	defer r.mu.Unlock()
	i, err := os.Stat(r.filename)
	if err != nil {
		return 0, err
	}
	if i.Size() >= r.size && r.size != 0 {
		r.mu.Unlock()
		err = r.move()
		if err != nil {
			return 0, err
		}
		r.mu.Lock()
	}
	return r.fp.Write(p)

}

func (r *RotateWriter) move() error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// close current logging file
	if r.fp != nil {
		err := r.fp.Close()
		r.fp = nil
		if err != nil {
			// handle error closing file
			return err
		}
	}

	// if file already exists, rename it
	_, err := os.Stat(r.filename)
	if err == nil {
		err = os.Rename(r.filename, r.filename+"."+time.Now().Format(r.timeformat))
		if err != nil {
			// handle error renaming file
			return err
		}
	}

	r.fp, err = os.Create(r.filename)
	if err != nil {
		// handle error starting new log file
		return err
	}
	return nil
}
