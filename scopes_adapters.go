package scopes

// Gateway are the possible actions with a Scope in the Gateways
type Gateway interface {
	Get(name URL) (*ScopeName, error)
	Create(name URL, scope ScopeName) (*ScopeName, error)
	Delete(name URL)
}
