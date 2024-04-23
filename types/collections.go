package types

import "fmt"

type Collection struct {
	Info     Info       `json:"info"`
	Item     []Item     `json:"item"`
	Auth     Auth       `json:"auth"`
	Varaible []Variable `json:"variable"`
}

func (c *Collection) InitAndValidate() error {

	if err := c.Info.InitAndValidate(); err != nil {
		return fmt.Errorf("collection info has error: %s", err)
	}

	name := c.Info.Name

	if len(c.Item) == 0 {
		return fmt.Errorf("%s collection has no items", name)
	}

	for i := range c.Item {
		if err := c.Item[i].InitAndValidate(); err != nil {
			return fmt.Errorf("%s collection's item at %d has error: %s", name, i+1, err)
		}
	}

	for i := range c.Varaible {
		if err := c.Varaible[i].InitAndValidate(); err != nil {
			return fmt.Errorf("%s collection's variable at %d has error: %s", name, i+1, err)
		}
	}

	if !c.Auth.IsEmpty() {
		if err := c.Auth.InitAndValidate(); err != nil {
			return fmt.Errorf("%s collection's auth has error: %s", name, err)
		}
	}

	return nil
}

func (c *Collection) IsEmpty() bool {
	return c.Info.IsEmpty()
}
