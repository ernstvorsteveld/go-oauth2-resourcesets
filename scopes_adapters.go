package scopes

// GatewayUseCases are the possible actions with a Scope in the Gateways
type GatewayUseCases interface {
	Get(name URL) (*ScopeName, error)
	Create(name URL, scope ScopeName) (*ScopeName, error)
	Upsert(name URL, scope ScopeName) (*ScopeName, error)
	Delete(name URL) error
}

// ControllerUseCases defines the Controller use cases
type ControllerUseCases interface {
	Get(name URL) ([]byte, error)
	Create(b []byte) ([]byte, error)
	Update(b []byte) ([]byte, error)
	Delete(name URL) error
}
