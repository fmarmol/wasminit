package main

import (
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"text/template"

	"github.com/spf13/pflag"
)

type Data struct {
	JS      string
	HTML    string
	Package string
}

const templateSrc = `package {{.Package}}

var (
	FileJS = "{{.JS}}"
	FileHTML = "{{.HTML}}"
)
`

func ExecTemplate(t string, w io.Writer, data Data) error {
	tmpl := template.Must(template.New("toto").Parse(t))
	return tmpl.Execute(w, data)
}

func raise(err error) {
	if err != nil {
		panic(err)
	}
}

func readFile(path string) string {
	fd, err := os.Open(path)
	raise(err)
	defer fd.Close()
	raw, err := io.ReadAll(fd)
	raise(err)
	return base64.StdEncoding.EncodeToString(raw)
}

func main() {
	dir := pflag.StringP("dir", "d", ".", "dst directory")
	file := pflag.StringP("file", "f", "wasm.go", "dst filename")
	pkg := pflag.StringP("package", "p", "main", "dst package name")

	pflag.Parse()

	contentJS := readFile("./cmd/packer/wasm_exec.js")
	contentHTML := readFile("./cmd/packer/wasm_exec.html")

	data := Data{
		HTML:    contentHTML,
		JS:      contentJS,
		Package: *pkg,
	}

	path := filepath.Join(*dir, *file)
	dst, err := os.Create(path)
	raise(err)
	defer dst.Close()

	err = ExecTemplate(templateSrc, dst, data)
	raise(err)
	fmt.Printf("SUCCESS: ouput file created at %v\n", path)
}
