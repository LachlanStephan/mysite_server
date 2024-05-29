package main

import (
	"html/template"
	"log"
	"os"
)

type application struct {
	errorLog  *log.Logger
	infoLog   *log.Logger
	fileCache map[string]*template.Template
}

var (
	infoLog  = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Llongfile)
)

func main() {
	fileCache, err := newFileCache()
	if err != nil {
		errorLog.Fatal(err)
	}

	app := &application{
		errorLog:  errorLog,
		infoLog:   infoLog,
		fileCache: fileCache,
	}

	runServer(getPort(), app)
}

func getPort() string {
	p := os.Getenv(port)
	if p == "" {
		return portDefault
	}

	return p
}
