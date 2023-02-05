package main

import (
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1> Welcome! </h1>")

	f, err := os.Open("./gzB.gz")
	if err != nil {
		log.Fatal(err)
	}

	gzipreader, err := gzip.NewReader(f)
	if err != nil {
		log.Fatal(err)
	}
	content, err := ioutil.ReadAll(gzipreader)
	if err != nil {
		log.Fatal(err)
	}

	var body string
	body = string(content)

	fmt.Println(body)
	fmt.Fprint(w, "<h1> OK! </h1>")
}

func main() {
	http.HandleFunc("/", handlerFunc)
	fmt.Println("Starting the server on :3000..")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
