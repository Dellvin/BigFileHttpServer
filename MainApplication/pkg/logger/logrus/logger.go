package logrus

import (
	"HttpBigFilesServer/MainApplication/pkg/logger"
	"github.com/sirupsen/logrus"
	//"time"

	"os"
)

type logging struct {
	l *logrus.Logger
}

func New() logger.Interface {
	l := setupLogrus()
	return logging{l: l}
}

func setupLogrus() *logrus.Logger {
	var log = logrus.New()
	log.Formatter = new(logrus.JSONFormatter)
	log.Formatter = new(logrus.TextFormatter)                     //default
	log.Formatter.(*logrus.TextFormatter).DisableColors = true    // remove colors
	log.Formatter.(*logrus.TextFormatter).DisableTimestamp = true // remove timestamp from test output
	log.Level = logrus.TraceLevel
	log.Out = os.Stdout
	return log
}

func (log logging) Info(err error) {
	log.l.WithFields(logrus.Fields{
		"error": err.Error(),
		//"date":  time.Now().Date(),
		//"time":  time.Now().Clock(),
	}).Info()
}
func (log logging) InfoStr(err string) {
	log.l.WithFields(logrus.Fields{
		"error": err,
		//"date":  time.Now().Date(),
		//"time":  time.Now().Clock(),
	}).Info()
}
func (log logging)Trace(err error){
	log.l.WithFields(logrus.Fields{
		"error": err.Error(),
		//"date":  time.Now().Date(),
		//"time":  time.Now().Clock(),
	}).Trace()
}
func (log logging)TraceStr(err string){
	log.l.WithFields(logrus.Fields{
		"error": err,
		//"date":  time.Now().Date(),
		//"time":  time.Now().Clock(),
	}).Trace()
}

func (log logging)Debug(err error){
	log.l.WithFields(logrus.Fields{
		"error": err.Error(),
		//"date":  time.Now().Date(),
		//"time":  time.Now().Clock(),
	}).Debug()
}
func (log logging)DebugStr(err string){
	log.l.WithFields(logrus.Fields{
		"error": err,
		//"date":  time.Now().Date(),
		//"time":  time.Now().Clock(),
	}).Debug()
}

func (log logging)Warning(err error){
	log.l.WithFields(logrus.Fields{
		"error": err.Error(),
		//"date":  time.Now().Date(),
		//"time":  time.Now().Clock(),
	}).Warning()
}
func (log logging)WarningStr(err string){
	log.l.WithFields(logrus.Fields{
		"error": err,
		//"date":  time.Now().Date(),
		//"time":  time.Now().Clock(),
	}).Warning()
}

func (log logging)Panic(err error){
	log.l.WithFields(logrus.Fields{
		"error": err.Error(),
		//"date":  time.Now().Date(),
		//"time":  time.Now().Clock(),
	}).Panic()
}
func (log logging)PanicStr(err string){
	log.l.WithFields(logrus.Fields{
		"error": err,
		//"date":  time.Now().Date(),
		//"time":  time.Now().Clock(),
	}).Panic()
}

func (log logging)Error(err error){
	log.l.WithFields(logrus.Fields{
		"error": err.Error(),
		//"date":  time.Now().Date(),
		//"time":  time.Now().Clock(),
	}).Error()
}
func (log logging)ErrorStr(err string){
	log.l.WithFields(logrus.Fields{
		"error": err,
		//"date":  time.Now().Date(),
		//"time":  time.Now().Clock(),
	}).Error()
}
