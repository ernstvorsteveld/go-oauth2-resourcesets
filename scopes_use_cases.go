package scopes

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
