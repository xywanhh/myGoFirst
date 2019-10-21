package main

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

const prefix  = "/list/"

func HandlerList(writer http.ResponseWriter, request *http.Request) error {
	if strings.Index(request.URL.Path, prefix) != 0 {
		return errors.New("path must satrt with " + prefix)
	}
	path := request.URL.Path[len(prefix):]
	file, err := os.Open(path)
	if err != nil {
		http.Error(writer,
			err.Error(),
			http.StatusInternalServerError)
		return err
	}
	defer file.Close()

	all, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	writer.Write(all)
	return nil
}

type appHandler func(writer http.ResponseWriter, request *http.Request) error

func errWrapper(handler1 appHandler) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			 if r := recover(); r != nil {
				 log.Printf("panic : %v", r)
				 http.Error(writer, http.StatusText(http.StatusInternalServerError),
					 http.StatusInternalServerError)
			 }
		}()
		err := handler1(writer, request)
		if err != nil {
			log.Printf("Error handling request %s", err.Error())

			if userError, ok := err.(userError); ok {
				http.Error(writer, userError.Message(), http.StatusBadRequest)
				return
			}

			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer,
				http.StatusText(code), code)
		}
	}
}

type userError interface {
	error
	Message() string
}

func main() {
	http.HandleFunc("/list/", errWrapper(HandlerList))

	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		panic(err)
	}
}
