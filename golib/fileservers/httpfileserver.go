package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

var (
	ENDPOINT = "/dbconfal"
	DEFAULTPORT = ":55555"

	FILEROOT = "/tmp/dbconfal/"
)

func sharefile(w http.ResponseWriter, req *http.Request) {
	filename := req.FormValue("file")

	if strings.Contains(filename, "..") {
		http.Error(w, "requested file is not available", http.StatusBadRequest)
	}

	if len(filename) != 0 {
		_, err := os.Stat(FILEROOT + filename)
		if err != nil {
			http.Error(w, "requested file is not available", http.StatusBadRequest)
		}

		bytes, err := ioutil.ReadFile(FILEROOT + filename)

		w.Write(bytes)
	}
}

func main() {

	//helloHandler := func(w http.ResponseWriter, req *http.Request) {
	//	io.WriteString(w, "Hello, world!\n")}
	http.HandleFunc(ENDPOINT, sharefile)
	log.Fatal(http.ListenAndServe(DEFAULTPORT, nil))
}