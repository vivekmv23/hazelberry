package types

import "fmt"

// Auth Types Supported
const (
	BASIC = "basic"
)

type AuthAttr struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Auth struct {
	Type  string     `json:"type"`
	Basic []AuthAttr `json:"basic"`
	creds map[string]string
}

func (a *Auth) InitAndValidate() error {
	if isValid := (a.Type != ""); !isValid {
		return fmt.Errorf("auth type is mandatory")
	}
	switch a.Type {
	case BASIC:
		if err := initAndValidateBasicAuthAttr(a); err != nil {
			return fmt.Errorf("basic auth has error: %s", err)
		}
	default:
		return fmt.Errorf("auth type \"%s\" is invalid/unsupported", a.Type)
	}
	return nil
}

func (a *Auth) IsEmpty() bool {
	return a.Type == ""
}

func initAndValidateBasicAuthAttr(ba *Auth) error {
	if len(ba.Basic) < 2 {
		return fmt.Errorf("values username & password are mandatory")
	}
	ba.creds = map[string]string{
		ba.Basic[0].Key: ba.Basic[0].Value,
		ba.Basic[1].Key: ba.Basic[1].Value,
	}

	if ba.creds["username"] == "" || ba.creds["password"] == "" {
		return fmt.Errorf("values username & password are mandatory")
	}
	return nil
}

func (a *Auth) getId() string {
	return a.creds["username"]
}

func (a *Auth) getPass() string {
	return a.creds["password"]
}
