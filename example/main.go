package main

import (
	"log"
	"os"

	"github.com/zerotone.ae/catlog"
)

func main() {
	catlog.Info("std log")
	catlog.SetOptions(catlog.WithLevel(catlog.DebugLevel))
	catlog.Debug("change std log to debug level")
	catlog.SetOptions(catlog.WithFormatter(&catlog.JsonFormatter{IgnoreBasicFields: false}))
	catlog.Debug("log in json format")
	catlog.Info("another log in json format")

	// output to file
	fd, err := os.OpenFile("test.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln("create file test.log failed")
	}
	defer fd.Close()

	l := catlog.New(catlog.WithLevel(catlog.InfoLevel), catlog.WithOutput(fd), catlog.WithFormatter(&catlog.JsonFormatter{IgnoreBasicFields: false}))
	l.Info("custom log with json formatter")
}
