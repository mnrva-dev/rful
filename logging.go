package rful

import (
	"fmt"
	"time"
)

const basicfmt = "%s %s %s\t%s\t\t%s"

func (r *Rful) Trace(p string) {
	r.w.Write([]byte(fmt.Sprintf(basicfmt, r.prefix, time.Now().Format(r.time), r.loc(), "TRACE", p)))
}
