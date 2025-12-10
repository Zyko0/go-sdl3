//go:build windows || unix

package shadercross

import (
	"runtime"
	"unsafe"

	"github.com/Zyko0/go-sdl3/internal"
	"github.com/Zyko0/go-sdl3/sdl"
	puregogen "github.com/Zyko0/purego-gen"
	"github.com/ebitengine/purego"
)

var (
	// Library handles
	_hnd_shadercross uintptr

	// Symbols
	_addr_ShaderCross_Init                            uintptr
	_addr_ShaderCross_Quit                            uintptr
	_addr_ShaderCross_GetSPIRVShaderFormats           uintptr
	_addr_ShaderCross_TranspileMSLFromSPIRV           uintptr
	_addr_ShaderCross_TranspileHLSLFromSPIRV          uintptr
	_addr_ShaderCross_CompileDXBCFromSPIRV            uintptr
	_addr_ShaderCross_CompileDXILFromSPIRV            uintptr
	_addr_ShaderCross_CompileGraphicsShaderFromSPIRV  uintptr
	_addr_ShaderCross_CompileComputePipelineFromSPIRV uintptr
	_addr_ShaderCross_ReflectGraphicsSPIRV            uintptr
	_addr_ShaderCross_ReflectComputeSPIRV             uintptr
	_addr_ShaderCross_GetHLSLShaderFormats            uintptr
	_addr_ShaderCross_CompileDXBCFromHLSL             uintptr
	_addr_ShaderCross_CompileDXILFromHLSL             uintptr
	_addr_ShaderCross_CompileSPIRVFromHLSL            uintptr
)

func initialize() {
	var err error

	// Symbols
	_addr_ShaderCross_Init, err = puregogen.OpenSymbol(_hnd_shadercross, "SDL_ShaderCross_Init")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: SDL_ShaderCross_Init: " + err.Error())
	}
	_addr_ShaderCross_Quit, err = puregogen.OpenSymbol(_hnd_shadercross, "SDL_ShaderCross_Quit")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: SDL_ShaderCross_Quit: " + err.Error())
	}
	_addr_ShaderCross_GetSPIRVShaderFormats, err = puregogen.OpenSymbol(_hnd_shadercross, "SDL_ShaderCross_GetSPIRVShaderFormats")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: SDL_ShaderCross_GetSPIRVShaderFormats: " + err.Error())
	}
	_addr_ShaderCross_TranspileMSLFromSPIRV, err = puregogen.OpenSymbol(_hnd_shadercross, "SDL_ShaderCross_TranspileMSLFromSPIRV")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: SDL_ShaderCross_TranspileMSLFromSPIRV: " + err.Error())
	}
	_addr_ShaderCross_TranspileHLSLFromSPIRV, err = puregogen.OpenSymbol(_hnd_shadercross, "SDL_ShaderCross_TranspileHLSLFromSPIRV")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: SDL_ShaderCross_TranspileHLSLFromSPIRV: " + err.Error())
	}
	_addr_ShaderCross_CompileDXBCFromSPIRV, err = puregogen.OpenSymbol(_hnd_shadercross, "SDL_ShaderCross_CompileDXBCFromSPIRV")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: SDL_ShaderCross_CompileDXBCFromSPIRV: " + err.Error())
	}
	_addr_ShaderCross_CompileDXILFromSPIRV, err = puregogen.OpenSymbol(_hnd_shadercross, "SDL_ShaderCross_CompileDXILFromSPIRV")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: SDL_ShaderCross_CompileDXILFromSPIRV: " + err.Error())
	}
	_addr_ShaderCross_CompileGraphicsShaderFromSPIRV, err = puregogen.OpenSymbol(_hnd_shadercross, "SDL_ShaderCross_CompileGraphicsShaderFromSPIRV")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: SDL_ShaderCross_CompileGraphicsShaderFromSPIRV: " + err.Error())
	}
	_addr_ShaderCross_CompileComputePipelineFromSPIRV, err = puregogen.OpenSymbol(_hnd_shadercross, "SDL_ShaderCross_CompileComputePipelineFromSPIRV")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: SDL_ShaderCross_CompileComputePipelineFromSPIRV: " + err.Error())
	}
	_addr_ShaderCross_ReflectGraphicsSPIRV, err = puregogen.OpenSymbol(_hnd_shadercross, "SDL_ShaderCross_ReflectGraphicsSPIRV")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: SDL_ShaderCross_ReflectGraphicsSPIRV: " + err.Error())
	}
	_addr_ShaderCross_ReflectComputeSPIRV, err = puregogen.OpenSymbol(_hnd_shadercross, "SDL_ShaderCross_ReflectComputeSPIRV")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: SDL_ShaderCross_ReflectComputeSPIRV: " + err.Error())
	}
	_addr_ShaderCross_GetHLSLShaderFormats, err = puregogen.OpenSymbol(_hnd_shadercross, "SDL_ShaderCross_GetHLSLShaderFormats")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: SDL_ShaderCross_GetHLSLShaderFormats: " + err.Error())
	}
	_addr_ShaderCross_CompileDXBCFromHLSL, err = puregogen.OpenSymbol(_hnd_shadercross, "SDL_ShaderCross_CompileDXBCFromHLSL")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: SDL_ShaderCross_CompileDXBCFromHLSL: " + err.Error())
	}
	_addr_ShaderCross_CompileDXILFromHLSL, err = puregogen.OpenSymbol(_hnd_shadercross, "SDL_ShaderCross_CompileDXILFromHLSL")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: SDL_ShaderCross_CompileDXILFromHLSL: " + err.Error())
	}
	_addr_ShaderCross_CompileSPIRVFromHLSL, err = puregogen.OpenSymbol(_hnd_shadercross, "SDL_ShaderCross_CompileSPIRVFromHLSL")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: SDL_ShaderCross_CompileSPIRVFromHLSL: " + err.Error())
	}
}

