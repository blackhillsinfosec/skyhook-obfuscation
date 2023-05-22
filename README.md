# Skyhook Obfuscators

This repository provides obfuscation objects and algorithms for [Skyhook](https://github.com/blackhillsinfosec/skyhook).

# Building WASM

This command can be used to build the WASM file. It can be dropped
into Skyhook directly, or incorporated into other tooling.

```bash
cd wasm
GOOS=js GOARCH=wasm go build -o algos.wasm
```