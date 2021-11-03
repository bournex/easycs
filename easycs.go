package easycs

import (
	"bytes"
	"context"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Response struct {
	Status int
	Body   []byte
}

type EasyC struct {
	c *http.Client

	method string          // http method
	url    string          // whole url with query strings, default is http://localhost:8000/
	scheme string          // protocol scheme, default is http
	host   string          // http host field, default is localhost
	path   string          // http path, default is /
	body   []byte          // request body, default is nil
	header http.Header     // request header, default is nil
	param  url.Values      // query parameters, default is nil
	form   url.Values      // form data, if both body and form are specified, form will be perform
	ctx    context.Context // request context
}

// Do the request
func (ec *EasyC) Do() (*http.Response, error) {
	if ec.method == "" {
		ec.method = "GET"
	}

	if ec.url == "" {
		if ec.scheme == "" {
			ec.scheme = "http"
		}
		if ec.host == "" {
			ec.host = "localhost:8000"
		}
		if ec.path == "" {
			ec.path = "/"
		}
		ec.url = ec.scheme + "://" + ec.host + ec.path
	}

	if len(ec.param) > 0 {
		ec.url = ec.url + "?" + ec.param.Encode()
	}

	var reader io.Reader
	if len(ec.body) == 0 {
		if len(ec.form) > 0 {
			reader = strings.NewReader(ec.form.Encode())
		}
	} else {
		reader = bytes.NewReader(ec.body)
	}

	var (
		request *http.Request
		err     error
	)

	defer func() {
		*ec = EasyC{}
	}()

	if ec.ctx != nil {
		request, err = http.NewRequestWithContext(ec.ctx, ec.method, ec.url, reader)
	} else {
		request, err = http.NewRequest(ec.method, ec.url, reader)
	}

	if err != nil {
		return nil, err
	}

	if len(ec.header) > 0 {
		request.Header = http.Header(ec.header)
	}

	if ec.c != nil {
		return ec.c.Do(request)
	}

	return http.DefaultClient.Do(request)
}

//
func (ec *EasyC) DoWithStatus(f func(*Response, error)) {
	response, err := ec.Do()
	if err != nil {
		f(nil, err)
	} else {
		defer response.Body.Close()
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			f(nil, err)
		} else {
			f(&Response{response.StatusCode, body}, nil)
		}
	}
}

// async request with callback wrapped
func (ec *EasyC) Done(f func(*http.Response, error)) {
	go func() {
		f(ec.Do())
	}()
}

// async request with callback wrapped with status and response body parameters
func (ec *EasyC) DoneWithStatus(f func(*Response, error)) {
	go func() {
		ec.DoWithStatus(f)
	}()
}

func (ec *EasyC) WithContext(ctx context.Context) *EasyC {
	ec.ctx = ctx
	return ec
}

// default is GET
func (ec *EasyC) WithMethod(method string) *EasyC {
	ec.method = method
	return ec
}

// default is http
func (ec *EasyC) WithScheme(scheme string) *EasyC {
	ec.scheme = scheme
	return ec
}

// default is localhost
func (ec *EasyC) WithHost(host string) *EasyC {
	ec.host = host
	return ec
}

// default is /
func (ec *EasyC) WithPath(path string) *EasyC {
	ec.path = path
	return ec
}

// full path with scheme, host, url path
func (ec *EasyC) WithUrl(url string) *EasyC {
	ec.url = url
	return ec
}

// default is http.DefaultClient from net/http package
func (ec *EasyC) WithClient(c *http.Client) *EasyC {
	ec.c = c
	return ec
}

// default is nil
func (ec *EasyC) WithBody(body []byte) *EasyC {
	ec.body = body
	return ec
}

// query string key-value pair
func (ec *EasyC) WithQuery(key, val string) *EasyC {
	if ec.param == nil {
		ec.param = url.Values{}
	}
	ec.param.Add(key, val)
	return ec
}

// query string object
func (ec *EasyC) WithQuerys(params url.Values) *EasyC {
	ec.param = params
	return ec
}

// form data key-value pair
func (ec *EasyC) WithForm(key, val string) *EasyC {
	if ec.form == nil {
		ec.form = url.Values{}
	}
	ec.form.Add(key, val)
	return ec
}

// form data object
func (ec *EasyC) WithForms(forms url.Values) *EasyC {
	ec.form = forms
	return ec
}

// header key-value pair
func (ec *EasyC) WithHeader(key, val string) *EasyC {
	if ec.header == nil {
		ec.header = http.Header{}
	}
	ec.header.Add(key, val)
	return ec
}

// header object
func (ec *EasyC) WithHeaders(headers http.Header) *EasyC {
	ec.header = headers
	return ec
}
