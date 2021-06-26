package logger_test

import (
	"fmt"
	log "main/src/logger"
	"testing"
)

func TestLog(_ *testing.T) {
	log.Info("info log")
	log.Warning("warning log")
	log.Error("error log")
	defer func()  {
		if err := recover(); err != nil {
			fmt.Println("catch a panic")
		}
	}()
	log.Fatal("fatal log")
}
