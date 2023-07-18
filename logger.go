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

var log Logger

func init() {
	log = NewLogger()
}

func WithModule(module string) Logger {
	return log.WithModule(module)
}

func Success(msgs ...string) {
	log.Success(msgs...)
}

func Info(msgs ...string) {
	log.Info(msgs...)
}

func Error(msgs ...string) {
	log.Error(msgs...)
}

func Warn(msgs ...string) {
	log.Warn(msgs...)
}

func Debug(msgs ...string) {
	log.Debug(msgs...)
}

func Fatal(msgs ...string) {
	log.Fatal(msgs...)
}

func Progress(fields ...*progress.Field) {
	log.Progress(fields...)
}

func NewLine() {
	log.NewLine()
}

type Logger interface {
	WithModule(module string) *logger
	Success(msgs ...string)
	Info(msgs ...string)
	Error(msgs ...string)
	Warn(msgs ...string)
	Debug(msgs ...string)
	Fatal(msgs ...string)
	Progress(fields ...*progress.Field)
	NewLine()
}

func NewLogger() Logger {
	return &logger{}
}

type logger struct {
	module string
}

func (l *logger) WithModule(module string) *logger {
	l.module = module
	return l
}

func (l *logger) Success(msgs ...string) {
	l.print(color.GreenString, msgs...)
}

func (l *logger) Info(msgs ...string) {
	l.print(color.CyanString, msgs...)
}

func (l *logger) Error(msgs ...string) {
	l.print(color.RedString, msgs...)
}

func (l *logger) Warn(msgs ...string) {
	l.print(color.YellowString, msgs...)
}

func (l *logger) Debug(msgs ...string) {
	l.print(color.BlueString, msgs...)
}

func (l *logger) Fatal(msgs ...string) {
	l.print(color.HiRedString, msgs...)
	os.Exit(1)
}

func (l *logger) Progress(fields ...*progress.Field) {
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

func (l *logger) NewLine() {
	fmt.Printf("\n")
}

func (l *logger) now() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

var emojiRegex = regexp.MustCompile(`[\x{1F600}-\x{1F64F}\x{1F300}-\x{1F5FF}\x{1F680}-\x{1F6FF}\x{1F1E0}-\x{1F1FF}\x{2600}-\x{26FF}\x{2700}-\x{27BF}\x{1F900}-\x{1F9FF}\x{1F018}-\x{1F270}\x{1F300}-\x{1F5FF}\x{1F1E6}-\x{1F1FF}\x{1F600}-\x{1F64F}\x{1F680}-\x{1F6C5}\x{1F30D}-\x{1F567}]`)

func (l *logger) isEmoji(str string) bool {
	return emojiRegex.MatchString(str)
}

func (l *logger) print(c func(format string, a ...interface{}) string, msgs ...string) {
	
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
		module = fmt.Sprintf(" %s ", color.New(color.FgMagenta).Sprintf("[%s]", l.module))
	}
	
	fmt.Printf("%s %s%s\n",
		color.WhiteString(l.now()),
		module,
		strings.Join([]string{strings.Join(msgs, " ")}, " "),
	)
}
