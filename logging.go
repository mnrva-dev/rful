package rful

import (
	"fmt"
	"strings"
	"time"
)

// Write a new log at the Trace level
func (r *Logger) Trace(v ...any) (int, error) {
	if r.lvl > LogLevelTrace {
		return 0, nil
	}
	str := fmt.Sprintf(r.fmt, r.prefix, time.Now().Format(r.time), r.loc(), "TRACE ", r.sep, fmt.Sprint(v...))
	str = strings.TrimSpace(str)
	str += "\n"
	if r.altWriters&AltTraceWriter == AltTraceWriter {
		return r.aw.Write([]byte(str))
	} else {
		return r.w.Write([]byte(str))
	}
}

// Write a new log at the Debug level
func (r *Logger) Debug(v ...any) (int, error) {
	if r.lvl > LogLevelDebug {
		return 0, nil
	}
	str := fmt.Sprintf(r.fmt, r.prefix, time.Now().Format(r.time), r.loc(), "DEBUG ", r.sep, fmt.Sprint(v...))
	str = strings.TrimSpace(str)
	str += "\n"
	if r.altWriters&AltDebugWriter == AltDebugWriter {
		return r.aw.Write([]byte(str))
	} else {
		return r.w.Write([]byte(str))
	}
}

// Write a new log at the Info level
func (r *Logger) Info(v ...any) (int, error) {
	if r.lvl > LogLevelInfo {
		return 0, nil
	}
	str := fmt.Sprintf(r.fmt, r.prefix, time.Now().Format(r.time), r.loc(), "INFO  ", r.sep, fmt.Sprint(v...))
	str = strings.TrimSpace(str)
	str += "\n"
	if r.altWriters&AltInfoWriter == AltInfoWriter {
		return r.aw.Write([]byte(str))
	} else {
		return r.w.Write([]byte(str))
	}
}

// Write a new log at the Warn level
func (r *Logger) Warn(v ...any) (int, error) {
	if r.lvl > LogLevelWarn {
		return 0, nil
	}
	str := fmt.Sprintf(r.fmt, r.prefix, time.Now().Format(r.time), r.loc(), "WARN  ", r.sep, fmt.Sprint(v...))
	str = strings.TrimSpace(str)
	str += "\n"
	if r.altWriters&AltWarnWriter == AltWarnWriter {
		return r.aw.Write([]byte(str))
	} else {
		return r.w.Write([]byte(str))
	}
}

// Write a new log at the Error level
func (r *Logger) Error(v ...any) (int, error) {
	if r.lvl > LogLevelError {
		return 0, nil
	}
	str := fmt.Sprintf(r.fmt, r.prefix, time.Now().Format(r.time), r.loc(), "ERROR ", r.sep, fmt.Sprint(v...))
	str = strings.TrimSpace(str)
	str += "\n"
	if r.altWriters&AltErrorWriter == AltErrorWriter {
		return r.aw.Write([]byte(str))
	} else {
		return r.w.Write([]byte(str))
	}
}

// Write a new log at the Panic level
func (r *Logger) Panic(v ...any) (int, error) {
	if r.lvl > LogLevelPanic {
		return 0, nil
	}
	str := fmt.Sprintf(r.fmt, r.prefix, time.Now().Format(r.time), r.loc(), "PANIC ", r.sep, fmt.Sprint(v...))
	str = strings.TrimSpace(str)
	str += "\n"
	if r.altWriters&AltPanicWriter == AltPanicWriter {
		return r.aw.Write([]byte(str))
	} else {
		return r.w.Write([]byte(str))
	}
}

// Write a new log at the Fatal level
func (r *Logger) Fatal(v ...any) (int, error) {
	str := fmt.Sprintf(r.fmt, r.prefix, time.Now().Format(r.time), r.loc(), "FATAL ", r.sep, fmt.Sprint(v...))
	str = strings.TrimSpace(str)
	str += "\n"
	if r.altWriters&AltFatalWriter == AltFatalWriter {
		return r.aw.Write([]byte(str))
	} else {
		return r.w.Write([]byte(str))
	}
}

