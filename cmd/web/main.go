package main

import (
	"log"
	"os"

	"github.com/LachlanStephan/mysite_server/cmd/files"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	files    *files.Files
}

var (
	infoLog  = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Llongfile)
	port     = ":8080"
)

func main() {
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		files:    &files.Files{},
	}

	err := runServer(port, app)
	if err != nil {
		log.Fatal(err)
	}
}
