package k8s_provider

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type Provider struct {
	clientSet *kubernetes.Clientset
	ns        string
}

type Endpoint struct {
	Address  string
	Port     int32
	PortName string
}

func New(ns, cfgPath string) (*Provider, error) {
	config, err := clientcmd.BuildConfigFromFlags("", cfgPath)
	if err != nil {
		return nil, err
	}

	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return &Provider{
		clientSet: clientSet,
		ns:        ns,
	}, nil
}

func (p *Provider) GetSvcList() ([]string, error) {
	svc, err := p.clientSet.CoreV1().Services(p.ns).List(metav1.ListOptions{})
	if err != nil {
		return []string{}, err
	}

	out := make([]string, len(svc.Items))
	for i := range svc.Items {
		out[i] = svc.Items[i].Name
	}

	return out, nil
}

func (p *Provider) GetSvcEndpoints(svcName string) ([]*Endpoint, error) {
	svc, err := p.clientSet.CoreV1().Services(p.ns).Get(svcName, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	out := make([]*Endpoint, len(svc.Spec.Ports))
	for i := range svc.Spec.Ports {
		p := svc.Spec.Ports[i]
		out[i] = &Endpoint{
			PortName:p.Name,
			Address: svc.Spec.ClusterIP,
			Port: p.Port,
		}
	}

	return out, nil
}
