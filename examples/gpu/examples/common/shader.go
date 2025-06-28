package common

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/Zyko0/go-sdl3/sdl"
)

func LoadShader(
	device *sdl.GPUDevice,
	shaderFilename string,
	samplerCount uint32,
	uniformBufferCount uint32,
	storageBufferCount uint32,
	storageTextureCount uint32,
) (*sdl.GPUShader, error) {
	var stage sdl.GPUShaderStage
	if strings.Contains(shaderFilename, ".vert") {
		stage = sdl.GPU_SHADERSTAGE_VERTEX
	} else if strings.Contains(shaderFilename, ".frag") {
		stage = sdl.GPU_SHADERSTAGE_FRAGMENT
	} else {
		return nil, errors.New("invalid shader stage")
	}

	fullPath := ""
	backendFormats := device.ShaderFormats()
	format := sdl.GPU_SHADERFORMAT_INVALID
	entrypoint := ""

	if backendFormats&sdl.GPU_SHADERFORMAT_SPIRV == sdl.GPU_SHADERFORMAT_SPIRV {
		fullPath = fmt.Sprintf("content/shaders/compiled/spirv/%s.spv", shaderFilename)
		format = sdl.GPU_SHADERFORMAT_SPIRV
		entrypoint = "main"
	} else if backendFormats&sdl.GPU_SHADERFORMAT_MSL == sdl.GPU_SHADERFORMAT_MSL {
		fullPath = fmt.Sprintf("content/shaders/compiled/spirv/%s.msl", shaderFilename)
		format = sdl.GPU_SHADERFORMAT_MSL
		entrypoint = "main0"
	} else if backendFormats&sdl.GPU_SHADERFORMAT_DXIL == sdl.GPU_SHADERFORMAT_DXIL {
		fullPath = fmt.Sprintf("content/shaders/compiled/spirv/%s.dxil", shaderFilename)
		format = sdl.GPU_SHADERFORMAT_DXIL
		entrypoint = "main"
	} else {
		return nil, errors.New("unrecognized backend shader format")
	}

	fullPath = filepath.Join(BasePath, fullPath)

	code, err := os.ReadFile(fullPath)
	if err != nil {
		return nil, errors.New("failed to open shader: " + err.Error())
	}

	shaderInfo := sdl.GPUShaderCreateInfo{
		Code:               code,
		CodeSize:           uint64(len(code)),
		Entrypoint:         entrypoint,
		Format:             format,
		Stage:              stage,
		NumSamplers:        samplerCount,
		NumUniformBuffers:  uniformBufferCount,
		NumStorageBuffers:  storageBufferCount,
		NumStorageTextures: storageTextureCount,
	}

	shader, err := device.CreateGPUShader(&shaderInfo)
	if err != nil {
		return nil, errors.New("failed to create shader: " + err.Error())
	}

	return shader, nil
}
