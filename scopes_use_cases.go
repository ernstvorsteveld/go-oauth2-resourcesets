package scopes

// ScopeUseCases is the set of functions we can do with a Scope:
// Get: retrieve the scope by its URL, throws error when not available,
// Create: create for the URL a new Scope. If the URL already has a scope, it is overwritten,
// Delete: deletes the scope that belongs to the URL.
type ScopeUseCases interface {
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

// Delete the scope for the URL
func (g *Gateway) Delete(name URL) {
	g.gw.Delete(name)
}

func upsert(g *Gateway, name URL, scope Scope) {
	sn := ScopeName{
		URL:   name,
		Scope: scope,
	}
	g.gw.Create(name, sn)
}
