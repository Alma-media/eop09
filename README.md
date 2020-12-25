# eop09

[![License][license-badge]][license-link]
[![CI][ci-badge]][ci-link]
[![Report Card][report-badge]][report-link]

Golang microservices. Task for "two hours".

## Installation
```bash
$ go get github.com/Alma-media/eop09
```

## Run
- server:
```bash
$ go run main.go -port 5050
```
- client:
```bash
$  go run main.go -file path/to/file -grpc-addr localhost:5050 -http-port 8080
```

## Further improvements

### Client

1. API
- HTTPs support
- advanced routing
- http middleware (jwt, auth, timeout etc)
- enable codec middleware and retrieve codec from the context to use corresponding encoder
- pass logger to the handler (intermal errors should be logged but not exposed to the user)

2. Config
- introduce config `struct` and use a library to parse the values from flags, toml, yaml, json, hcl ...

### Server

1. Service
- TLS
- server side interceptors
- register reflection API
- graceful shutdown

2. Storage
- choose another storage implementation that uses the context and is able to fail / return an error (e.g. database, blockchain etc)


[license-badge]: https://img.shields.io/:license-MIT-green.svg
[license-link]: https://opensource.org/licenses/MIT
[ci-badge]: https://github.com/Alma-media/eop09/workflows/.github/workflows/tests.yaml/badge.svg
[ci-link]: https://github.com/Alma-media/eop09/actions
[report-badge]: https://goreportcard.com/badge/github.com/Alma-media/eop09
[report-link]: https://goreportcard.com/report/github.com/Alma-media/eop09

