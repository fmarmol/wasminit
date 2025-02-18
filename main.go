package main

import (
	"bytes"
	"encoding/base64"
	"io"
	"os"
)

func genjs() {
	src, err := base64.StdEncoding.DecodeString(FileJS)
	if err != nil {
		panic(err)
	}
	out, err := os.Create("wasm_exec.js")
	if err != nil {
		panic(err)
	}
	defer out.Close()
	_, err = io.Copy(out, bytes.NewBuffer(src))
	if err != nil {
		panic(err)
	}
}

func genHtml() {
	src, err := base64.StdEncoding.DecodeString(FileHTML)
	if err != nil {
		panic(err)
	}
	out, err := os.Create("index.html")
	if err != nil {
		panic(err)
	}
	defer out.Close()
	_, err = io.Copy(out, bytes.NewBuffer(src))
	if err != nil {
		panic(err)
	}
}

func main() {
	genjs()
	genHtml()
}
