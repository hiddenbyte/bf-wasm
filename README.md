# bf-wasm
a Brainfuck compiler targeting WASM

`bf-wasm` compiles Brainfuck source code from the _standard input_ into a WASM module. The WASM module is written to the _standard ouput_. The WASM module is written in WebAssembly text format. The [wat2wasm](https://github.com/WebAssembly/wabt) - from The WebAssembly Binary Toolkit - should be used translate WebAssembly text format to the binary format.

**What is Brainfuck?**

>  Brainfuck is an esoteric programming language created in 1993 by Urban Müller.

(from [https://en.wikipedia.org/wiki/Brainfuck](https://en.wikipedia.org/wiki/Brainfuck)) 

## Installing

### Using `go get`

```bash
go get github.com/hiddenbyte/bf-wasm
```

### Build yourself

```bash
make build # build bf-wasm
```

The `bf-wasm` executable is available at `./dist` after running `make build`.

## Running bf-wasm

```bash
echo ",++." | bf-wasm > example.wat # produces a WebAssemply module in a text format
```