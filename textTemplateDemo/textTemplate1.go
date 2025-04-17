package main

import "html/template"

func getTemplate(name, text string) *template.Template {
	// get a new template by a given name
	tmpl := template.New(name)
	// parse the text for '{{}}' pattern
	return template.Must(tmpl.Parse(text))
}

func main() {

}