func Init() error {
	r0, _, _ := purego.SyscallN(_addr_ShaderCross_Init)
	if r0 != 0 {
		return internal.LastErr()
	}

	return nil
}

func Quit() {
	purego.SyscallN(_addr_ShaderCross_Quit)
}

func GetSPIRVShaderFormats() (sdl.GPUShaderFormat, error) {
	r0, _, _ := purego.SyscallN(_addr_ShaderCross_GetSPIRVShaderFormats)
	if sdl.GPUShaderFormat(r0) == sdl.GPU_SHADERFORMAT_INVALID {
		return sdl.GPU_SHADERFORMAT_INVALID, internal.LastErr()
	}

	return sdl.GPUShaderFormat(r0), nil
}

func TranspileMSLFromSPIRV(info *SPIRVInfo) (string, error) {
	_info := info.as()
	r0, _, _ := purego.SyscallN(_addr_ShaderCross_TranspileMSLFromSPIRV, uintptr(unsafe.Pointer(_info)))
	if r0 == 0 {
		return "", internal.LastErr()
	}
	str := internal.ClonePtrString(r0)
	internal.Free(r0)
	runtime.KeepAlive(_info)

	return str, nil
}

func TranspileHLSLFromSPIRV(info *SPIRVInfo) (string, error) {
	_info := info.as()
	r0, _, _ := purego.SyscallN(_addr_ShaderCross_TranspileHLSLFromSPIRV, uintptr(unsafe.Pointer(_info)))
	if r0 == 0 {
		return "", internal.LastErr()
	}
	str := internal.ClonePtrString(r0)
	internal.Free(r0)
	runtime.KeepAlive(_info)

	return str, nil
}

func CompileDXBCFromSPIRV(info *SPIRVInfo) ([]byte, error) {
	var size int32

	_info := info.as()
	r0, _, _ := purego.SyscallN(_addr_ShaderCross_CompileDXBCFromSPIRV, uintptr(unsafe.Pointer(_info)), uintptr(unsafe.Pointer(&size)))
	if r0 == 0 {
		return nil, internal.LastErr()
	}
	code := internal.ClonePtrSlice[byte](r0, int(size))
	internal.Free(r0)
	runtime.KeepAlive(_info)

	return code, nil
}

func CompileDXILFromSPIRV(info *SPIRVInfo) ([]byte, error) {
	var size int32

	_info := info.as()
	r0, _, _ := purego.SyscallN(_addr_ShaderCross_CompileDXILFromSPIRV, uintptr(unsafe.Pointer(_info)), uintptr(unsafe.Pointer(&size)))
	if r0 == 0 {
		return nil, internal.LastErr()
	}
	code := internal.ClonePtrSlice[byte](r0, int(size))
	internal.Free(r0)
	runtime.KeepAlive(_info)

	return code, nil
}

func CompileGraphicsShaderFromSPIRV(device *sdl.GPUDevice, info *SPIRVInfo, resourceInfo *GraphicsShaderResourceInfo, props sdl.PropertiesID) (*sdl.GPUShader, error) {
	_info := info.as()
	r0, _, _ := purego.SyscallN(_addr_ShaderCross_CompileGraphicsShaderFromSPIRV, uintptr(unsafe.Pointer(device)), uintptr(unsafe.Pointer(_info)), uintptr(unsafe.Pointer(resourceInfo)), uintptr(props))
	if r0 == 0 {
		return nil, internal.LastErr()
	}

	runtime.KeepAlive(_info)
	runtime.KeepAlive(resourceInfo)

	return (*sdl.GPUShader)(unsafe.Pointer(r0)), nil
}

