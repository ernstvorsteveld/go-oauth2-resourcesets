package scopes

import (
	"encoding/json"
	"testing"
)

func Test_should_not_get_when_empty_database(t *testing.T) {
	var uc = Gateway{
		gw: NewInMemoryDB(),
	}

	var u, e = GetURL("https://not.oke.com")
	s, e := uc.Get(*u)
	if e == nil {
		t.Errorf("Should have an error, because Get should fail")
	}
	if s != nil {
		t.Errorf("Should not have an object that does not exist")
	}
}

func Test_should_create_and_get(t *testing.T) {
	var g Gateway = Gateway{
		gw: NewInMemoryDB(),
	}
	var gu, e = GetURL("https://not.oke.com")
	if e != nil {
		t.Errorf("GetURL should pass, but failed with error %v", e)
	}
	u, e := GetURL("http://geenidee")
	scope := Scope{
		Description: "view",
		IconURI:     u,
	}
	g.Create(*gu, scope)
	s, e := g.Get(*gu)
	if e != nil {
		t.Errorf("Getting should work")
	}
	if s == nil {
		t.Errorf("Scope shoudl not be nil")
	}
}

func Test_should_be_able_to_delete_idempotent(t *testing.T) {
	var g Gateway = Gateway{
		gw: NewInMemoryDB(),
	}

	u, _ := GetURL("https://not.oke.com")
	g.Delete(*u)
	g.Delete(*u)
	g.Delete(*u)
}

func Test_should_marshall_and_unmarshall(t *testing.T) {
	n := "view"
	u := "http://www.example.com/icons/view.png"
	var s = NewScope(n, u)
	var sb, e = json.Marshal(s)
	if e != nil {
		t.Errorf("Marshalling failed, %v\n", e)
	}
	expected := "{\"description\":\"view\",\"icon_uri\":\"http://www.example.com/icons/view.png\"}"
	if string(sb) != expected {
		t.Errorf("Marshalling failed, expected %s, have %s\n", expected, string(sb))
	}

	var scope Scope = Scope{}
	e = json.Unmarshal([]byte(expected), &scope)
	if e != nil {
		t.Errorf("UnMarshalling of a scope failed, %v\n", e)
	}

	if n != scope.Description {
		t.Errorf("Description value is incorrect, expected %s, have %s\n", n, scope.Description)
	}
	if u != scope.IconURI.URL.String() {
		t.Errorf("URL value is incorrect, expected %s, have %s\n", u, scope.IconURI.URL.String())
	}
}
