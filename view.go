package carbonate

import (
	"bytes"
)

// The View interface
type View interface {
	ControllerChild
	setPackageName(string)
}

// The struct to be inherited by other module specific views
// Contains the name of the package and the output buffer
type BaseView struct {
	base   string
	output bytes.Buffer
}

// Set the parent controller
func (v *BaseView) setParent(c *BaseController) {
	c.view = v
}

// Set the name of the module
func (v *BaseView) setPackageName(base string) {
	v.base = base
	checkAndParsePackageTemplates(base)
}

// Render a template and write it to the output buffer
func (v *BaseView) RenderTemplate(name string, data interface{}) {
	html := v.RenderTemplateAsString(name, data)
	v.output.WriteString(html)
}

// Render a template and return the results as string
func (v *BaseView) RenderTemplateAsString(name string, data interface{}) string {
	template := getRelevantTemplate(name, v.base)
	var b bytes.Buffer

	// Check if the template is valid
	if template != nil {
		// If it is, then render it
		template.Execute(&b, data)
	} else {
		// Otherwise, we have a problem
		b.WriteString("No template found")
	}

	return b.String()
}
