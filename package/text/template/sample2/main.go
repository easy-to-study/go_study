package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"flag"
	"io"
	"os"
	"text/template"
)

const Version = "0.1"

var (
	globFlag       bool
	jsonFilename   string
	outputFilename string
	versionFlag    bool
)

func init() {
	flag.BoolVar(&globFlag, "g", false, "enable globbing template files")
	flag.StringVar(&jsonFilename, "j", "-", "json filename (default stdin)")
	flag.StringVar(&outputFilename, "o", "-", "output filename (default stdout)")
	flag.BoolVar(&versionFlag, "v", false, "show version and exit")
}

func main() {
	flag.Parse()

	if versionFlag {
		println(Version)
		return
	}

	if flag.NArg() == 0 {
		showErrAndExit(errors.New("Please pass template filenames as arguments"))
	}

	var err error
	var tpl *template.Template
	if globFlag {
		tpl, err = template.ParseGlob(flag.Arg(0))
	} else {
		tpl, err = template.ParseFiles(flag.Args()...)
	}
	if err != nil {
		showErrAndExit(err)
	}

	var data interface{}
	if jsonFilename == "-" {
		data, err = buildDataFromReader(os.Stdin)
	} else {
		data, err = buildDataFromFile(jsonFilename)
	}
	if err != nil {
		showErrAndExit(err)
	}

	if outputFilename == "-" {
		err = executeTemplateWriter(tpl, data, os.Stdout)
	} else {
		err = executeTemplateFile(tpl, data, outputFilename)
	}
	if err != nil {
		showErrAndExit(err)
	}
}

func showErrAndExit(err error) {
	println(err.Error())
	os.Exit(1)
}

func buildDataFromFile(filename string) (interface{}, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return buildDataFromReader(f)
}

func buildDataFromReader(r io.Reader) (interface{}, error) {
	var v interface{}
	err := json.NewDecoder(bufio.NewReader(r)).Decode(&v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func executeTemplateFile(tpl *template.Template, data interface{}, filename string) error {
	w, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer w.Close()
	return executeTemplateWriter(tpl, data, w)
}

func executeTemplateWriter(tpl *template.Template, data interface{}, writer io.Writer) error {
	w := bufio.NewWriter(writer)
	err := tpl.Execute(w, data)
	if err != nil {
		return err
	}
	return w.Flush()
}
