package scopes

import "net/url"

// ScopeName contains the url of the scope and the scope itself
type ScopeName struct {
	url   url.URL
	scope Scope
}

// Scope is the human understandable name/description and the icon url
type Scope struct {
	description string
	iconURI     url.URL
}

// ScopeDescription is the set of functions we can do with a Scope:
// Get: retrieve the scope by its URL, throws error when not available,
// Create: create for the URL a new Scope. If the URL already has a scope, it is overwritten,
// Delete: deletes the scope that belongs to the URL.
type ScopeDescription interface {
	Get(name url.URL) (*Scope, error)
	Create(name url.URL, scope Scope)
	Delete(name url.URL)
}

// ScopeDescriptionUseCase is the use case
type ScopeDescriptionUseCase struct {
	ScopeDb ScopeDbUseCase
}

// Get the scope by its URL
func (s *ScopeDescriptionUseCase) Get(name url.URL) (*Scope, error) {
	sn, error := s.ScopeDb.Get(name)

	if error != nil {
		return nil, error
	}
	return &Scope{sn.scope.description, sn.scope.iconURI}, nil
}

// Create the scope for an URL
func (s *ScopeDescriptionUseCase) Create(name url.URL, scope Scope) {
	sn := ScopeName{
		url:   name,
		scope: scope,
	}
	s.ScopeDb.Create(name, sn)
}

// Delete the scope for the URL
func (s *ScopeDescriptionUseCase) Delete(name url.URL) {
	s.ScopeDb.Delete(name)
}

// NewScopeName is to be used for creating a new scope
func NewScopeName(u url.URL, s string, i url.URL) ScopeName {
	return ScopeName{
		url:   u,
		scope: NewScope(s, i),
	}
}

// NewScope to create a new Scope
func NewScope(s string, u url.URL) Scope {
	return Scope{
		description: s,
		iconURI:     u,
	}
}
