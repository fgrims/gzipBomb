package main

import (
	"compress/gzip"
	"io/ioutil"
	"log"
	"os"
)

func handlerFunc() {
	println("Start")

	f, err := os.Open("./gzB.gz")
	if err != nil {
		log.Fatal(err)
	}

	gzipreader, err := gzip.NewReader(f)
	if err != nil {
		log.Fatal(err)
	}
	println("Start ReadAll")
	content, err := ioutil.ReadAll(gzipreader)
	if err != nil {
		log.Fatal(err)
	}
	print(content)
}

func main() {
	handlerFunc()
}
