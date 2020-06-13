package scopes

import "net/url"

// ScopeName contains the url of the scope and the scope itself
type ScopeName struct {
	scope url.URL
	Scope
}

// Scope is the name and the icon url
type Scope struct {
	name    string
	iconURI url.URL
}

// ScopeDescription is the set of functions we have available on a Scope
type ScopeDescription interface {
	Get(name url.URL) (*Scope, error)
	Create(name url.URL, scope Scope)
	Delete(name url.URL)
}

// ScopeDescriptionUseCase is the use case
type ScopeDescriptionUseCase struct {
	scopeDb ScopeDescription
}

// Get is for retrieval of a Scope
func (s *ScopeDescriptionUseCase) Get(name url.URL) (*Scope, error) {
	return nil, nil
}

// Create is for creating a scope
func (s *ScopeDescriptionUseCase) Create(name url.URL, scope Scope) {
}

// Delete is for deleting a scope
func (s *ScopeDescriptionUseCase) Delete(name url.URL) {
}

// NewScopeName is to be used for creating a new scope
func NewScopeName(u url.URL, n string, i url.URL) ScopeName {
	return ScopeName{u, Scope{n, i}}
}
