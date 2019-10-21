package main

import (
	"io/ioutil"
	"net/http"
	"os"
)

// http://localhost:9090/list/ff.txt
func main() {
	http.HandleFunc("/list/",
		func(writer http.ResponseWriter, request *http.Request) {
			path := request.URL.Path[len("/list/"):]
			file, err := os.Open(path)
			if err != nil {
				panic(err)
			}
			defer file.Close()

			all, err := ioutil.ReadAll(file)
			if err != nil {
				panic(err)
			}

			writer.Write(all)
		})

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		panic(err)
	}
}
