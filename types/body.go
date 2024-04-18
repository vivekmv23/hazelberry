package types

import (
	"fmt"
)

// Body Modes Supported
const (
	RAW = "raw"
)

type Body struct {
	Mode string `json:"mode"`
	Raw  string `json:"raw"`
}

func (b *Body) InitAndValidate() error {
	if isValid := b.Mode != ""; !isValid {
		return fmt.Errorf("field body.mode is mandatory")
	}
	switch b.Mode {
	case RAW:
		return nil
	default:
		return fmt.Errorf("body.mode %s is invalid or not yet supported", b.Mode)
	}
}

func (b *Body) IsEmpty() bool {
	return b.Mode == ""
}
