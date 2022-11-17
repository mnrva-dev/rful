package rful

import (
	"io"
)

type Logger struct {
	prefix     string
	time       string
	sep        string
	loc        func() string
	fmt        string
	lvl        int
	altWriters int
	w          io.Writer
	aw         io.Writer
}

// Returns a new Logger with the default config. To change the
// config, use SetPreface() and/or SetFormat().
func New(w io.Writer) *Logger {
	r := &Logger{
		prefix:     "",
		time:       TimeDefaultFmt,
		loc:        func() string { return "" },
		sep:        "-> ",
		fmt:        "%s%s%s%s\t%s%s",
		lvl:        LogLevelAll,
		altWriters: 0,
		w:          w,
		aw:         nil,
	}

	return r
}

// Logger does not provide a default prefix or location. The default time
// format is a slightly modified RFC3339, "2006-01-02 15:04:05".
//
// Logger provides many time format constants, which are identical to go's
// time formats, with time zones removed and added whitespace for clarity.
//
// The seperator is what is printed between the prefix string (time, location, etc)
// and the actual log text. The default is "-> ".
func (r *Logger) SetPreface(prefix string, timeformat string, location func() string, separator string) *Logger {
	r.prefix = prefix
	r.time = timeformat
	r.loc = location
	r.sep = separator
	return r
}

// Set a string to prefix every log entry
func (r *Logger) SetPrefix(prefix string) *Logger {
	r.prefix = prefix
	return r
}

// Determines the format for how time is printed.
func (r *Logger) SetTimeFormat(format string) *Logger {
	r.time = format
	return r
}

// The location function takes no arguments and should return the location of the event
// being logged in the desired format.
func (r *Logger) SetLocation(loc func() string) *Logger {
	r.loc = loc
	return r
}

// Define the separator
func (r *Logger) SetSeparator(separator string) *Logger {
	r.sep = separator
	return r
}

// Set the format for each log, in printf-style syntax. The default
// format is "%s%s%s%s\t%s%s".
//
// Make sure to include all arguments. The order of arguments
// is prefix, time, location, separator, and type.
func (r *Logger) SetFormat(format string) *Logger {
	r.fmt = format
	return r
}

// Define the log level. All logs below the specified log level will not be printed.
func (r *Logger) SetLevel(lvl int) *Logger {
	r.lvl = lvl
	return r
}

// Use the AltWriterXxx flags to enable the alternate io.Writer for specific log levels
func (r *Logger) EnableAltWriter(flags int) *Logger {
	r.altWriters = flags
	return r
}

// Set an alternate io.Writer that can be used for particular log levels.
// Enable the writer using the EnableAltWriter method.
func (r *Logger) SetAltWriter(w io.Writer) *Logger {
	r.aw = w
	return r
}
