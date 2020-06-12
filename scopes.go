package scopes

import "net/url"

type ScopeName struct {
	scope url.URL
	Scope
}

type Scope struct {
	name    string
	iconUri url.URL
}

type ScopeDescription interface {
	Get(name url.URL) (*Scope, error)
	Create(name url.URL, scope Scope)
	Delete(name url.URL)
}

type ScopeDescriptionUseCase struct {
	scopeDb ScopeDescription
}

func (s *ScopeDescriptionUseCase) Get(name url.URL) (*Scope, error) {
	return nil, nil
}

func NewScopeName(u url.URL, n string, i url.URL) ScopeName {
	return ScopeName { u, Scope {n, i}}
}
