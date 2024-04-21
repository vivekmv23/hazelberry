package types

import "fmt"

type Info struct {
	Name string `json:"name"`
}

func (inf *Info) InitAndValidate() error {
	if inf.Name == "" {
		return fmt.Errorf("info name is mandatory")
	}
	return nil
}

func (inf *Info) IsEmpty() bool {
	return inf.Name == ""
}
