package consul

import (
	"github.com/go-kratos/gateway/discovery"
	"github.com/go-kratos/kratos/v2/registry"
	kuberegistry "github.com/zzz577/zeus-common/registry/kubernetes"
)

func init() {
	discovery.Register("kubernetes", New)
}

func New(_ string) (registry.Discovery, error) {
	return kuberegistry.NewRegistry()
}
