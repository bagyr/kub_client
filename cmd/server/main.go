package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"

	"github.com/bagyr/kub_client/internal/service"
)


func main() {
	cfg := filepath.Join(os.Getenv("HOME"), ".kube", "dev_cluster_developers")

	//k8s, err := k8s_provider.New("travel", cfg)
	//if err != nil {
	//	panic(err.Error())
	//}

	adapter, err := service.New("travel", cfg)
	if err != nil {
		panic(err.Error())
	}

	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/services/", adapter.ServiceHandler())
	r.HandleFunc("/services/{svc}", EndpointHandler)
	http.Handle("/", r)
	fmt.Println("Serving ")

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err.Error())
	}

}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("under construction"))
}

func ServiceHandler(w http.ResponseWriter, r *http.Request) {


}

func EndpointHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	svcName, ok := vars["svc"]
	if !ok {
		svcName = "not_found"
	}
	w.WriteHeader(200)
	w.Write([]byte("hello " + svcName))
}
//func main() {
//	var kubeconfig *string
//	if home := homeDir(); home != "" {
//		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "dev_cluster_developers"), "(optional) absolute path to the kubeconfig file")
//	} else {
//		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
//	}
//	flag.Parse()
//
//	// use the current context in kubeconfig
//	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
//	if err != nil {
//		panic(err.Error())
//	}
//
//	// create the clientset
//	clientset, err := kubernetes.NewForConfig(config)
//	if err != nil {
//		panic(err.Error())
//	}
//
//	ns := "travel"
//	svc, err := clientset.CoreV1().Services(ns).List(metav1.ListOptions{LabelSelector: "app=travel-order-api"})
//	if err != nil {
//		panic(err.Error())
//	}
//
//	for _, s := range svc.Items {
//		fmt.Println(s.Name)
//	}
//
//	order, err := clientset.CoreV1().Services(ns).Get("travel-order-api", metav1.GetOptions{})
//	if err != nil {
//		panic(err.Error())
//	}
//	ipAddr := order.Spec.ClusterIP
//
//	for _, p := range order.Spec.Ports {
//		if strings.ToLower(p.Name) == "http" {
//			ipAddr += fmt.Sprintf(":%d", p.Port)
//		}
//	}
//
//}
//
//func homeDir() string {
//	if h := os.Getenv("HOME"); h != "" {
//		return h
//	}
//	return os.Getenv("USERPROFILE") // windows
//}
