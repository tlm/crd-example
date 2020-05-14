package main

import (
	"flag"
	"time"

	crdexample "github.com/tlm/crd-example/pkg/apis/crdexample/v1"
	clientset "github.com/tlm/crd-example/pkg/generated/clientset/versioned"
	informers "github.com/tlm/crd-example/pkg/generated/informers/externalversions"
	"github.com/tlm/crd-example/pkg/signal"

	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog"
)

var (
	kubeconfig string
	masterURL  string
)

func main() {
	klog.InitFlags(nil)
	flag.Parse()

	stopCh := signals.SetupSignalHandler()

	cfg, err := clientcmd.BuildConfigFromFlags(masterURL, kubeconfig)
	if err != nil {
		klog.Fatalf("error building kubeconfig: %v", err)
	}

	exampleClient, err := clientset.NewForConfig(cfg)
	if err != nil {
		klog.Fatalf("error building kubernetes clientset: %v", err)
	}

	factory := informers.NewSharedInformerFactory(exampleClient, time.Second*30)
	factory.Crdexample().V1().Tweets().Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			tweet, ok := obj.(*crdexample.Tweet)
			if !ok {
				klog.Error("error converting informer obj to tweet")
				return
			}
			klog.Infof("new tweet %s with message %s", tweet.Name, tweet.Spec.Message)
		},
		DeleteFunc: func(obj interface{}) {
			tweet, ok := obj.(*crdexample.Tweet)
			if !ok {
				klog.Error("error converting informer obj to tweet")
				return
			}
			klog.Infof("deleted tweet %s with message %s", tweet.Name, tweet.Spec.Message)
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			klog.Infof("in update function")
		},
	})

	factory.Start(stopCh)

	select {
	case <-stopCh:
	}
}

func init() {
	flag.StringVar(&kubeconfig, "kubeconfig", "", "Path to a kubeconfig. Only required if out-of-cluster.")
	flag.StringVar(&masterURL, "master", "", "The address of the Kubernetes API server. Overrides any value in kubeconfig. Only required if out-of-cluster.")
}
