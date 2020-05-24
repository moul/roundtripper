package roundtripper

import (
	"net/http"
)

type Transport struct {
	ExtraHeader http.Header
	Base        http.RoundTripper
}

func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	reqBodyClosed := false
	if req.Body != nil {
		defer func() {
			if !reqBodyClosed {
				req.Body.Close()
			}
		}()
	}

	req2 := new(http.Request)
	*req2 = *req
	req2.Header = make(http.Header, len(req.Header))
	for k, s := range req.Header {
		req2.Header[k] = append([]string(nil), s...)
	}
	for k, s := range t.ExtraHeader {
		if _, found := req2.Header[k]; !found {
			req2.Header[k] = []string(nil)
		}
		req2.Header[k] = append(req2.Header[k], s...)
	}

	// req.Body is assumed to be closed by the base RoundTripper.
	reqBodyClosed = true
	return http.DefaultTransport.RoundTrip(req2)
}
