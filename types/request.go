package types

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Request Methods Supported
var supportedMethods []string = []string{
	http.MethodGet,
	http.MethodPost,
	http.MethodPatch,
	http.MethodPut,
	http.MethodDelete,
	http.MethodHead,
	http.MethodOptions,
}

type Request struct {
	UrlParsed interface{}  `json:"url"`
	Auth      Auth         `json:"auth"`
	Method    string       `json:"method"`
	Header    []HeaderAttr `json:"header"` // header is oneOf string||header-list. Not supporting string yet
	Body      Body         `json:"body"`
	Url       Url
}

func (r *Request) InitAndValidate() error {

	if err := checkAndValidateUrl(r); err != nil {
		return err
	}

	if err := checkAndValidateMethod(r.Method); err != nil {
		return err
	}

	if err := checkAndValidate(&r.Auth); err != nil {
		return err
	}

	if err := checkAndValidate(&r.Body); err != nil {
		return err

	}

	for _, header := range r.Header {
		if err := header.InitAndValidate(); err != nil {
			return err
		}
	}

	return nil
}

func checkAndValidate(tp Type) error {
	if tp.IsEmpty() {
		return nil
	}
	return tp.InitAndValidate()
}

func checkAndValidateUrl(r *Request) error {
	if r.UrlParsed == nil {
		return fmt.Errorf("Url is mandatory")
	}
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
		return checkAndValidate(&r.Url)
	}
}

func checkAndValidateMethod(method string) error {
	isSupportedMethod := false
	for _, supportedMethod := range supportedMethods {
		if supportedMethod == method {
			isSupportedMethod = true
			break
		}
	}

	if !isSupportedMethod {
		return fmt.Errorf("http method %s is not supported", method)
	}
	return nil
}
