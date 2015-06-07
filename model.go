package framework

type Model interface {
	ControllerChild
}

type BaseModel struct {
}

func (m *BaseModel) SetParent(c *BaseController) {
	c.model = m
}
