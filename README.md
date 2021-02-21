# bf-wasm
a brainf*ck compiler targeting wasm (text format)

## Installing

### go get

```bash
go get github.com/hiddenbyte/bf-wasm
```

## Running bf-wasm

```bash
echo ",++." | bf-wasm > example.wat # Produces a WebAssemply module in a text format
```