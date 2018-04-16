package crash

import (
	"runtime"

	"github.com/ansel1/merry"
	"github.com/getsentry/raven-go"
	"github.com/percolate/shisa/context"
	"github.com/percolate/shisa/errorx"
	"github.com/percolate/shisa/httpx"
)

// SentryReporter reports errors to an external Sentry service
type SentryReporter struct {
	client *raven.Client
}

// NewSentryReporter instantiates a *raven.Client and *SentryReporter
// from a sentry DSN string
func NewSentryReporter(dsn string) (*SentryReporter, merry.Error) {
	client, err := raven.New(dsn)
	if err != nil {
		return nil, merry.Prepend(err, "crash: new sentry reporter")
	}
	return &SentryReporter{client: client}, nil
}

// Report checks if the provided merry.Error is a panic (as defined by
// errors.IsPanic), and if so, constructs a stacktrace and sends an an
// exception to sentry
func (s *SentryReporter) Report(ctx context.Context, r *httpx.Request, err merry.Error) {
	if !errorx.IsPanic(err) {
		return
	}

	exception := raven.NewException(err, merryToStacktrace(err))
	user := &raven.User{ID: ctx.Actor().ID()}

	sentryInterfaces := []raven.Interface{exception, user, raven.NewHttp(r.Request)}

	packet := raven.NewPacket(err.Error(), sentryInterfaces...)
	tags := map[string]string{"request-id": ctx.RequestID()}

	s.client.Capture(packet, tags)
}

// Close safely closes the *raven.Client
func (s *SentryReporter) Close() (exception merry.Error) {
	defer func() {
		arg := recover()
		if arg == nil {
			return
		}
		exception = errorx.CapturePanic(arg, "crash: panic in sentry reporter close")
	}()

	s.client.Close()

	return
}

func merryToStacktrace(err merry.Error) *raven.Stacktrace {
	var frames []*raven.StacktraceFrame
	for _, f := range merry.Stack(err) {
		pc := uintptr(f) - 1
		fn := runtime.FuncForPC(pc)
		var file string
		var line int
		if fn != nil {
			file, line = fn.FileLine(pc)
		} else {
			file = "unknown"
		}
		frame := raven.NewStacktraceFrame(pc, file, line, 0, nil)
		if frame != nil {
			frames = append([]*raven.StacktraceFrame{frame}, frames...)
		}
	}
	return &raven.Stacktrace{Frames: frames}
}
