package scopes

import (
	"encoding/json"
	"net/url"

	"github.com/pkg/errors"
)

// URL is type for urls
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
		return errors.Wrap(&json.SyntaxError{
			Offset: 0,
		}, "Unmarshalling failed")
	}
	url, _ := GetURL(URL)
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

// NewScopeName is to be used for creating a new scope
func NewScopeName(u URL, s string, i string) ScopeName {
	return ScopeName{
		URL:   u,
		Scope: NewScope(s, i),
	}
}

// NewScope to create a new Scope
func NewScope(s string, u string) Scope {
	url, _ := GetURL(u)
	return Scope{
		Description: s,
		IconURI:     url,
	}
}

// GetURL makes a url of a string value
func GetURL(s string) (*URL, error) {
	u, err := url.Parse(s)
	if err != nil {
		return nil, err
	}
	url := URL{URL: *u}
	return &url, nil
}
