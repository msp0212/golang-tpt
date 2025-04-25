package main

import (
	"fmt"
	"os"
	"text/template"
)

var pln = fmt.Println

func getTemplate(name, text string) *template.Template {
	// get a new template by a given name
	tmpl := template.New(name)
	// parse the text for '{{}}' pattern
	return template.Must(tmpl.Parse(text))
}

func textTemplateStrDemo() {
	t := getTemplate("tmpl1", "Hello, {{.}}\n")
	t.Execute(os.Stdout, "world")
}

func textTemplateStrSliceDemo() {
	t := getTemplate("tmpl1", "Hello, {{.}}\n")
	t.Execute(os.Stdout, []string{"foo", "bar", "baz"})
}

func textTemplateStructDemo() {
	type employee struct {
		Name string // exported field (start with Capital letter)
		age  uint   //age wont be printed since it is not exported
	}
	t := getTemplate("tmpl1", "Hello, {{.Name}} {{.age}}\n")
	t.Execute(os.Stdout, employee{"emp1", 25})
	pln()
}

func textTemplateMapDemo() {
	t := getTemplate("tmpl1", "Hello, {{.Key1}} {{.key2}}\n")
	t.Execute(os.Stdout, map[string]string{"Key1": "Val1", "key2": "val2"})
}

func textTemplateIfElseDemo() {
	/*
		"if" check considers a value as false if the value
		is default value of the given type.
		"-" trims the whitespaces
	*/
	t := getTemplate("tmpl1", "Hello, {{if . -}}      world {{else -}}				empty-world {{end}}\n")
	t.Execute(os.Stdout, "non-empty-string")
	t.Execute(os.Stdout, "")
	t.Execute(os.Stdout, 1)
	t.Execute(os.Stdout, 0)
	t.Execute(os.Stdout, 1.1)
	t.Execute(os.Stdout, 0.0)
}

func textTemplateRangeDemo() {
	t := getTemplate("tmpl1", "List of items:\n{{range .}}<item>{{.}}</item>\n{{end}}End\n")
	t.Execute(os.Stdout, []string{"foo", "bar", "baz", ""})
	t.Execute(os.Stdout, []int{-1, 1, 2, 101})
	t.Execute(os.Stdout, map[string]string{"Key1": "Val1", "key2": "val2"})
}

func textTemplateDemo() {
	pln("textTemplateDemo...")
	textTemplateStrDemo()
	textTemplateStrSliceDemo()
	textTemplateStructDemo()
	textTemplateMapDemo()
	textTemplateIfElseDemo()
	textTemplateRangeDemo()
}

func main() {
	textTemplateDemo()
}
