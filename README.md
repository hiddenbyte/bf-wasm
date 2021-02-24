# bf-wasm
a Brainfuck compiler targeting WASM

`bf-wasm` compiles Brainfuck source code from the _standard input_ into a WASM module. The WASM module is written to the _standard ouput_. The WASM module is written in WebAssembly text format.

The [wat2wasm](https://github.com/WebAssembly/wabt) - from The WebAssembly Binary Toolkit - should be used to translate WebAssembly text format to the binary format.

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

The `bf-wasm` executable is available at the `./dist` folder after running `make build`.

## Usage
`bf-wasm` compiles Brainfuck source code from the _standard input_ into a WASM module. The WASM module is written to the _standard ouput_.

```bash
cat example.bf | bf-wasm > example.wat
```

### Examples

#### Hello World

**1. Create a Brainfuck source code file named `helloworld.bf` containg the following code.**

```brainfuck
[ This program prints "Hello World!" and a newline to the screen, its
  length is 106 active command characters. [It is not the shortest.]

  This loop is an "initial comment loop", a simple way of adding a comment
  to a BF program such that you don't have to worry about any command
  characters. Any ".", ",", "+", "-", "<" and ">" characters are simply
  ignored, the "[" and "]" characters just have to be balanced. This
  loop and the commands it contains are ignored because the current cell
  defaults to a value of 0; the 0 value causes this loop to be skipped.
]
++++++++               Set Cell #0 to 8
[
    >++++               Add 4 to Cell #1; this will always set Cell #1 to 4
    [                   as the cell will be cleared by the loop
        >++             Add 2 to Cell #2
        >+++            Add 3 to Cell #3
        >+++            Add 3 to Cell #4
        >+              Add 1 to Cell #5
        <<<<-           Decrement the loop counter in Cell #1
    ]                   Loop till Cell #1 is zero; number of iterations is 4
    >+                  Add 1 to Cell #2
    >+                  Add 1 to Cell #3
    >-                  Subtract 1 from Cell #4
    >>+                 Add 1 to Cell #6
    [<]                 Move back to the first zero cell you find; this will
                        be Cell #1 which was cleared by the previous loop
    <-                  Decrement the loop Counter in Cell #0
]                       Loop till Cell #0 is zero; number of iterations is 8

The result of this is:
Cell No :   0   1   2   3   4   5   6
Contents:   0   0  72 104  88  32   8
Pointer :   ^

>>.                     Cell #2 has value 72 which is 'H'
>---.                   Subtract 3 from Cell #3 to get 101 which is 'e'
+++++++..+++.           Likewise for 'llo' from Cell #3
>>.                     Cell #5 is 32 for the space
<-.                     Subtract 1 from Cell #4 for 87 to give a 'W'
<.                      Cell #3 was set to 'o' from the end of 'Hello'
+++.------.--------.    Cell #3 for 'rl' and 'd'
>>+.                    Add 1 to Cell #5 gives us an exclamation point
>++.                    And finally a newline from Cell #6
```

The code above was taken from [https://en.wikipedia.org/wiki/Brainfuck](https://en.wikipedia.org/wiki/Brainfuck).

**2. Compile `helloworld.bf` into WASM using bf-wasm**

```bash
cat helloworld.bf | bf-wasm > helloworld.wat
```

The `bf-wasm` produces the WASM module in text format. The WASM module needs to be translated to a binary format using [wat2wasm](https://github.com/WebAssembly/wabt).

```bash
wat2wasm helloworld.wat
```

**3. Copy `brainfck_env.js`**

`brainfck_env.js` is the brainfck WASM runtime environment.

```bash
curl https://raw.githubusercontent.com/hiddenbyte/bf-wasm/main/brainfck_env.js  > brainfck_env.js
```

**4. Create a HTML document named `index.html` with the following code**

```html
<!DOCTYPE html>
<html>
    <head>
        <title>Hello World</title>
    </head>
    
    <!-- Load brainfck WASM runtime environment -->
    <script src="brainfck_env.js"></script>

    <script>
            // This function is called when a brainfck program wants to read from std input. ',' command.
            const input = () => 0; // 

            // This function is called when a brainfck program wants to write to std output. '.' command.
            const output = data => { 
                document.write(String.fromCharCode(data)); 
            };

            // Create a brainfck WASM runtime environment with the specified I/O sources.
            const importObj = Brainfck.newImportObject(input, ouput);
            
            // Load and execute the compiled 'helloworld.bf' program
            WebAssembly.instantiateStreaming(fetch("helloworld.wasm"), importObj).then(obj => {
                obj.instance.exports.main(); // Execute
            });
    </script>
    <body>
    </body>
</html>
```

**5. Execute**

The `helloworld.wasm`, `brainfck_env.js` and `index.html` files should be in the same folder. These files should be served by a HTTP server.

Start a HTTP Server.

```bash
go get github.com/shurcooL/goexec
goexec 'http.ListenAndServe(`:8080`, http.FileServer(http.Dir(`.`)))' # Start a HTTP server
```

Open [http://localhost:8080/index.html](http://localhost:8080/index.html). 