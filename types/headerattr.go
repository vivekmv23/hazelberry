package types

import "fmt"

type HeaderAttr struct {
	Key      string `json:"key"`
	Value    string `json:"value"`
	Disabled bool   `json:"disabled"`
}

func (ha *HeaderAttr) InitAndValidate() error {
	// Values are also mandatory as per schema
	// Ignoring the validations to support empty string values
	if ha.Key == "" {
		return fmt.Errorf("key is mandatory")
	}
	return nil
}

func (ha *HeaderAttr) GetKey() string {
	return ha.Key
}

func (ha *HeaderAttr) GetValue() string {
	return ha.Value
}

func (ha *HeaderAttr) IsDisabled() bool {
	return ha.Disabled
}
