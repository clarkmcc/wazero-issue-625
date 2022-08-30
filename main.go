package main

import (
	"context"
	_ "embed"
	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/wasi_snapshot_preview1"
	"os"
)

//go:generate cargo build --target wasm32-wasi --release
//go:embed target/wasm32-wasi/release/issue_625_repr.wasm
var wasm []byte

func main() {
	ctx := context.Background()
	r := wazero.NewRuntime(ctx)
	defer r.Close(ctx)
	_, err := wasi_snapshot_preview1.Instantiate(ctx, r)
	if err != nil {
		panic(err)
	}
	compiled, err := r.CompileModule(ctx, wasm, wazero.NewCompileConfig())
	if err != nil {
		panic(err)
	}
	module, err := r.InstantiateModule(ctx, compiled, wazero.NewModuleConfig().
		WithStdout(os.Stdout).
		WithStderr(os.Stderr))
	if err != nil {
		panic(err)
	}
	_, err = module.ExportedFunction("entrypoint").Call(ctx)
	if err != nil {
		panic(err)
	}
}
