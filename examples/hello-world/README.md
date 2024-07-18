## Install tinygo

```
export PATH=/Users/rajatjindal/tinygo/bin:$PATH
export TINYGOROOT=~/tinygo
```

## Build wasm

```
tinygo build -target=wasip2 --wit-package $(go list -mod=readonly -m -f '{{.Dir}}' github.com/rajatjindal/wasm-go-sdk)/wit --wit-world cli -o main.wasm main.go
```

## Run

```
wasmtime run -Shttp main.wasm
```