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
	hostString string
}

func (u *Url) InitAndValidate() error {
	// raw might not be mandatory as per schema
	if u.Raw == "" {
		return fmt.Errorf("url.raw cannot be empty")
	}
	hostString, err := u.GetHost()
	if err != nil {
		return err
	}
	// set the determined hostString for later use
	u.hostString = hostString
	for _, qp := range u.Query {
		if err := qp.InitAndValidate(); err != nil {
			return err
		}
	}
	return nil
}

// gets the host string if not already set in hostString
func (u *Url) GetHost() (string, error) {
	if u.hostString != "" {
		return u.hostString, nil
	}
	// type determination to support string || []string
	switch u.Host.(type) {
	case string:
		return u.Host.(string), nil
	case []interface{}:
		for _, subDomain := range u.Host.([]interface{}) {
			u.hostString += subDomain.(string) + "."
		}
		return strings.TrimSuffix(u.hostString, "."), nil
	default:
		return "", fmt.Errorf("url.host cannot be of type %T", u.Host)
	}

}

type QueryParam struct {
	Key      string `json:"key"`
	Value    string `json:"value"`
	Disabled bool   `json:"disabled"`
}

func (qp *QueryParam) InitAndValidate() error {
	if qp.Key == "" {
		return fmt.Errorf("query parameter key is mandatory")
	}
	return nil
}

func (qp *QueryParam) getKey() string {
	return qp.Key
}

func (qp *QueryParam) getValue() string {
	return qp.Value
}
