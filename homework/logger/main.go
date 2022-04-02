package logger

type Logger struct {
	OutPutPath   string
	OutPutMethod string
}

func NewLogger(OutPutPath, OutPutMethod string) *Logger {
	return &Logger{
		OutPutPath,
		OutPutMethod,
	}
}

func (l *Logger) Debug() {

}

func (l *Logger) Info() {

}

func main() {

}
