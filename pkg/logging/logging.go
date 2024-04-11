package logging

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	DebugLogger   log.Logger
	InfoLogger    log.Logger
	WarningLogger log.Logger
	ErrorLogger   log.Logger
	FatalLogger   log.Logger
}

func Init() Logger {
	logger := Logger{
		DebugLogger:   *log.New(os.Stderr, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile),
		InfoLogger:    *log.New(os.Stderr, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
		WarningLogger: *log.New(os.Stderr, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile),
		ErrorLogger:   *log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
		FatalLogger:   *log.New(os.Stderr, "FATAL: ", log.Ldate|log.Ltime|log.Lshortfile),
	}

	err := os.Mkdir("logs", 0o666)
	if err != nil {
		panic(err)
	}
	// allFile, err = os.OpenFile("logs/all.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o666)
	if err != nil {
		panic(err)
	}

	logger.DebugLogger.SetOutput(io.Discard)
	logger.InfoLogger.SetOutput(io.Discard)
	logger.WarningLogger.SetOutput(io.Discard)
	logger.ErrorLogger.SetOutput(io.Discard)
	logger.FatalLogger.SetOutput(io.Discard)
	return logger
}
