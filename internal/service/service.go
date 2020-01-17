package service

import (
	"encoding/json"
	"net/http"

	"github.com/bagyr/kub_client/internal/k8s_provider"
)

type K8sAdapter struct {
	k8s *k8s_provider.Provider
}

func New(ns, cfg string) (*K8sAdapter, error) {
	k8s, err := k8s_provider.New("travel", cfg)
	if err != nil {
		return nil, err
	}

	return &K8sAdapter{
		k8s: k8s,
	}, nil
}

func (a *K8sAdapter) ServiceHandler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		svc, err := a.k8s.GetSvcList()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		out, err := json.Marshal(svc)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(out)
	}
}