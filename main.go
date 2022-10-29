package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/common-nighthawk/go-figure"
)

// set default port number if env var $PORT isn't set
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// gets `font` header to change font type
// the font value must be supported via https://github.com/common-nighthawk/go-figure#supported-fonts
// if the `font` header is empty, the r.Header.Get will return `""`
func requestHandler(w http.ResponseWriter, r *http.Request) {

	// handle PUTs
	if r.Method == "PUT" {
		// attempt to read the payload of the PUT
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("server: could not read request body: %s\n", err)
		}

		str := strings.TrimSpace(string(reqBody)) // trim leading & trailing whitespace
		// check if reqBody-derived string is empty
		if len(str) == 0 {
			str = "empty request body"
			fmt.Printf("PUT request with empty body\n")
		} else {
			fmt.Printf("PUT request with body `%s`\n", str)
		}

		asciiStr := figure.NewFigure(str, r.Header.Get("font"), false) // `false` means we'll replace non-ASCII chars with `?`
		io.WriteString(w, asciiStr.String())
		return
	}
	// otherwise handle GETs
	if r.URL.Path != "/" {
		fmt.Printf("GET request at %s\n", r.URL.Path)
		//http.NotFound(w, r)
		str := strings.Replace(r.URL.Path, "/", " ", -1)               // replace slashes in path with spaces so `hi/there` results in `hi there`
		str = strings.TrimSpace(str)                                   // trim leading & trailing whitespace
		asciiStr := figure.NewFigure(str, r.Header.Get("font"), false) // `false` means we'll replace non-ASCII chars with `?`
		io.WriteString(w, asciiStr.String())
		return
	}
	fmt.Printf("GET request at /\n")
	str := "base path"
	asciiStr := figure.NewFigure(str, r.Header.Get("font"), true)
	io.WriteString(w, asciiStr.String())
	return
}

func main() {
	port := getEnv("PORT", "8080")

	//http.HandleFunc("/*")
	http.HandleFunc("/", requestHandler)

	http.ListenAndServe(":"+port, nil)
}
