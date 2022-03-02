module mux

go 1.17

require (
	CRUD/models v0.0.0-00010101000000-000000000000
	github.com/gorilla/mux v1.8.0
	github.com/gorilla/sessions v1.2.1
)

require (
	github.com/go-redis/redis v6.15.9+incompatible // indirect
	github.com/gorilla/securecookie v1.1.1 // indirect
	github.com/nxadm/tail v1.4.8 // indirect
	golang.org/x/crypto v0.0.0-20220214200702-86341886e292 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace CRUD/models => ./models