func (r *Logger) Tracef(format string, v ...any) (int, error) {
	if r.lvl > LogLevelTrace {
		return 0, nil
	}
	p := fmt.Sprintf(format, v...)
	str := fmt.Sprintf(r.fmt, r.prefix, time.Now().Format(r.time), r.loc(), "TRACE ", r.sep, p)
	str = strings.TrimSpace(str)
	str += "\n"
	if r.altWriters&AltTraceWriter == AltTraceWriter {
		return r.aw.Write([]byte(str))
	} else {
		return r.w.Write([]byte(str))
	}
}

func (r *Logger) Debugf(format string, v ...any) (int, error) {
	if r.lvl > LogLevelDebug {
		return 0, nil
	}
	p := fmt.Sprintf(format, v...)
	str := fmt.Sprintf(r.fmt, r.prefix, time.Now().Format(r.time), r.loc(), "DEBUG ", r.sep, p)
	str = strings.TrimSpace(str)
	str += "\n"
	if r.altWriters&AltDebugWriter == AltDebugWriter {
		return r.aw.Write([]byte(str))
	} else {
		return r.w.Write([]byte(str))
	}
}

func (r *Logger) Infof(format string, v ...any) (int, error) {
	if r.lvl > LogLevelInfo {
		return 0, nil
	}
	p := fmt.Sprintf(format, v...)
	str := fmt.Sprintf(r.fmt, r.prefix, time.Now().Format(r.time), r.loc(), "INFO  ", r.sep, p)
	str = strings.TrimSpace(str)
	str += "\n"
	if r.altWriters&AltInfoWriter == AltInfoWriter {
		return r.aw.Write([]byte(str))
	} else {
		return r.w.Write([]byte(str))
	}
}

func (r *Logger) Warnf(format string, v ...any) (int, error) {
	if r.lvl > LogLevelWarn {
		return 0, nil
	}
	p := fmt.Sprintf(format, v...)
	str := fmt.Sprintf(r.fmt, r.prefix, time.Now().Format(r.time), r.loc(), "WARN  ", r.sep, p)
	str = strings.TrimSpace(str)
	str += "\n"
	if r.altWriters&AltWarnWriter == AltWarnWriter {
		return r.aw.Write([]byte(str))
	} else {
		return r.w.Write([]byte(str))
	}
}

func (r *Logger) Errorf(format string, v ...any) (int, error) {
	if r.lvl > LogLevelError {
		return 0, nil
	}
	p := fmt.Sprintf(format, v...)
	str := fmt.Sprintf(r.fmt, r.prefix, time.Now().Format(r.time), r.loc(), "ERROR ", r.sep, p)
	str = strings.TrimSpace(str)
	str += "\n"
	if r.altWriters&AltErrorWriter == AltErrorWriter {
		return r.aw.Write([]byte(str))
	} else {
		return r.w.Write([]byte(str))
	}
}

func (r *Logger) Panicf(format string, v ...any) (int, error) {
	if r.lvl > LogLevelPanic {
		return 0, nil
	}
	p := fmt.Sprintf(format, v...)
	str := fmt.Sprintf(r.fmt, r.prefix, time.Now().Format(r.time), r.loc(), "PANIC ", r.sep, p)
	str = strings.TrimSpace(str)
	str += "\n"
	if r.altWriters&AltPanicWriter == AltPanicWriter {
		return r.aw.Write([]byte(str))
	} else {
		return r.w.Write([]byte(str))
	}
}

func (r *Logger) Fatalf(format string, v ...any) (int, error) {
	p := fmt.Sprintf(format, v...)
	str := fmt.Sprintf(r.fmt, r.prefix, time.Now().Format(r.time), r.loc(), "FATAL ", r.sep, p)
	str = strings.TrimSpace(str)
	str += "\n"
	if r.altWriters&AltFatalWriter == AltFatalWriter {
		return r.aw.Write([]byte(str))
	} else {
		return r.w.Write([]byte(str))
	}
}
