package log

import (
	"github.com/pkg/errors"
)

var (
	StackSourceFileName     = "source"
	StackSourceLineName     = "line"
	StackSourceFunctionName = "func"
)

type state struct {
	b []byte
}

// Write implement fmt.Formatter interface.
func (s *state) Write(b []byte) (n int, err error) {
	s.b = b
	return len(b), nil
}

// Width implement fmt.Formatter interface.
func (s *state) Width() (wid int, ok bool) {
	return 0, false
}

// Precision implement fmt.Formatter interface.
func (s *state) Precision() (prec int, ok bool) {
	return 0, false
}

// Flag implement fmt.Formatter interface.
func (s *state) Flag(c int) bool {
	return true
}

func frameField(f errors.Frame, s *state, c rune) string {
	f.Format(s, c)
	return string(s.b)
}

// DebugMarshalStack MarshalStack implements pkg/errors stack trace marshaling.
//
// zerolog.ErrorStackMarshaler = MarshalStack
func DebugMarshalStack(err error) interface{} {
	type stackTracer interface {
		StackTrace() errors.StackTrace
	}
	var sterr stackTracer
	var ok bool
	for err != nil {
		sterr, ok = err.(stackTracer)
		if ok {
			break
		}

		u, ok := err.(interface {
			Unwrap() error
		})
		if !ok {
			return nil
		}

		err = u.Unwrap()
	}
	if sterr == nil {
		return nil
	}

	st := sterr.StackTrace()
	s := &state{}
	out := make([]map[string]string, 0, len(st))
	for _, frame := range st {
		out = append(out, map[string]string{
			StackSourceFileName:     frameField(frame, s, 's') + ":" + frameField(frame, s, 'd'),
			StackSourceFunctionName: frameField(frame, s, 'n'),
		})
	}
	return out
}
