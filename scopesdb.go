package scopes

import "net/url"

// ScopeNameDB is the type that is stored in the DB
type ScopeNameDB struct {
	URL     string
	ScopeDB ScopeDB
}

// ScopeDB is the type that is stored in the DB
type ScopeDB struct {
	Description string
	IconURI     string
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
		URL: scope.URL.String(),
		ScopeDB: ScopeDB{
			Description: scope.Scope.Description,
			IconURI:     scope.Scope.IconURI.String(),
		},
	}
}

// MapFrom is to map a DB Scope to a Scope
func MapFrom(scope ScopeNameDB) *ScopeName {
	return &ScopeName{
		URL: GetURL(scope.URL),
		Scope: Scope{
			Description: scope.ScopeDB.Description,
			IconURI:     GetURL(scope.ScopeDB.IconURI),
		},
	}
}
