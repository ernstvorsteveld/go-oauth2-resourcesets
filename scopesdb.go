package scopes

// ScopeNameDB is the type that is stored in the DB
type ScopeNameDB struct {
	URL     URL
	ScopeDB ScopeDB
}

// ScopeDB is the type that is stored in the DB
type ScopeDB struct {
	Description string
	IconURI     URL
}

// ScopeDbUseCase are the possible actions with a Scope in the DB
type ScopeDbUseCase interface {
	Get(name URL) (*ScopeName, error)
	Create(name URL, scope ScopeName) (*ScopeName, error)
	Delete(name URL)
}

// MapTo is used to map a Scope to a DB Scope
func MapTo(scope ScopeName) *ScopeNameDB {
	u := scope.Scope.IconURI
	return &ScopeNameDB{
		URL: scope.URL,
		ScopeDB: ScopeDB{
			Description: scope.Scope.Description,
			IconURI:     *u,
		},
	}
}

// MapFrom is to map a DB Scope to a Scope
func MapFrom(scope ScopeNameDB) *ScopeName {
	u := scope.ScopeDB.IconURI
	return &ScopeName{
		URL: scope.URL,
		Scope: Scope{
			Description: scope.ScopeDB.Description,
			IconURI:     &u,
		},
	}
}
