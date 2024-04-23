package types

import "fmt"

// Auth Types Supported
const (
	BASIC   = "basic"
	NO_AUTH = "noauth"
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
		return fmt.Errorf("type is mandatory")
	}
	switch a.Type {
	case BASIC:
		if err := initAndValidateBasicAuthAttr(a); err != nil {
			return fmt.Errorf("type basic has error: %s", err)
		}
	case NO_AUTH:
		return nil // requires no specific validation
	default:
		return fmt.Errorf("type \"%s\" is invalid/unsupported", a.Type)
	}
	return nil
}

func (a *Auth) IsEmpty() bool {
	return a.Type == "" || a.Type == NO_AUTH
}

func initAndValidateBasicAuthAttr(ba *Auth) error {
	if len(ba.Basic) < 2 {
		return fmt.Errorf("username & password are mandatory")
	}
	ba.creds = map[string]string{
		ba.Basic[0].Key: ba.Basic[0].Value,
		ba.Basic[1].Key: ba.Basic[1].Value,
	}

	if ba.creds["username"] == "" || ba.creds["password"] == "" {
		return fmt.Errorf("username & password are mandatory")
	}
	return nil
}

func (a *Auth) GetId() string {
	return a.creds["username"]
}

func (a *Auth) GetPass() string {
	return a.creds["password"]
}
