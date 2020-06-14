package scopes

import "net/url"

// ScopeNameDB is the type that is stored in the DB
type ScopeNameDB struct {
	url     url.URL
	scopeDB ScopeDB
}

// ScopeDB is the type that is stored in the DB
type ScopeDB struct {
	description string
	iconURI     url.URL
}

// ScopeDbUseCase are the possible actions with a Scope in the DB
type ScopeDbUseCase interface {
	Get(name url.URL) (*ScopeName, error)
	Create(name url.URL, scope ScopeName) (*ScopeName, error)
	Delete(name url.URL)
}

// MapTo is used to map a Scope to a DB Scope
func MapTo(scope ScopeName) *ScopeNameDB {
	return &ScopeNameDB{
		url: scope.url,
		scopeDB: ScopeDB{
			description: scope.scope.description,
			iconURI:     scope.scope.iconURI},
	}
}

// MapFrom is to map a DB Scope to a Scope
func MapFrom(scope ScopeNameDB) *ScopeName {
	return &ScopeName{
		url: scope.url,
		scope: Scope{
			description: scope.scopeDB.description,
			iconURI:     scope.scopeDB.iconURI,
		},
	}
}
