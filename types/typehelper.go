package types

import (
	"encoding/json"
)

func ConvertParsedUrl(r *Request) error {
	switch r.UrlParsed.(type) {
	case string:
		r.Url = Url{Raw: r.UrlParsed.(string)}
		return nil
	default:
		mar, err := json.Marshal(r.UrlParsed)
		if err != nil {
			return err
		}
		if err := json.Unmarshal(mar, &r.Url); err != nil {
			return err
		}
		return nil
	}
}
