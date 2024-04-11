package types

import "fmt"

// Auth Types
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
		return fmt.Errorf("field auth.type is mandatory")
	}
	switch a.Type {
	case BASIC:
		return initAndValidateBasicAuthAttr(a)
	default:
		return fmt.Errorf("auth type is invalid or not yet supported")
	}

}

func initAndValidateBasicAuthAttr(ba *Auth) error {
	isValid := (len(ba.Basic) == 2)
	if !isValid {
		return fmt.Errorf("values username & password are mandatory for basic authentication")
	}
	ba.creds = map[string]string{
		ba.Basic[0].Key: ba.Basic[0].Value,
		ba.Basic[1].Key: ba.Basic[1].Value,
	}
	isValid = isValid && (ba.creds["username"] != "") && (ba.creds["password"] != "")
	if !isValid {
		return fmt.Errorf("values username & password are mandatory for basic authentication")
	}
	return nil
}

func (a *Auth) getId() string {
	return a.creds["username"]
}

func (a *Auth) getPass() string {
	return a.creds["password"]
}
