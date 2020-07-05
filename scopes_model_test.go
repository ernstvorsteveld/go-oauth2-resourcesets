package scopes

import (
	"testing"
)

func Test_NewScope(t *testing.T) {
	expected := "http://www.example.com/scopes/view"
	var s = NewScope("view", expected)
	if s.Description != "view" {
		t.Errorf("The descriptionn is not initialized correctly, expected %s, but have %s", "view", s.Description)
	}
	if s.IconURI.URL.String() != expected {
		t.Errorf("Scope iconURL is incorrect, expected %s, have %s\n", expected, s.IconURI.URL.String())
	}
}
