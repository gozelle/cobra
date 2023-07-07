package cobra

import (
	"fmt"
	"github.com/gozelle/cobra/progress"
	"github.com/gozelle/color"
	"os"
	"regexp"
	"strings"
	"time"
)

var logger *Logger

func init() {
	logger = NewLogger()
}

func WithModule(module string) *Logger {
	return logger.WithModule(module)
}

func Success(msgs ...string) {
	logger.Success(msgs...)
}

func Info(msgs ...string) {
	logger.Info(msgs...)
}

func Error(msgs ...string) {
	logger.Error(msgs...)
}

func Warn(msgs ...string) {
	logger.Warn(msgs...)
}

func Debug(msgs ...string) {
	logger.Debug(msgs...)
}

func Fatal(msgs ...string) {
	logger.Fatal(msgs...)
}

func Progress(fields ...*progress.Field) {
	logger.Progress(fields...)
}

func NewLine() {
	logger.NewLine()
}

func NewLogger() *Logger {
	return &Logger{}
}

type Logger struct {
	module string
}

func (l *Logger) WithModule(module string) *Logger {
	l.module = module
	return l
}

func (l *Logger) Success(msgs ...string) {
	l.print(color.GreenString, msgs...)
}

func (l *Logger) Info(msgs ...string) {
	l.print(color.CyanString, msgs...)
}

func (l *Logger) Error(msgs ...string) {
	l.print(color.RedString, msgs...)
}

func (l *Logger) Warn(msgs ...string) {
	l.print(color.YellowString, msgs...)
}

func (l *Logger) Debug(msgs ...string) {
	l.print(color.BlueString, msgs...)
}

func (l *Logger) Fatal(msgs ...string) {
	l.print(color.HiRedString, msgs...)
	os.Exit(1)
}

func (l *Logger) Progress(fields ...*progress.Field) {
	var f []string
	var vals []any
	for _, v := range fields {
		f = append(f, v.Format())
		vals = append(vals, v.Value)
	}
	if len(f) > 0 {
		s := fmt.Sprintf("\r%s", strings.Join(f, " "))
		fmt.Printf(s, vals...)
	}
}

func (l *Logger) NewLine() {
	fmt.Printf("\n")
}

func (l *Logger) now() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

var emojiRegex = regexp.MustCompile(`[\x{1F600}-\x{1F64F}\x{1F300}-\x{1F5FF}\x{1F680}-\x{1F6FF}\x{1F1E0}-\x{1F1FF}\x{2600}-\x{26FF}\x{2700}-\x{27BF}\x{1F900}-\x{1F9FF}\x{1F018}-\x{1F270}\x{1F300}-\x{1F5FF}\x{1F1E6}-\x{1F1FF}\x{1F600}-\x{1F64F}\x{1F680}-\x{1F6C5}\x{1F30D}-\x{1F567}]`)

func (l *Logger) isEmoji(str string) bool {
	return emojiRegex.MatchString(str)
}

func (l *Logger) print(c func(format string, a ...interface{}) string, msgs ...string) {
	
	if n := len(msgs); n > 0 {
		if !l.isEmoji(msgs[0]) {
			msgs[0] = c(msgs[0])
		} else if n > 1 {
			msgs[1] = c(msgs[1])
		}
	}
	
	for i := 0; i < len(msgs); i++ {
		if msgs[i] == "" {
			msgs = append(msgs[0:i], msgs[i+1:]...)
		}
	}
	
	var module string
	if l.module != "" {
		module = fmt.Sprintf(" %s ", color.New(color.FgBlack, color.BgWhite).Sprintf("[%s]", l.module))
	}
	
	fmt.Printf("%s %s%s\n",
		color.WhiteString(l.now()),
		module,
		strings.Join([]string{strings.Join(msgs, " ")}, " "),
	)
}
