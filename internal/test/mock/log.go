package mock

type Logger struct {
	info   [][]any
	warns  [][]any
	errors [][]any
	prints [][]any
}

func (l *Logger) Info(v ...any) {
	l.info = append(l.info, v)
}
func (l *Logger) Debug(v ...any) {}

func (l *Logger) Warning(v ...any) {
	l.warns = append(l.warns, v)
}

func (l *Logger) Error(v ...any) {
	l.errors = append(l.errors, v)
}

func (l *Logger) Print(v ...any) {
	l.prints = append(l.prints, v)
}
