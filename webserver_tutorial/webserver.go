package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
)

func indexHandler(writer http.ResponseWriter, reader *http.Request) {
	fmt.Fprintf(writer, "Whoa, go is neat!")
	fmt.Println(reader)
}

func aboutHandler(writer http.ResponseWriter, reader *http.Request) {
	fmt.Fprintf(writer, "now we're at the about page")
	fmt.Println(reader)
}

func verifyHandler(writer http.ResponseWriter, reader *http.Request) {
	fmt.Fprintf(writer, "let's verify!")
	fmt.Println(reader)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/verify", verifyHandler)
	go http.ListenAndServe(":8000", nil)
	reader := bufio.NewReader(os.Stdin)
	reader.ReadRune()
}
