# GRPC testing

### About
Testing GRPC to see if it could be of some use in the needed metadata flow in Open Core.  In particular
the need to transform some base RDF graphs/vocs into things like schema.org, DataCite and more.

### Steps
Installs
```
go get google.golang.org/grpc
go get -u github.com/golang/protobuf/protoc-gen-go
```

The compiler plugin, protoc-gen-go, will be installed in $GOBIN, defaulting to $GOPATH/bin. It must be in your $PATH for the protocol compiler, protoc, to find it.

```
export PATH=$PATH:$GOPATH/bin
```

Proto building
```
protoc -I helloworld/ helloworld/helloworld.proto --go_out=plugins=grpc:helloworld
```
