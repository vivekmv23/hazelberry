package types

import "fmt"

type Variable struct {
	Id       string `json:"id"`
	Key      string `json:"key"`
	Value    string `json:"value"`
	Name     string `json:"name"`
	Disabled bool   `json:"disabled"`
}

func (v *Variable) InitAndValidate() error {
	if v.Key == "" && v.Id == "" {
		return fmt.Errorf("key/id is mandatory")
	}
	return nil
}

func (v *Variable) GetKey() string {
	if v.Id != "" {
		return v.Id
	} else {
		return v.Key
	}
}

func (v *Variable) GetValue() string {
	return v.Value
}

func (v *Variable) IsDisabled() bool {
	return v.Disabled
}
