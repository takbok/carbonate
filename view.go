package framework

import (
	"bytes"
)

type View interface {
	ControllerChild
	setPackageName(string)
}

type BaseView struct {
	base   string
	output bytes.Buffer
}

func (v *BaseView) SetParent(c *BaseController) {
	c.view = v
}

func (v *BaseView) setPackageName(base string) {
	v.base = base
	checkAndParsePackageTemplates(base)
}

func (v *BaseView) RenderTemplate(name string, data interface{}) {
	html := v.RenderTemplateAsString(name, data)
	v.output.WriteString(html)
}

func (v *BaseView) RenderTemplateAsString(name string, data interface{}) string {
	template := getRelevantTemplate(name, v.base)
	var b bytes.Buffer

	if template != nil {
		template.Execute(&b, data)
	} else {
		b.WriteString("No template found")
	}

	return b.String()
}
