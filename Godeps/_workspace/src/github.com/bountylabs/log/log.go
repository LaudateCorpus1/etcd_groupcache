package log

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

type StructuredLog struct {
	Path       string
	Duration   time.Duration
	RemoteAddr string
	Referer    string
	Status     int
}

type Logger struct {
	info           *log.Logger
	warning        *log.Logger
	error          *log.Logger
	slack          *log.Logger
	request        *log.Logger
	requestEncoder *json.Encoder
	callDepth      int
}

func (this *Logger) ErrLogger() *log.Logger {
	return this.error
}

func New(prefix string, depth int) *Logger {

	requestLogger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	l := &Logger{
		info:           log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile),
		warning:        log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile),
		error:          log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Lshortfile),
		slack:          log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile),
		request:        requestLogger,
		requestEncoder: json.NewEncoder(os.Stdout),
		callDepth:      depth,
	}

	l.SetPrefix(prefix)
	return l
}

var defaultLogger = New("", 3)

type BLogger interface {
	Errorf(format string, v ...interface{})
	Errorln(v ...interface{})
	Warningf(format string, v ...interface{})
	Warningln(v ...interface{})
	Infof(format string, v ...interface{})
	Infoln(v ...interface{})
	Requestln(v ...interface{})
	Panicln(v ...interface{})
	Panicf(format string, v ...interface{})
	Slackf(format string, v ...interface{})
	SlackLn(v ...interface{})
}

func (this *Logger) SetPrefix(prefix string) {

	if prefix != "" {
		prefix = prefix + " "
	}

	this.info.SetPrefix(prefix + "I: ")
	this.warning.SetPrefix(prefix + "W: ")
	this.error.SetPrefix(prefix + "E: ")
	this.slack.SetPrefix(prefix + "SLACK: ")
	this.request.SetPrefix(prefix + "R: ")
}

func (this *Logger) Errorf(format string, v ...interface{}) {
	this.error.Output(this.callDepth, f(format, v...))
}

func (this *Logger) Errorln(v ...interface{}) {
	this.error.Output(this.callDepth, ln(v...))
}

func (this *Logger) Warningf(format string, v ...interface{}) {
	this.warning.Output(this.callDepth, f(format, v...))
}

func (this *Logger) Warningln(v ...interface{}) {
	this.warning.Output(this.callDepth, ln(v...))
}

func (this *Logger) Infof(format string, v ...interface{}) {
	this.info.Output(this.callDepth, f(format, v...))
}

func (this *Logger) Infoln(v ...interface{}) {
	this.info.Output(this.callDepth, ln(v...))
}

func (this *Logger) Requestln(v ...interface{}) {
	this.request.Println(v...)
}

func (this *Logger) RequestEncoder() *json.Encoder {
	return this.requestEncoder
}

func (this *Logger) Panicln(v ...interface{}) {
	string := ln(v...)
	this.error.Output(this.callDepth, string)
	panic(string)
}

func (this *Logger) Panicf(format string, v ...interface{}) {
	string := f(format, v...)
	this.error.Output(this.callDepth, string)
	panic(string)
}

func (this *Logger) Slackf(format string, v ...interface{}) {
	this.slack.Output(this.callDepth, f(format, v...))
}

func (this *Logger) SlackLn(v ...interface{}) {
	this.slack.Output(this.callDepth, ln(v...))
}

//Globals
func DefaultLogger() *Logger {
	return defaultLogger
}

func Errorf(format string, v ...interface{}) {
	defaultLogger.Errorf(format, v...)
}

func Errorln(v ...interface{}) {
	defaultLogger.Errorln(v...)
}

func Warningf(format string, v ...interface{}) {
	defaultLogger.Warningf(format, v...)
}

func Warningln(v ...interface{}) {
	defaultLogger.Warningln(v...)
}

func Infof(format string, v ...interface{}) {
	defaultLogger.Infof(format, v...)
}

func Infoln(v ...interface{}) {
	defaultLogger.Infoln(v...)
}

func Requestln(v ...interface{}) {
	defaultLogger.Requestln(v...)
}

func RequestObject(obj StructuredLog) {
	defaultLogger.requestEncoder.Encode(obj)
}

func Panicln(v ...interface{}) {
	defaultLogger.Panicln(v...)
}

func Panicf(format string, v ...interface{}) {
	defaultLogger.Panicf(format, v...)
}

func Slackf(format string, v ...interface{}) {
	defaultLogger.Slackf(format, v...)
}

func SlackLn(v ...interface{}) {
	defaultLogger.SlackLn(v...)
}

func SetPrefix(p string) {
	defaultLogger.SetPrefix(p)
}

func f(format string, v ...interface{}) string {
	return ln(fmt.Sprintf(format, v...))
}

func ln(a ...interface{}) string {
	return fmt.Sprintln(a...)
}