func CompileComputePipelineFromSPIRV(device *sdl.GPUDevice, info *SPIRVInfo, metadata *ComputePipelineMetadata, props sdl.PropertiesID) (*sdl.GPUComputePipeline, error) {
	_info := info.as()
	r0, _, _ := purego.SyscallN(_addr_ShaderCross_CompileComputePipelineFromSPIRV, uintptr(unsafe.Pointer(device)), uintptr(unsafe.Pointer(_info)), uintptr(unsafe.Pointer(metadata)), uintptr(props))
	if r0 == 0 {
		return nil, internal.LastErr()
	}

	runtime.KeepAlive(_info)
	runtime.KeepAlive(metadata)

	return (*sdl.GPUComputePipeline)(unsafe.Pointer(r0)), nil
}

func ReflectGraphicsSPIRV(bytecode []byte, props sdl.PropertiesID) (*GraphicsShaderMetadata, error) {
	_bytecode := unsafe.SliceData(bytecode)
	r0, _, _ := purego.SyscallN(_addr_ShaderCross_ReflectGraphicsSPIRV, uintptr(unsafe.Pointer(_bytecode)), uintptr(len(bytecode)), uintptr(props))
	if r0 == 0 {
		return nil, internal.LastErr()
	}

	ret := (*graphicsShaderMetadata)(unsafe.Pointer(r0))
	var info *GraphicsShaderResourceInfo
	if ret.ResourceInfo != nil {
		info = &GraphicsShaderResourceInfo{}
		*info = *ret.ResourceInfo
	}
	md := &GraphicsShaderMetadata{
		ResourceInfo: info,
		Inputs:       internal.ClonePtrSlice[IOVarMetadata](uintptr(unsafe.Pointer(ret.Inputs)), int(ret.NumInputs)),
		Outputs:      internal.ClonePtrSlice[IOVarMetadata](uintptr(unsafe.Pointer(ret.Outputs)), int(ret.NumOutputs)),
	}
	internal.Free(r0)

	runtime.KeepAlive(bytecode)

	return md, nil
}

func ReflectComputeSPIRV(bytecode []byte, props sdl.PropertiesID) (*ComputePipelineMetadata, error) {
	_bytecode := unsafe.SliceData(bytecode)
	r0, _, _ := purego.SyscallN(_addr_ShaderCross_ReflectComputeSPIRV, uintptr(unsafe.Pointer(_bytecode)), uintptr(len(bytecode)), uintptr(props))
	if r0 == 0 {
		return nil, internal.LastErr()
	}

	md := *(*ComputePipelineMetadata)(unsafe.Pointer(r0))
	internal.Free(r0)

	runtime.KeepAlive(bytecode)

	return &md, nil
}

func GetHLSLShaderFormats() (sdl.GPUShaderFormat, error) {
	r0, _, _ := purego.SyscallN(_addr_ShaderCross_GetHLSLShaderFormats)
	if sdl.GPUShaderFormat(r0) == sdl.GPU_SHADERFORMAT_INVALID {
		return sdl.GPU_SHADERFORMAT_INVALID, internal.LastErr()
	}

	return sdl.GPUShaderFormat(r0), nil
}

func CompileDXBCFromHLSL(info *HLSLInfo) ([]byte, error) {
	var size int32

	_info := info.as()
	r0, _, _ := purego.SyscallN(_addr_ShaderCross_CompileDXBCFromHLSL, uintptr(unsafe.Pointer(_info)), uintptr(unsafe.Pointer(&size)))
	if r0 == 0 {
		return nil, internal.LastErr()
	}
	code := internal.ClonePtrSlice[byte](r0, int(size))
	internal.Free(r0)
	runtime.KeepAlive(_info)

	return code, nil
}

func CompileDXILFromHLSL(info *HLSLInfo) ([]byte, error) {
	var size int32

	_info := info.as()
	r0, _, _ := purego.SyscallN(_addr_ShaderCross_CompileDXILFromHLSL, uintptr(unsafe.Pointer(_info)), uintptr(unsafe.Pointer(&size)))
	if r0 == 0 {
		return nil, internal.LastErr()
	}
	code := internal.ClonePtrSlice[byte](r0, int(size))
	internal.Free(r0)
	runtime.KeepAlive(_info)

	return code, nil
}

func CompileSPIRVFromHLSL(info *HLSLInfo) ([]byte, error) {
	var size int32

	_info := info.as()
	r0, _, _ := purego.SyscallN(_addr_ShaderCross_CompileSPIRVFromHLSL, uintptr(unsafe.Pointer(_info)), uintptr(unsafe.Pointer(&size)))
	if r0 == 0 {
		return nil, internal.LastErr()
	}
	code := internal.ClonePtrSlice[byte](r0, int(size))
	internal.Free(r0)
	runtime.KeepAlive(_info)

	return code, nil
}
