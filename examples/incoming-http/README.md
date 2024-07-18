## Install tinygo

```
export PATH=/Users/rajatjindal/tinygo/bin:$PATH
export TINYGOROOT=~/tinygo
```

## Build wasm

```
$ tinygo build -target=wasip2 --wit-package $(go list -mod=readonly -m -f '{{.Dir}}' github.com/rajatjindal/wasi-go-sdk)/wit --wit-world sdk -o main.wasm main.go
```

## Run

```
wasmtime serve -Scli main.wasm
Serving HTTP on http://0.0.0.0:8080/
```

## Curl

```
$ curl http://0.0.0.0:8080/
{"timestamp":1721269951458,"fact":"Bats are the only mammals that can fly"}
```
