package pluginsdk

type Logger interface {
	LogEntry(string, ...interface{})
	LogSpan(string, ...interface{})
}
