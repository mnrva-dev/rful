package rful_test

import (
	"os"
	"testing"

	"github.com/mnrva-dev/rful"
)

func TestXxx(t *testing.T) {
	l := rful.New(os.Stdout)
	l.SetFormat("", rful.TimeRFC3339, rful.LocShortFileFunc, "->")
	l.Trace("This is a trace log")
}
