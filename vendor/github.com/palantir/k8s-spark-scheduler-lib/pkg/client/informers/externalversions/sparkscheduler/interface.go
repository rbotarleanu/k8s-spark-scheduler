// Code generated by informer-gen. DO NOT EDIT.

package sparkscheduler

import (
	internalinterfaces "github.com/palantir/k8s-spark-scheduler-lib/pkg/client/informers/externalversions/internalinterfaces"
	v1beta1 "github.com/palantir/k8s-spark-scheduler-lib/pkg/client/informers/externalversions/sparkscheduler/v1beta1"
	v1beta2 "github.com/palantir/k8s-spark-scheduler-lib/pkg/client/informers/externalversions/sparkscheduler/v1beta2"
)

// Interface provides access to each of this group's versions.
type Interface interface {
	// V1beta1 provides access to shared informers for resources in V1beta1.
	V1beta1() v1beta1.Interface
	// V1beta2 provides access to shared informers for resources in V1beta2.
	V1beta2() v1beta2.Interface
}

type group struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &group{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// V1beta1 returns a new v1beta1.Interface.
func (g *group) V1beta1() v1beta1.Interface {
	return v1beta1.New(g.factory, g.namespace, g.tweakListOptions)
}

// V1beta2 returns a new v1beta2.Interface.
func (g *group) V1beta2() v1beta2.Interface {
	return v1beta2.New(g.factory, g.namespace, g.tweakListOptions)
}
