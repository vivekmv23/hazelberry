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
	if b.Mode == "" {
		return fmt.Errorf("mode is mandatory")
	}
	switch b.Mode {
	case RAW:
		return nil
	default:
		return fmt.Errorf("mode \"%s\" is invalid/unsupported", b.Mode)
	}
}

func (b *Body) IsEmpty() bool {
	return b.Mode == ""
}
