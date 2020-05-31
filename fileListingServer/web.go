package main

import (
	"io/ioutil"
	"net/http"
	"os"
)

func HandleFileList(writer http.ResponseWriter, request *http.Request) {
	path := request.URL.Path[len("/list/"):]
	file, err := os.Open(path)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	all, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	writer.Write(all)
}

func main() {
	http.HandleFunc("/list/", HandleFileList)
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
