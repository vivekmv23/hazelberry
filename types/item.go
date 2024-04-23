package types

import "fmt"

type Item struct {
	Id       string     `json:"id"`
	Name     string     `json:"name"`
	Request  Request    `json:"request"` // not supportiing request of type string
	Variable []Variable `json:"variable"`
}

func (itm *Item) InitAndValidate() error {
	if itm.Name == "" {
		return fmt.Errorf("name is mandatory")
	}
	if itm.Request.IsEmpty() {
		return fmt.Errorf("request is mandatory")
	}
	if err := itm.Request.InitAndValidate(); err != nil {
		return fmt.Errorf("requests has error: %s", err)
	}

	if len(itm.Variable) > 0 {
		for i := range itm.Variable {
			if err := itm.Variable[i].InitAndValidate(); err != nil {
				return fmt.Errorf("variable at %d has error: %s", i+1, err)
			}
		}
	}

	return nil

}

func (i *Item) IsEmpty() bool {
	return i.Name == ""
}
