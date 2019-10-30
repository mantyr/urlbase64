package main

import (
	"fmt"
	"net/url"
	"encoding/base64"
	"os"
	"io/ioutil"
)

func main() {
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	s, err := url.QueryUnescape(string(data))
	if err != nil {
		panic(err)
	}
	data, err = base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n\n", string(data))
}
