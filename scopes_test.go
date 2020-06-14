package scopes

import (
	"testing"
)

func Test_should_not_get_when_empty_database(t *testing.T) {
	var uc ScopeDescriptionUseCase = ScopeDescriptionUseCase{
		scopeDb: NewInMemoryDB(),
	}

	s, e := uc.Get(GetURL("https://not.oke.com"))
	if e == nil {
		t.Errorf("Should have an error, because Get should fail")
	}
	if s != nil {
		t.Errorf("Should not have an object that does not exist")
	}
}

func Test_should_create_and_get(t *testing.T) {
	var uc ScopeDescriptionUseCase = ScopeDescriptionUseCase{
		scopeDb: NewInMemoryDB(),
	}
	gu := GetURL("https://not.oke.com")
	scope := Scope{
		description: "view",
		iconURI:     GetURL("http://geenidee"),
	}
	uc.Create(gu, scope)
	s, e := uc.Get(gu)
	if e != nil {
		t.Errorf("Getting should work")
	}
	if s == nil {
		t.Errorf("Scope shoudl not be nil")
	}
}

func Test_should_be_able_to_delete_idempotent(t *testing.T) {
	var uc ScopeDescriptionUseCase = ScopeDescriptionUseCase{
		scopeDb: NewInMemoryDB(),
	}

	uc.Delete(GetURL("https://not.oke.com"))
	uc.Delete(GetURL("https://not.oke.com"))
	uc.Delete(GetURL("https://not.oke.com"))
}
