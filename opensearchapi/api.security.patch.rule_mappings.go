package opensearchapi

import (
	"context"
	"io"
	"net/http"
	"strings"
)

func newBulkUpsertSecurityRuleMappingFunc(t Transport) BulkUpsertSecurityRuleMapping {
	return func(name string, body io.Reader, o ...func(*BulkUpsertSecurityRuleMappingRequest)) (*Response, error) {
		var r = BulkUpsertSecurityRuleMappingRequest{Name: name, Body: body}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

// BulkUpsertSecurityRuleMapping Bulk Upsert multiple role mappings
//
//	To use this API, you must have at least the manage_security cluster privilege.
//		https://opensearch.org/docs/2.3/security/access-control/api/#BulkUpsert-role-mapping
type BulkUpsertSecurityRuleMapping func(name string, body io.Reader, o ...func(*BulkUpsertSecurityRuleMappingRequest)) (*Response, error)

// BulkUpsertSecurityRuleMappingRequest configures the BulkUpsert Security Rule Mapping API request.
type BulkUpsertSecurityRuleMappingRequest struct {
	Name string

	Body io.Reader

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context
}

// Do will execute the request and returns response or error.
func (r BulkUpsertSecurityRuleMappingRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = http.MethodPatch

	path.Grow(len("/_plugins/_security/api/rolesmapping/"))
	path.WriteString("/_plugins/_security/api/rolesmapping/")

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

	req, err := newRequest(method, path.String(), r.Body)
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
func (f BulkUpsertSecurityRuleMapping) WithContext(v context.Context) func(*BulkUpsertSecurityRuleMappingRequest) {
	return func(r *BulkUpsertSecurityRuleMappingRequest) {
		r.ctx = v
	}
}

// WithPretty makes the response body pretty-printed.
func (f BulkUpsertSecurityRuleMapping) WithPretty() func(*BulkUpsertSecurityRuleMappingRequest) {
	return func(r *BulkUpsertSecurityRuleMappingRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f BulkUpsertSecurityRuleMapping) WithHuman() func(*BulkUpsertSecurityRuleMappingRequest) {
	return func(r *BulkUpsertSecurityRuleMappingRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f BulkUpsertSecurityRuleMapping) WithErrorTrace() func(*BulkUpsertSecurityRuleMappingRequest) {
	return func(r *BulkUpsertSecurityRuleMappingRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f BulkUpsertSecurityRuleMapping) WithFilterPath(v ...string) func(*BulkUpsertSecurityRuleMappingRequest) {
	return func(r *BulkUpsertSecurityRuleMappingRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f BulkUpsertSecurityRuleMapping) WithHeader(h map[string]string) func(*BulkUpsertSecurityRuleMappingRequest) {
	return func(r *BulkUpsertSecurityRuleMappingRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f BulkUpsertSecurityRuleMapping) WithOpaqueID(s string) func(*BulkUpsertSecurityRuleMappingRequest) {
	return func(r *BulkUpsertSecurityRuleMappingRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
