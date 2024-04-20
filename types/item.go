package types

import "fmt"

type Item struct {
	Id      string  `json:"id"`
	Name    string  `json:"name"`
	Request Request `json:"request"` // not supportiing request of type string
}

func (i *Item) InitAndValidate() error {
	if i.Name == "" {
		return fmt.Errorf("item name is mandatory")
	}
	if i.Request.IsEmpty() {
		return fmt.Errorf("item request is mandatory")
	}
	return i.Request.InitAndValidate()

}

func (i *Item) IsEmpty() bool {
	return i.Name == ""
}
