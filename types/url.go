package types

import (
	"fmt"
	"strings"
)

type Url struct {
	Raw        string       `json:"raw"`
	Protocol   string       `json:"protocol"`
	Host       interface{}  `json:"host"` // host is oneOf string || []string
	Port       string       `json:"port"`
	Query      []QueryParam `json:"query"`
	Variable   []Variable   `json:"variable"`
	hostString string
}

func (u *Url) InitAndValidate() error {
	// raw might not be mandatory as per schema
	if u.Raw == "" {
		return fmt.Errorf("raw is mandatory")
	}
	hostString, err := u.GetHost()
	if err != nil {
		return fmt.Errorf("host has error: %s", err)
	}
	// set the determined hostString for later use
	u.hostString = hostString

	if len(u.Query) > 0 {
		for i := range u.Query {
			if err := u.Query[i].InitAndValidate(); err != nil {
				return fmt.Errorf("query at %d has error: %s", i+1, err)
			}
		}
	}

	if len(u.Variable) > 0 {
		for i := range u.Variable {
			if err := u.Variable[i].InitAndValidate(); err != nil {
				return fmt.Errorf("variable at %d has error: %s", i+1, err)
			}
		}
	}

	return nil
}

func (u *Url) IsEmpty() bool {
	return u.Raw == ""
}

// gets the host string if not already set in hostString
func (u *Url) GetHost() (string, error) {
	if u.hostString != "" {
		return u.hostString, nil
	}
	// type determination to support string || []string
	switch u.Host.(type) {
	case nil:
		return "", nil
	case string:
		return u.Host.(string), nil
	case []interface{}:
		for _, subDomain := range u.Host.([]interface{}) {
			u.hostString += subDomain.(string) + "."
		}
		return strings.TrimSuffix(u.hostString, "."), nil
	default:
		return "", fmt.Errorf("url host cannot be of type %T", u.Host)
	}

}

type QueryParam struct {
	Key      string `json:"key"`
	Value    string `json:"value"`
	Disabled bool   `json:"disabled"`
}

func (qp *QueryParam) InitAndValidate() error {
	if qp.Key == "" {
		return fmt.Errorf("key is mandatory")
	}
	return nil
}

func (qp *QueryParam) GetKey() string {
	return qp.Key
}

func (qp *QueryParam) GetValue() string {
	return qp.Value
}

func (qp *QueryParam) IsDisabled() bool {
	return qp.Disabled
}
