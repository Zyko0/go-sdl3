# GPU

Examples from https://github.com/TheSpydog/SDL_gpu_examples

Run `go run .`, optionally use example name as argument

## Building

Run `CGO_ENABLED=0 go build -o gpu .`

Can cross compile with `GOOS=windows` and test in Wine

Force Direct3D 12 using `SDL_GPU_DRIVER=direct3d12`
