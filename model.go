package brahma

// The model interface
type Model interface {
	ControllerChild
}

// The base model struct
type BaseModel struct {
}

// Set the parent controller for the model
func (m *BaseModel) setParent(c *BaseController) {
	c.model = m
}
