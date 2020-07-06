package scopes

import "encoding/json"

// WebError is the error object returned on an error
type WebError struct {
	Request       string `json:"request_url"`
	DateTime      string `json:"date_time"`
	CorrelationID string `json:"correlation_id"`
	Code          string `json:"error_code"`
	Description   string `json:"error_description"`
}

// Controller is containing the use cases
type Controller struct {
	uc ScopeUseCases
}

// Get gets the scope by url
func (c *Controller) Get(name URL) ([]byte, error) {
	var s, e1 = c.uc.Get(name)
	if e1 != nil {
		// handle the error!
		return nil, e1
	}
	var b, e2 = json.Marshal(s)
	if e2 != nil {
		//handle the error
		return nil, e2
	}
	return b, nil
}

// Create creates a new scope
func (c *Controller) Create(b []byte) ([]byte, error) {
	return nil, nil
}

// Update (partially) updates a scope
func (c *Controller) Update(b []byte) ([]byte, error) {
	return nil, nil
}

// Delete removes a scope with name
func (c *Controller) Delete(name URL) error {
	c.uc.Delete(name) // old version does not return error yet!!!
	return nil
}
