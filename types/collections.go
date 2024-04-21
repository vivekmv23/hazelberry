package types

import "fmt"

type Collection struct {
	Info Info   `json:"info"`
	Item []Item `json:"item"`
	Auth Auth   `json:"auth"`
}

func (c *Collection) InitAndValidate() error {

	if err := c.Info.InitAndValidate(); err != nil {
		return fmt.Errorf("collection info has error: %s", err)
	}

	if len(c.Item) == 0 {
		return fmt.Errorf("collection item is mandatory")
	}

	for i := range c.Item {
		if err := c.Item[i].InitAndValidate(); err != nil {
			return fmt.Errorf("collection item at %d has error: %s", i, err)
		}
	}

	if !c.Auth.IsEmpty() {
		if err := c.Auth.InitAndValidate(); err != nil {
			return fmt.Errorf("collection auth has error: %s", err)
		}
	}

	return nil
}

func (c *Collection) IsEmpty() bool {
	return c.Info.IsEmpty()
}
