package scopes

import (
	"fmt"
	"strconv"
	"testing"
)

func Test_get_in_empty_database_should_fail(t *testing.T) {
	var db GatewayUseCases = NewInMemoryDB()
	gu := GetURL("http://www.example.com/scopes.view")
	sn, e := db.Get(gu)

	if e == nil {
		t.Errorf("Should not find an non existing scope.")
	}

	if sn != nil {
		t.Errorf("Found a non-existing object!")
	}
}

func Test_create_and_get(t *testing.T) {
	var db GatewayUseCases = NewInMemoryDB()
	gu := GetURL("http://www.example.com/scopes.view")
	scope := NewScopeName(gu, "view", "http://geenidee")
	scopeName, error := db.Create(gu, scope)

	if error != nil {
		t.Errorf("Creating scope in DB failed.")
	}
	if scopeName.URL != gu {
		t.Errorf("Scope url not equals.")
	}

	getDb, e2 := db.Get(gu)
	if e2 != nil {
		t.Errorf("Getting error")
	}
	if scopeName.URL != getDb.URL {
		t.Errorf("Wrong one returned.")
	}
}

func Test_create_and_delete(t *testing.T) {
	var db GatewayUseCases = NewInMemoryDB()
	expectScopesInDb(db, 100)

	gu := GetURL("http://www.example.com/scopes.view5")
	var e error = db.Delete(gu)
	if e != nil {
		t.Errorf("Failed to delete")
	}

	_, e2 := db.Get(gu)
	if e2 == nil {
		t.Errorf("Can still find the deleted object")
	}
}

func expectScopesInDb(gw GatewayUseCases, n int) {
	for i := 0; i < n; i++ {
		gu := GetURL("http://www.example.com/scopes.view" + strconv.Itoa(i))
		scope := NewScopeName(gu, "view"+strconv.Itoa(i), "http://geenidee"+strconv.Itoa(i))
		_, e := gw.Create(gu, scope)
		if e != nil {
			fmt.Printf("Could not create a scope.")
		}
	}
}
