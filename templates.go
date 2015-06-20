package brahma

import "html/template"

// The template store
var Templates = make(map[string]*template.Template)

// Get the name of the applicable template
func getRelevantTemplate(name, module string) *template.Template {
	extensions := []string{"", ".tmpl", ".html"}
	t := Templates[module]

	if t == nil {
		// Have we even parsed the module templates?
		return nil
	}

	for _, suffix := range extensions {
		// Return if a relevant template is found
		if temp := t.Lookup(name + suffix); temp != nil {
			return temp
		}
	}

	// We should never reach this point.
	// If we get here, it means that the template specified was invalid
	return nil
}

// Parse all templates from a module
// Templates are parsed on demand, as and when a module is requested
func checkAndParsePackageTemplates(base string) {
	if _, exists := Templates[base]; exists {
		// If the templates have already been parsed, don't do anything
		return
	}

	// Parse the templates and store them under the module name
	Templates[base] = template.Must(template.ParseGlob(`packages/` + base + `/templates/*.tmpl`))
}
