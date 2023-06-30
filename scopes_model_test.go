package scopes

import (
	"encoding/json"
	"fmt"
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

func Test_mashalling(t *testing.T) {
	expectedURI := "http://www.example.com/api/resourcesets/s1"
	expectedIconURI := "http://icon.example.com/s1"
	expectedDescription := "scope s1"
	var source []byte = []byte("{ \"url\" : \"" + expectedURI + "\",\"scope\" : { \"description\" : \"" + expectedDescription + "\", \"icon_uri\" : \"" + expectedIconURI + "\"}}")
	var sn ScopeName = ScopeName{}
	var e error
	e = json.Unmarshal(source, &sn)
	if e != nil {
		t.Errorf("Unmarshalling of string failed, %s", e)
	}
	if sn.URL.String() != expectedURI {
		t.Errorf("Unmarshalling failed, expected URL %s, got %s", expectedURI, sn.URL.String())
	}
	if sn.Scope.Description != expectedDescription {
		t.Errorf("Unmarshalling failed, expected description %s, got %s", expectedDescription, sn.Scope.Description)
	}

	if sn.Scope.IconURI.String() != expectedIconURI {
		t.Errorf("Unarshalling failed, expected icon uri %s, got %s", expectedIconURI, sn.Scope.IconURI)
	}
}

func Test_unmarshalling(t *testing.T) {
	expectedURI := ""
	expectedIconURI := "http://icon.example.com/s1"
	expectedDescription := "scope s1"
	document := `{ "url" : "` + expectedURI + `", "scope" : { "description" : "` + expectedDescription + `", "icon_uri" : "` + expectedIconURI + `"}}`
	fmt.Printf("Document: \n%s\n", document)
	var source []byte = []byte(document)
	var sn ScopeName = ScopeName{}
	var e error
	e = json.Unmarshal(source, &sn)
	if e == nil {
		t.Errorf("Unmarshalling error expected!")
	}
}
