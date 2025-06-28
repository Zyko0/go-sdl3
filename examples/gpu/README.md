# GPU

Examples from https://github.com/TheSpydog/SDL_gpu_examples

Run `go run ./examples` or add example name as argument

## Building

Run `CGO_ENABLED=0 go build -o gpu ./examples`

Can cross compile with `GOOS=windows` and test in Wine

<!-- SDL_HINT_GPU_DRIVER=d3d12 -->
