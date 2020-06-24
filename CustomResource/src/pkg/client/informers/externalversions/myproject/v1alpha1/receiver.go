/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	myprojectv1alpha1 "github.com/cisco/CustomResource/src/pkg/apis/myproject/v1alpha1"
	versioned "github.com/cisco/CustomResource/src/pkg/client/clientset/versioned"
	internalinterfaces "github.com/cisco/CustomResource/src/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/cisco/CustomResource/src/pkg/client/listers/myproject/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ReceiverInformer provides access to a shared informer and lister for
// Receivers.
type ReceiverInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ReceiverLister
}

type receiverInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewReceiverInformer constructs a new informer for Receiver type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewReceiverInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredReceiverInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredReceiverInformer constructs a new informer for Receiver type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredReceiverInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SampleprojectV1alpha1().Receivers(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SampleprojectV1alpha1().Receivers(namespace).Watch(context.TODO(), options)
			},
		},
		&myprojectv1alpha1.Receiver{},
		resyncPeriod,
		indexers,
	)
}

func (f *receiverInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredReceiverInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *receiverInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&myprojectv1alpha1.Receiver{}, f.defaultInformer)
}

func (f *receiverInformer) Lister() v1alpha1.ReceiverLister {
	return v1alpha1.NewReceiverLister(f.Informer().GetIndexer())
}