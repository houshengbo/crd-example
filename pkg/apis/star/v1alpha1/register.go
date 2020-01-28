package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/client-go/rest"
)

const (
	CRDPlural   string = "stars"
	CRDGroup    string = "example.crd.com"
	CRDVersion  string = "v1alpha1"
)

var SchemeGroupVersion = schema.GroupVersion{Group: CRDGroup, Version: CRDVersion}

func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion,
		&Star{},
		&StarList{},
	)
	meta_v1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}


func NewClient(cfg *rest.Config) (*StarV1Alpha1Client, error) {
	scheme := runtime.NewScheme()
	SchemeBuilder := runtime.NewSchemeBuilder(addKnownTypes)
	if err := SchemeBuilder.AddToScheme(scheme); err != nil {
		return nil, err
	}

	codecs := serializer.NewCodecFactory(scheme)
	config := *cfg
	config.GroupVersion = &SchemeGroupVersion
	config.APIPath = "/apis"
	config.ContentType = runtime.ContentTypeJSON

	config.NegotiatedSerializer = codecs.WithoutConversion()
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &StarV1Alpha1Client{restClient: client}, nil
}
