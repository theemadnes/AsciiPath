package main

import (
	"fmt"
	"io"
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
func getHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		fmt.Printf("Req: %s\n", r.URL.Path)
		//http.NotFound(w, r)
		str := strings.Replace(r.URL.Path, "/", " ", -1)               // replace slashes in path with spaces so `hi/there` results in `hi there`
		str = strings.TrimSpace(str)                                   // trim leading & trailing whitespace
		asciiStr := figure.NewFigure(str, r.Header.Get("font"), false) // `false` means we'll replace non-ASCII chars with `?`
		io.WriteString(w, asciiStr.String())
		return
	}
	fmt.Printf("got / request\n")
	str := "base path"
	asciiStr := figure.NewFigure(str, r.Header.Get("font"), true)
	io.WriteString(w, asciiStr.String())
}

func main() {
	port := getEnv("PORT", "8080")

	//http.HandleFunc("/*")
	http.HandleFunc("/", getHandler)

	http.ListenAndServe(":"+port, nil)
}
