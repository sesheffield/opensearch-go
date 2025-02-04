package opensearchapi

import (
	"context"
	"net/http"
	"strings"
)

func newGetSecurityRoleMappingFunc(t Transport) GetSecurityRoleMapping {
	return func(name string, o ...func(*GetSecurityRoleMappingRequest)) (*Response, error) {
		var r = GetSecurityRoleMappingRequest{Name: name}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

// GetSecurityRoleMapping Gets a role mapping
//
//	To use this API, you must have at least the manage_security cluster privilege.
type GetSecurityRoleMapping func(name string, o ...func(*GetSecurityRoleMappingRequest)) (*Response, error)

// GetSecurityRoleMappingRequest configures the Get Security Rule Mapping API request.
type GetSecurityRoleMappingRequest struct {
	Name string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context
}

// Do will execute the request and returns response or error.
func (r GetSecurityRoleMappingRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = http.MethodGet

	path.Grow(len("/_plugins/_security/api/rolesmapping/") + len(r.Name))
	path.WriteString("/_plugins/_security/api/rolesmapping/")
	path.WriteString(r.Name)

	params = make(map[string]string)
	if r.Pretty {
		params["pretty"] = "true"
	}

	if r.Human {
		params["human"] = "true"
	}

	if r.ErrorTrace {
		params["error_trace"] = "true"
	}

	if len(r.FilterPath) > 0 {
		params["filter_path"] = strings.Join(r.FilterPath, ",")
	}

	req, err := newRequest(method, path.String(), nil)
	if err != nil {
		return nil, err
	}

	if len(params) > 0 {
		q := req.URL.Query()
		for k, v := range params {
			q.Set(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	if len(r.Header) > 0 {
		if len(req.Header) == 0 {
			req.Header = r.Header
		} else {
			for k, vv := range r.Header {
				for _, v := range vv {
					req.Header.Add(k, v)
				}
			}
		}
	}

	if ctx != nil {
		req = req.WithContext(ctx)
	}

	res, err := transport.Perform(req)
	if err != nil {
		return nil, err
	}

	response := Response{
		StatusCode: res.StatusCode,
		Body:       res.Body,
		Header:     res.Header,
	}

	return &response, nil
}

// WithContext sets the request context.
func (f GetSecurityRoleMapping) WithContext(v context.Context) func(*GetSecurityRoleMappingRequest) {
	return func(r *GetSecurityRoleMappingRequest) {
		r.ctx = v
	}
}

// WithPretty makes the response body pretty-printed.
func (f GetSecurityRoleMapping) WithPretty() func(*GetSecurityRoleMappingRequest) {
	return func(r *GetSecurityRoleMappingRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f GetSecurityRoleMapping) WithHuman() func(*GetSecurityRoleMappingRequest) {
	return func(r *GetSecurityRoleMappingRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f GetSecurityRoleMapping) WithErrorTrace() func(*GetSecurityRoleMappingRequest) {
	return func(r *GetSecurityRoleMappingRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f GetSecurityRoleMapping) WithFilterPath(v ...string) func(*GetSecurityRoleMappingRequest) {
	return func(r *GetSecurityRoleMappingRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f GetSecurityRoleMapping) WithHeader(h map[string]string) func(*GetSecurityRoleMappingRequest) {
	return func(r *GetSecurityRoleMappingRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f GetSecurityRoleMapping) WithOpaqueID(s string) func(*GetSecurityRoleMappingRequest) {
	return func(r *GetSecurityRoleMappingRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
