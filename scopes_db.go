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
