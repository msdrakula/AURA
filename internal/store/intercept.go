package store

import "meb/internal/httpio"

// InterceptRequest parks a request in the UI queue if intercept is enabled.
func (s *Store) InterceptRequest(id string, req *httpio.Request) (string, string) {
	flow := &Flow{ID: id, Created: nowTS(), Request: req}
	action := s.Intercept(flow, "request")
	return action, flow.EditedRequest
}

// InterceptResponse parks a response in the UI queue if intercept is enabled.
func (s *Store) InterceptResponse(id string, req *httpio.Request, resp *httpio.Response) (string, string) {
	flow := &Flow{ID: id, Created: nowTS(), Request: req, Response: resp}
	action := s.Intercept(flow, "response")
	return action, flow.EditedResponse
}
