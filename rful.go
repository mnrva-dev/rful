package rful

import (
	"io"
)

type Rful struct {
	prefix string
	time   string
	sep    string
	loc    func() string
	w      io.Writer
}

// Returns a new Rful logger with the default config. To change the
// config, use SetFormat()
func New(w io.Writer) Rful {
	r := Rful{
		prefix: "",
		time:   "2006-01-02T15:04:05",
		loc:    func() string { return "" },
		sep:    "->",
		w:      w,
	}

	return r
}

// Rful does not provide a default prefix or location. The default time
// format is RFC3339, "2006-01-02T15:04:05"
//
// Rful provides many time format constants, which are identical to go's
// time formats, with time zones removed
//
// The seperator is what is printed between the prefix string (time, location, etc)
// and the actual log.
func (r *Rful) SetFormat(prefix string, timeformat string, location func() string, separator string) {
	r.prefix = prefix
	r.time = timeformat
	r.loc = location
	r.sep = separator
}
