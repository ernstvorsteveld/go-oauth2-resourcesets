package scopes

import (
	"fmt"
	"net/url"
)

// InMemoryDB is the in-memory database
type InMemoryDB struct {
	store map[url.URL]*ScopeNameDB
}

// NewInMemoryDB Create a new InMemory DB
func NewInMemoryDB() *InMemoryDB {
	return &InMemoryDB{
		store: make(map[url.URL]*ScopeNameDB),
	}
}

// Get is for getting
func (db *InMemoryDB) Get(name url.URL) (*ScopeName, error) {
	scopeNameDB := db.store[name]
	if scopeNameDB == nil {
		return nil, fmt.Errorf("Scope with url %#v was not found", name)
	}
	return MapFrom(*scopeNameDB), nil
}

// Create is for creating
func (db *InMemoryDB) Create(name url.URL, scope ScopeName) (*ScopeName, error) {
	db.store[name] = MapTo(scope)
	return &scope, nil
}

// Delete is for deleting
func (db *InMemoryDB) Delete(name url.URL) {
	if db.store[name] != nil {
		delete(db.store, name)
	}
}
