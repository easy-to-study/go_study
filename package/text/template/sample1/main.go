package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

func main() {
	// First we create a FuncMap with which to register the function.
	funcMap := template.FuncMap{
		// The name "title" is what the function will be called in the template text.
		//「title」という名前は、テンプレートテキストで関数が呼び出されるものです。
		"title": strings.Title,
	}

	// A simple template definition to test our function.
	// We print the input text several ways:
	// - the original
	// - title-cased
	// - title-cased and then printed with %q
	// - printed with %q and then title-cased.
	//関数をテストするためのシンプルなテンプレート定義。
	//入力テキストをいくつかの方法で印刷します。
	// - オリジナル
	//-タイトルを大文字にして
	//-タイトルを大文字にして、％qで出力します
	//-％qで印刷してから、タイトルを大文字にします。
	const templateText = `
Input: {{printf "%q" .}}
Output 0: {{title .}}
Output 1: {{title . | printf "%q"}}
Output 2: {{printf "%q" . | title}}
Output 3: {{body .}}
`

	// Create a template, add the function map, and parse the text.
	tmpl, err := template.New("titleTest").Funcs(funcMap).Parse(templateText)
	if err != nil {
		log.Fatalf("parsing: %s", err)
	}

	// Run the template to verify the output.
	err = tmpl.Execute(os.Stdout, "the go programming language")
	if err != nil {
		log.Fatalf("execution: %s", err)
	}

}

// 実行結果
//  t-sataga@MBA  ~/go/src/github.com/easy-to-study/go_study/package/text/template/sample1   master 
// -> % go run main.go                                                                                                                                             16:29:15 - 16:29:18

// Input: "the go programming language"
// Output 0: The Go Programming Language
// Output 1: "The Go Programming Language"
// Output 2: "The Go Programming Language"
//  t-sataga@MBA  ~/go/src/github.com/easy-to-study/go_study/package/text/template/sample1   master 
// -> %
