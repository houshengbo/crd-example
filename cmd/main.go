package main

import (
	"flag"
	"fmt"

	"github.com/golang/glog"
	"github.com/houshengbo/crd-example/pkg/apis/star/v1alpha1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
)

func main() {
	flag.Parse()
	var err error

	var config *rest.Config
	if config, err = rest.InClusterConfig(); err != nil {
		glog.Fatalf("error creating the client configuration: %v", err)
	}

	// Create a new clientset which include our CRD schema
	crdclient, err := v1alpha1.NewClient(config)
	if err != nil {
		panic(err)
	}

	// Create a new Star object
	star := &v1alpha1.Star{
		ObjectMeta: meta_v1.ObjectMeta{
			Name:   "sun",
			Labels: map[string]string{"testlabel": "test"},
		},
		Spec: v1alpha1.StarSpec{
			Type:   "G",
			Location:    "sun",
		},
		Status: v1alpha1.StarStatus{
			State:   "Created",
			Message: "Created",
		},
	}
	// Create the Star object we create above in the k8s cluster
	resp, err := crdclient.Stars("default").Create(star)
	if err != nil {
		fmt.Printf("error creating the CR: %v\n", err)
	} else {
		fmt.Printf("CR created: %v\n", resp)
	}

	obj, err := crdclient.Stars("default").Get(star.ObjectMeta.Name)
	if err != nil {
		glog.Infof("error while getting the CR %v\n", err)
	}
	fmt.Printf("CR Star Found: \n%v\n", obj)
}
