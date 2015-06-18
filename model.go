package framework

// The model interface
type Model interface {
	ControllerChild
}

// The base model struct
type BaseModel struct {
}

// Set the parent controller for the model
func (m *BaseModel) SetParent(c *BaseController) {
	c.model = m
}
