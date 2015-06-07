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
