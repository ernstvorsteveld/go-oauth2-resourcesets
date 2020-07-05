package scopes

// Controller is the entrypoint for web/devices doing requests
type Controller interface {
}

// GatewayUseCases are the possible actions with a Scope in the Gateways
type GatewayUseCases interface {
	Get(name URL) (*ScopeName, error)
	Create(name URL, scope ScopeName) (*ScopeName, error)
	Upsert(name URL, scope ScopeName) (*ScopeName, error)
	Delete(name URL)
}
