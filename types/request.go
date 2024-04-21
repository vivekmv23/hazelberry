package types

import (
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
		return fmt.Errorf("request url has error: %s", err)
	}

	if err := checkAndValidateMethod(r.Method); err != nil {
		return fmt.Errorf("request method has error: %s", err)
	}

	if err := checkAndValidate(&r.Auth); err != nil {
		return fmt.Errorf("request auth has error: %s", err)
	}

	if err := checkAndValidate(&r.Body); err != nil {
		return fmt.Errorf("request body has error: %s", err)

	}

	if len(r.Header) > 0 {
		for i := range r.Header {
			if err := r.Header[i].InitAndValidate(); err != nil {
				return fmt.Errorf("request header at %d has error: %s", i+1, err)
			}
		}
	}

	return nil
}

func (r *Request) IsEmpty() bool {
	return r.Method == ""
}

func checkAndValidate(tp Type) error {
	if tp.IsEmpty() {
		return nil
	}
	return tp.InitAndValidate()
}

func checkAndValidateUrl(r *Request) error {
	if r.UrlParsed == nil {
		return fmt.Errorf("url is mandatory")
	}
	if err := ConvertParsedUrl(r); err != nil {
		return err
	}
	return checkAndValidate(&r.Url)
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
		return fmt.Errorf("http method \"%s\" is invalid/unsupported", method)
	}
	return nil
}
