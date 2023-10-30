package main

import (
	"log"
	"os"
)

func GlobalLogger() {
	logFile, _ := os.OpenFile("log/log.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	log.SetOutput(logFile)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetPrefix("Pipe_rest : ")
}
