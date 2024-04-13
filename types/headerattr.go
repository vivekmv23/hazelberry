package types

import "fmt"

type HeaderAttr struct {
	Key      string `json:"key"`
	Value    string `json:"value"`
	Disabled bool `json:"disabled"`
}

func (ha *HeaderAttr) InitAndValidate() error {
	// Values are also mandatory as per schema 
	// Ignoring the validations to support empty string values
	if ha.Key == "" {
		return fmt.Errorf("header key is mandatory")
	}
	return nil
}

func (ha *HeaderAttr) getKey() string {
	return ha.Key
}

func (ha *HeaderAttr) getValue() string {
	return ha.Value
}




