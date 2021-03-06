// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	"context"
	time "time"

	crdexamplev1 "github.com/tlm/crd-example/pkg/apis/crdexample/v1"
	versioned "github.com/tlm/crd-example/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/tlm/crd-example/pkg/generated/informers/externalversions/internalinterfaces"
	v1 "github.com/tlm/crd-example/pkg/generated/listers/crdexample/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// TweetInformer provides access to a shared informer and lister for
// Tweets.
type TweetInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.TweetLister
}

type tweetInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewTweetInformer constructs a new informer for Tweet type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewTweetInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredTweetInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredTweetInformer constructs a new informer for Tweet type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredTweetInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CrdexampleV1().Tweets(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CrdexampleV1().Tweets(namespace).Watch(context.TODO(), options)
			},
		},
		&crdexamplev1.Tweet{},
		resyncPeriod,
		indexers,
	)
}

func (f *tweetInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredTweetInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *tweetInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&crdexamplev1.Tweet{}, f.defaultInformer)
}

func (f *tweetInformer) Lister() v1.TweetLister {
	return v1.NewTweetLister(f.Informer().GetIndexer())
}
