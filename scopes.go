package scopes

import (
	"encoding/json"
	"net/url"
)

// URL is type for url's
type URL struct {
	URL url.URL
}

func (u *URL) String() string {
	if u != nil {
		return u.URL.String()
	}
	return ""
}

// MarshalJSON for marshalling to json
func (u *URL) MarshalJSON() ([]byte, error) {
	if u == nil {
		return []byte("null"), nil
	}
	return json.Marshal(u.String())
}

// UnmarshalJSON for unmarshalling URL
func (u *URL) UnmarshalJSON(b []byte) error {
	URL, error := unMarshallString(b)
	if error != nil {
		return error
	}
	url := GetURL(URL)
	u.URL = url.URL
	return nil
}

func unMarshallString(b []byte) (string, error) {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return "", err
	}
	return s, nil
}

// ScopeName contains the url of the scope and the scope itself
type ScopeName struct {
	URL   URL   `json:"url"`
	Scope Scope `json:"scope"`
}

// Scope is the human understandable name/description and the icon url
type Scope struct {
	Description string `json:"description"`
	IconURI     *URL   `json:"icon_uri"`
}

// UseCases is the set of functions we can do with a Scope:
// Get: retrieve the scope by its URL, throws error when not available,
// Create: create for the URL a new Scope. If the URL already has a scope, it is overwritten,
// Delete: deletes the scope that belongs to the URL.
type UseCases interface {
	Get(name URL) (*Scope, error)
	Create(name URL, scope Scope)
	Upsert(name URL, scope Scope)
	Delete(name URL)
}

// Gateway is the use case
type Gateway struct {
	gw GatewayUseCases
}

// Get the scope by its URL
func (g *Gateway) Get(name URL) (*Scope, error) {
	sn, error := g.gw.Get(name)

	if error != nil {
		return nil, error
	}
	return &Scope{sn.Scope.Description, sn.Scope.IconURI}, nil
}

// Create the scope for an URL
func (g *Gateway) Create(name URL, scope Scope) {
	upsert(g, name, scope)
}

// Upsert will create when new, update when exists
func (g *Gateway) Upsert(name URL, scope Scope) {
	upsert(g, name, scope)
}

func upsert(g *Gateway, name URL, scope Scope) {
	sn := ScopeName{
		URL:   name,
		Scope: scope,
	}
	g.gw.Create(name, sn)
}

// Delete the scope for the URL
func (g *Gateway) Delete(name URL) {
	g.gw.Delete(name)
}

// NewScopeName is to be used for creating a new scope
func NewScopeName(u URL, s string, i string) ScopeName {
	return ScopeName{
		URL:   u,
		Scope: NewScope(s, i),
	}
}

// NewScope to create a new Scope
func NewScope(s string, u string) Scope {
	url := GetURL(u)
	return Scope{
		Description: s,
		IconURI:     &url,
	}
}

// GetURL makes a url of a string value
func GetURL(s string) URL {
	u, _ := url.Parse(s)
	url := URL{URL: *u}
	return url
}
