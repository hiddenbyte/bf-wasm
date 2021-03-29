# Example: Hello World

## Requirements

* bf-wasm
* wat2wasm (from The WebAssembly Binary Toolkit - https://github.com/WebAssembly/wabt)

## Build

### Compile `helloworld.bf` into WASM

```bash
cat helloworld.bf | bf-wasm > helloworld.wat
```

`bf-wasm` produces a WASM module in text format. The WASM module needs to be translated to a binary format using `wat2wasm`.

```bash
wat2wasm helloworld.wat
```

## Run

### Start a HTTP Server

```bash
go get github.com/shurcooL/goexec
goexec 'http.ListenAndServe(`:8080`, http.FileServer(http.Dir(`.`)))' # Start a HTTP server
```

Open [http://localhost:8080/index.html](http://localhost:8080/index.html).