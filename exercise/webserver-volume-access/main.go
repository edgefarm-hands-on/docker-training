package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

var inputfile string

func main() {
	i := flag.String("i", "", "input file")
	flag.Parse()
	inputfile = *i
	fmt.Printf("using '%s' as input file\n", inputfile)
	fmt.Println("webserver-volume-access starting...")
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":8080", nil)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	dat, err := os.ReadFile(inputfile)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintf(w, "%s", string(dat))
}
