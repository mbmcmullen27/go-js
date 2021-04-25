# go-js

to build:

```sh
GOOS=js GOARCH=wasm go build -o main.wasm
cp main.wasm ~/js/wasm-test/
```

- /js/wasm-test holds a local copy of (go-wasm)[https://github.com/mbmcmullen27/go-wasm]
