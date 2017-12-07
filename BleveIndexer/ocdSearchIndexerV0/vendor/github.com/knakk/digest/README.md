## Digest


Package digest is an implementation of http.RoundTripper that handles digest authentication. The source is
taken from https://github.com/swook/mlab-ns2-rtt/tree/master/ns/digest

Updated to fix some minor issues.

### Original copyright notice

<pre>
Copyright 2013 M-Lab

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

The digest package provides an implementation of http.RoundTripper that takes
care of HTTP Digest Authentication (http://www.ietf.org/rfc/rfc2617.txt).
This only implements the MD5 and "auth" portions of the RFC, but that covers
the majority of avalible server side implementations including apache web
server.
</pre>

### Example usage
```go
t := NewTransport("myUserName", "myP@55w0rd")
req, err := http.NewRequest("GET", "http://notreal.com/path?arg=1", nil)
if err != nil {
	return err
}
resp, err := t.RoundTrip(req)
if err != nil {
	return err
}
```

Or it can be used as a client:

```go
c, err := t.Client()
if err != nil {
	return err
}
resp, err := c.Get("http://notreal.com/path?arg=1")
if err != nil {
	return err
}
```
