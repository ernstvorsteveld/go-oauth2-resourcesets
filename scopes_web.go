package scopes

import "encoding/json"

// Controler is containing the use cases
type Controler struct {
	uc UseCases
}

// Get gets the scope by url
func (c *Controler) Get(name URL) ([]byte, error) {
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
func (c *Controler) Create(b []byte) ([]byte, error) {
	return nil, nil
}

// Update (partially) updates a scope
func (c *Controler) Update(b []byte) ([]byte, error) {
	return nil, nil
}

// Delete removes a scope with name
func (c *Controler) Delete(name URL) error {
	c.uc.Delete(name) // old version does not return error yet!!!
	return nil
}
