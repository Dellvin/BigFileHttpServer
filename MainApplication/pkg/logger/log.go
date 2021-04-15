package logger

type Interface interface {
	Info(err error)
	InfoStr(err string)

	Trace(err error)
	TraceStr(err string)

	Debug(err error)
	DebugStr(err string)

	Warning(err error)
	WarningStr(err string)

	Panic(err error)
	PanicStr(err string)

	Error(err error)
	ErrorStr(err string)
}