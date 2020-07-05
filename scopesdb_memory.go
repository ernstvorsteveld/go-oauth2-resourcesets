package scopes

import (
	"fmt"
)

// InMemoryDB is the in-memory database
type InMemoryDB struct {
	store map[URL]*ScopeNameDB
}

// NewInMemoryDB Create a new InMemory DB
func NewInMemoryDB() *InMemoryDB {
	return &InMemoryDB{
		store: make(map[URL]*ScopeNameDB),
	}
}

// Get is for getting
func (db *InMemoryDB) Get(name URL) (*ScopeName, error) {
	scopeNameDB := db.store[name]
	if scopeNameDB == nil {
		return nil, fmt.Errorf("Scope with url %#v was not found", name)
	}
	return MapFrom(*scopeNameDB), nil
}

// Create is for creating
func (db *InMemoryDB) Create(name URL, scope ScopeName) (*ScopeName, error) {
	return createOrUpdate(db, name, scope)
}

// Upsert will create when new otherwise update
func (db *InMemoryDB) Upsert(name URL, scope ScopeName) (*ScopeName, error) {
	return createOrUpdate(db, name, scope)
}

func createOrUpdate(db *InMemoryDB, name URL, scope ScopeName) (*ScopeName, error) {
	db.store[name] = MapTo(scope)
	return &scope, nil
}

// Delete is for deleting
func (db *InMemoryDB) Delete(name URL) {
	if db.store[name] != nil {
		delete(db.store, name)
	}
}
