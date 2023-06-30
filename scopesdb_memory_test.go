package scopes

import (
	"fmt"
	"strconv"
	"testing"
)

func Test_get_in_empty_database_should_fail(t *testing.T) {
	var db GatewayUseCases = NewInMemoryDB()
	u, e := GetURL("http://www.example.com/scopes.view")
	sn, e := db.Get(*u)
	if e == nil {
		t.Errorf("Should not find an non existing scope.")
	}

	if sn != nil {
		t.Errorf("Found a non-existing object!")
	}
}

func Test_create_and_get(t *testing.T) {
	var db GatewayUseCases = NewInMemoryDB()
	u, e := GetURL("http://www.example.com/scopes.view")
	if e != nil {
		t.Errorf("GetURL should pass, but failed with error %v", e)
	}
	scope := NewScopeName(*u, "view", "http://geenidee")
	scopeName, error := db.Create(*u, scope)

	if error != nil {
		t.Errorf("Creating scope in DB failed.")
	}
	if scopeName.URL != *u {
		t.Errorf("Scope url not equals.")
	}

	getDb, e2 := db.Get(*u)
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

	u, e := GetURL("http://www.example.com/scopes.view5")
	e = db.Delete(*u)
	if e != nil {
		t.Errorf("Failed to delete")
	}

	_, e2 := db.Get(*u)
	if e2 == nil {
		t.Errorf("Can still find the deleted object")
	}
}

func expectScopesInDb(gw GatewayUseCases, n int) {
	for i := 0; i < n; i++ {
		u, e := GetURL("http://www.example.com/scopes.view" + strconv.Itoa(i))
		scope := NewScopeName(*u, "view"+strconv.Itoa(i), "http://geenidee"+strconv.Itoa(i))
		_, e = gw.Create(*u, scope)
		if e != nil {
			fmt.Printf("Could not create a scope.")
		}
	}
}
