const Brainfck = {
    newImportObject: (inputFunc, outputFunc) => ({
        io: {
            readInput: () => inputFunc(),
            writeOutput: (data) => { outputFunc(data); }
        }
    })
};