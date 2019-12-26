package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://www.google.com")
	if err != nil {
		fmt.Println("ERROR ", err)
		os.Exit(1)
	}

	fmt.Println(resp)
	fmt.Println("***** Reading HTML BODY by implementing Reader  ******")

	bs := make([]byte, 9999)
	resp.Body.Read(bs)
	fmt.Println(string(bs))

	lw := logWriter{}
	io.Copy(lw, resp.Body)

	fmt.Println("***** COPY HTML BODY by implementing Write  ******")
	io.Copy(os.Stdout, resp.Body)
}

// logWriter now implementing the Write Interface
func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("MANY byte", len(bs))
	return len(bs), nil
}
