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

protoc is in my home bin dir...
```
export PATH=$PATH:$GOPATH/bin
export PATH=~/bin/:$PATH
```

Proto building  (no REST gateway option in proto file)
```
protoc -I helloworld/ helloworld/helloworld.proto --go_out=plugins=grpc:helloworld
```

Proto building  (REST gateway option in proto file)
```
protoc -I /usr/local/include -I helloworld/ -I$GOPATH/src    -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc:. helloworld/helloworld.proto

```

Building the python library with protoc
```
export SRC_DIR=source_dir
~/bin/protoc -I=$SRC_DIR --python_out=./libraries $SRC_DIR/bleve.proto

To get client code needed..  you should really use
~/anaconda/bin/python -m grpc_tools.protoc -I./protobufs --python_out=./libraries/ --grpc_python_out=./libraries/ ./protobufs/bleve.proto
```


### REST gateway

```
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
go get -u github.com/golang/protobuf/protoc-gen-go
```




### Copyright
The original source code in this directory is just my learning modification of 
Google code.  That code can only be exposed under the following conditions.

```
/*
 *
 * Copyright 2015, Google Inc.
 * All rights reserved.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions are
 * met:
 *
 *     * Redistributions of source code must retain the above copyright
 * notice, this list of conditions and the following disclaimer.
 *     * Redistributions in binary form must reproduce the above
 * copyright notice, this list of conditions and the following disclaimer
 * in the documentation and/or other materials provided with the
 * distribution.
 *     * Neither the name of Google Inc. nor the names of its
 * contributors may be used to endorse or promote products derived from
 * this software without specific prior written permission.
 *
 * THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
 * "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
 * LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
 * A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
 * OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
 * SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
 * LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
 * DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
 * THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
 * (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
 * OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 *
 */
```