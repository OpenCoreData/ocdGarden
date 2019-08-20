BINARY := sqlbuilder
VERSION :=`cat VERSION`
.DEFAULT_GOAL := sqlbuilder

gbuilder:
	cd cmd/$(BINARY) ; \
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 env go build -o gbuilder

sqlbuilder:
	cd cmd/$(BINARY) ; \
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 env go build -o $(BINARY)
