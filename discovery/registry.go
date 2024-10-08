package discovery

import (
	"fmt"
	"github.com/go-kratos/kratos/v2/registry"
)

var globalRegistry = NewRegistry()

type Factory func(scheme string) (registry.Discovery, error)

// Registry is the interface for callers to get registered middleware.
type Registry interface {
	Register(name string, factory Factory)
	Create(scheme string) (registry.Discovery, error)
}

type discoveryRegistry struct {
	discovery map[string]Factory
}

// NewRegistry returns a new middleware registry.
func NewRegistry() Registry {
	return &discoveryRegistry{
		discovery: map[string]Factory{},
	}
}

func (d *discoveryRegistry) Register(name string, factory Factory) {
	d.discovery[name] = factory
}

func (d *discoveryRegistry) Create(scheme string) (registry.Discovery, error) {
	factory, ok := d.discovery[scheme]
	if !ok {
		return nil, fmt.Errorf("discovery %s has not been registered", scheme)
	}

	impl, err := factory(scheme)
	if err != nil {
		return nil, fmt.Errorf("create discovery error: %s", err)
	}
	return impl, nil
}

// Register registers one discovery.
func Register(name string, factory Factory) {
	globalRegistry.Register(name, factory)
}

// Create instantiates a discovery based on scheme.
func Create(scheme string) (registry.Discovery, error) {
	return globalRegistry.Create(scheme)
}
