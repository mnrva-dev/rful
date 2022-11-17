package rful

// Enable the alternate writer for specific log levels using
// these flags
const (
	AltTraceWriter = 1 << iota
	AltDebugWriter
	AltInfoWriter
	AltWarnWriter
	AltErrorWriter
	AltPanicWriter
	AltFatalWriter
)
