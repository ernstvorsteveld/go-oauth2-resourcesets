package scopes

import (
	"net/url"
	"testing"
)

func TestCreate(t *testing.T) {
	db := NewInMemoryDB()
	url, _ := url.Parse("http://www.example.com/scopes/view")
	sn, e := db.Get(*url)

	if e == nil {
		t.Errorf("Should not find an non existing scope.")
	}

	if sn != nil {
		t.Errorf("Found a non-existing object!")
	}

	gu := getURL("http://www.example.com/scopes.view")
	scope := NewScopeName(gu, "view", getURL("http://geenidee"))
	scopeName, error := db.Create(*url, scope)

	if error != nil {
		t.Errorf("Creating scope in DB failed.")
	}
	if scopeName.scope != gu {
		t.Errorf("Scope url not equals.")
	}

	getDb, e2 := db.Get(*url)
	if e2 != nil {
		t.Errorf("Getting error")
	}
	if scopeName.scope != getDb.scope {
		t.Errorf("Wrong one returned.")
	}
}

func getURL(s string) url.URL {
	url, _ := url.Parse(s)
	return *url
}
