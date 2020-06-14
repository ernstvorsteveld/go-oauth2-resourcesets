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
