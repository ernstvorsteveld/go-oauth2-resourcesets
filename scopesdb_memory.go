package scopes

import (
	"fmt"
	"net/url"
)

type InMemoryDB struct {
	store map[url.URL]*ScopeNameDB
}

func NewInMemoryDB() *InMemoryDB {
	return &InMemoryDB {
		store: make(map[url.URL]*ScopeNameDB),
	}
}

func (db *InMemoryDB) Get(name url.URL) (*ScopeName, error) {
	scopeNameDB := db.store[name]
	if scopeNameDB == nil {
		return nil, fmt.Errorf("Scope with url %#v was not found", name)
	}
	return MapFrom(*scopeNameDB), nil
}

func (db *InMemoryDB) Create(name url.URL, scope ScopeName) (*ScopeName, error) {
	db.store[name] = MapTo(scope)
	return &scope, nil
}

func (db *InMemoryDB) Delete(name url.URL) {
	if db.store[name] != nil {
 		delete(db.store, name)
	}
}
