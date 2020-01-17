module github.com/bagyr/kub_client

require (
	github.com/gorilla/mux v1.7.3
	github.com/imdario/mergo v0.3.8 // indirect
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d // indirect
	golang.org/x/time v0.0.0-20191024005414-555d28b269f0 // indirect
	k8s.io/apimachinery v0.17.1
	k8s.io/client-go v0.17.1
)

replace github.com/bagyr/kub_client/internal => ./internal

go 1.13
