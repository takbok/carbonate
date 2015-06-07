package framework

import "html/template"

var Templates = make(map[string]*template.Template)

func getRelevantTemplate(name, module string) *template.Template {
	extensions := []string{"", ".tmpl", ".html"}
	t := Templates[module]

	if t == nil {
		return nil
	}

	for _, suffix := range extensions {

		if temp := t.Lookup(name + suffix); temp != nil {
			return temp
		}
	}

	return nil
}

func checkAndParsePackageTemplates(base string) {
	if _, exists := Templates[base]; exists {
		return
	}

	Templates[base] = template.Must(template.ParseGlob(`packages/` + base + `/templates/*.tmpl`))
}
