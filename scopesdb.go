package scopes

import "net/url"

type ScopeNameDB struct {
	scope url.URL
	ScopeDB
}

type ScopeDB struct {
	name    string
	iconUri url.URL
}

type ScopeDbUseCase interface {
	Get(name url.URL) (*ScopeName, error)
	Create(name url.URL, scope ScopeName) (*ScopeName, error)
	Delete(name url.URL)
}

func MapTo(scope ScopeName) *ScopeNameDB {
	return &ScopeNameDB{scope.scope, ScopeDB{scope.name, scope.iconUri}}
}

func MapFrom(scope ScopeNameDB) *ScopeName {
	return &ScopeName{scope.scope, Scope{scope.name, scope.iconUri}}
}