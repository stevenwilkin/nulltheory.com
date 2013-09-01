package main

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"io/ioutil"
)


func handler(w http.ResponseWriter, r *http.Request) {
	var responseCode int

	// check for requested path under public/
	pwd, _ := os.Getwd()
	filename := path.Join(pwd, "public", r.URL.Path)
	stat, err := os.Stat(filename); 

	// append index.html if path is a directory
	if !os.IsNotExist(err) && stat.Mode().IsDir() {
		filename = path.Join(filename, "index.html")
		stat, err = os.Stat(filename); 
	}

	if os.IsNotExist(err) {
		responseCode = http.StatusNotFound
		w.WriteHeader(responseCode)
	} else {
		responseCode = http.StatusOK
		contents, _ := ioutil.ReadFile(filename)
		w.Write(contents)
	}

	fmt.Printf("%s %s %d\n", r.Method, r.URL.Path, responseCode)
}

func main() {
	fmt.Printf("> Starting on http://0.0.0.0:8080\n")
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting!")
	}
}
