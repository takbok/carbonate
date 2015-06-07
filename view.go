package framework

import (
	"bytes"
)

type BaseView struct {
	base   string
	output bytes.Buffer
}
