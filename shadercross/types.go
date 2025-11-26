package shadercross

import (
	"unsafe"

	"github.com/Zyko0/go-sdl3/internal"
	"github.com/Zyko0/go-sdl3/sdl"
)

type IOVarType int32

const (
	IOVAR_TYPE_UNKNOWN IOVarType = iota
	IOVAR_TYPE_INT8
	IOVAR_TYPE_UINT8
	IOVAR_TYPE_INT16
	IOVAR_TYPE_UINT16
	IOVAR_TYPE_INT32
	IOVAR_TYPE_UINT32
	IOVAR_TYPE_INT64
	IOVAR_TYPE_UINT64
	IOVAR_TYPE_FLOAT16
	IOVAR_TYPE_FLOAT32
	IOVAR_TYPE_FLOAT64
)

type ShaderStage int32

const (
	SHADERSTAGE_VERTEX ShaderStage = iota
	SHADERSTAGE_FRAGMENT
	SHADERSTAGE_COMPUTE
)

type IOVarMetadata struct {
	Name       string    /**< The UTF-8 name of the variable. */
	Location   uint32    /**< The location of the variable. */
	VectorType IOVarType /**< The vector type of the variable. */
	VectorSize uint32    /**< The number of components in the vector type of the variable. */
}

type ioVarMetadata struct {
	Name       *byte
	Location   uint32
	VectorType IOVarType
	VectorSize uint32
}

func (v *IOVarMetadata) as() *ioVarMetadata {
	if v == nil {
		return nil
	}
	return &ioVarMetadata{
		Name:       internal.StringToPtr(v.Name),
		Location:   v.Location,
		VectorType: v.VectorType,
		VectorSize: v.VectorSize,
	}
}

type GraphicsShaderResourceInfo struct {
	NumSamplers        uint32 /**< The number of samplers defined in the shader. */
	NumStorageTextures uint32 /**< The number of storage textures defined in the shader. */
	NumStorageBuffers  uint32 /**< The number of storage buffers defined in the shader. */
	NumUniformBuffers  uint32 /**< The number of uniform buffers defined in the shader. */
}

type GraphicsShaderMetadata struct {
	ResourceInfo *GraphicsShaderResourceInfo /**< Sub-struct containing the resource info of the shader. */
	Inputs       []IOVarMetadata             /**< The inputs defined in the shader. */
	Outputs      []IOVarMetadata             /**< The outputs defined in the shader. */
}

type graphicsShaderMetadata struct {
	ResourceInfo *GraphicsShaderResourceInfo
	NumInputs    uint32
	Inputs       *ioVarMetadata
	NumOutputs   uint32
	Outputs      *ioVarMetadata
}

func (m *GraphicsShaderMetadata) as() *graphicsShaderMetadata {
	if m == nil {
		return nil
	}

	inputs := make([]ioVarMetadata, len(m.Inputs))
	for i := range m.Inputs {
		inputs[i] = *m.Inputs[i].as()
	}
	outputs := make([]ioVarMetadata, len(m.Outputs))
	for i := range m.Outputs {
		outputs[i] = *m.Outputs[i].as()
	}

	return &graphicsShaderMetadata{
		ResourceInfo: m.ResourceInfo,
		NumInputs:    uint32(len(inputs)),
		Inputs:       unsafe.SliceData(inputs),
		NumOutputs:   uint32(len(outputs)),
		Outputs:      unsafe.SliceData(outputs),
	}
}

type ComputePipelineMetadata struct {
	NumSamplers                 uint32 /**< The number of samplers defined in the shader. */
	NumReadonlyStorageTextures  uint32 /**< The number of readonly storage textures defined in the shader. */
	NumReadonlyStorageBuffers   uint32 /**< The number of readonly storage buffers defined in the shader. */
	NumReadwriteStorageTextures uint32 /**< The number of read-write storage textures defined in the shader. */
	NumReadwriteStorageBuffers  uint32 /**< The number of read-write storage buffers defined in the shader. */
	NumUniformBuffers           uint32 /**< The number of uniform buffers defined in the shader. */
	ThreadcountX                uint32 /**< The number of threads in the X dimension. */
	ThreadcountY                uint32 /**< The number of threads in the Y dimension. */
	ThreadcountZ                uint32 /**< The number of threads in the Z dimension. */
}

type SPIRVInfo struct {
	Bytecode    []byte      /**< The SPIRV bytecode. */
	Entrypoint  string      /**< The entry point function name for the shader in UTF-8. */
	ShaderStage ShaderStage /**< The shader stage to transpile the shader with. */

	Props sdl.PropertiesID /**< A properties ID for extensions. Should be 0 if no extensions are needed. */
}

type spirvInfo struct {
	Bytecode     *byte
	BytecodeSize uint32
	Entrypoint   *byte
	ShaderStage  ShaderStage

	Props sdl.PropertiesID
}

func (v *SPIRVInfo) as() *spirvInfo {
	if v == nil {
		return nil
	}
	return &spirvInfo{
		Bytecode:     unsafe.SliceData(v.Bytecode),
		BytecodeSize: uint32(len(v.Bytecode)),
		Entrypoint:   internal.StringToPtr(v.Entrypoint),
		ShaderStage:  v.ShaderStage,
		Props:        v.Props,
	}
}

const (
	PROP_SHADER_DEBUG_ENABLE_BOOLEAN         = "SDL_shadercross.spirv.debug.enable"
	PROP_SHADER_DEBUG_NAME_STRING            = "SDL_shadercross.spirv.debug.name"
	PROP_SHADER_CULL_UNUSED_BINDINGS_BOOLEAN = "SDL_shadercross.spirv.cull_unused_bindings"
	PROP_SPIRV_PSSL_COMPATIBILITY_BOOLEAN    = "SDL_shadercross.spirv.pssl.compatibility"
	PROP_SPIRV_MSL_VERSION_STRING            = "SDL_shadercross.spirv.msl.version"
)

type HLSLDefine struct {
	Name  string /**< The define name. */
	Value string /**< An optional value for the define. Can be NULL. */
}

type hlslDefine struct {
	Name  *byte
	Value *byte
}

func (d *HLSLDefine) as() *hlslDefine {
	if d == nil {
		return nil
	}
	return &hlslDefine{
		Name:  internal.StringToPtr(d.Name),
		Value: internal.StringToNullablePtr(d.Value),
	}
}

type HLSLInfo struct {
	Source      string       /**< The HLSL source code for the shader. */
	Entrypoint  string       /**< The entry point function name for the shader in UTF-8. */
	IncludeDir  string       /**< The include directory for shader code. Optional, can be NULL. */
	Defines     []HLSLDefine /**< An array of defines. Optional, can be NULL. If not NULL, must be terminated with a fully NULL define struct. */
	ShaderStage ShaderStage  /**< The shader stage to compile the shader with. */

	Props sdl.PropertiesID /**< A properties ID for extensions. Should be 0 if no extensions are needed. */
}

type hlslInfo struct {
	Source      *byte
	Entrypoint  *byte
	IncludeDir  *byte
	Defines     *hlslDefine
	ShaderStage ShaderStage

	Props sdl.PropertiesID
}

func (v *HLSLInfo) as() *hlslInfo {
	if v == nil {
		return nil
	}

	var definesPtr *hlslDefine
	if len(v.Defines) > 0 {
		defines := make([]hlslDefine, 0, len(v.Defines))
		for i := range v.Defines {
			def := v.Defines[i].as()
			if def != nil {
				defines = append(defines, *def)
			}
		}
		defines = append(defines, hlslDefine{})
		definesPtr = unsafe.SliceData(defines)
	}

	return &hlslInfo{
		Source:      internal.StringToPtr(v.Source),
		Entrypoint:  internal.StringToPtr(v.Entrypoint),
		IncludeDir:  internal.StringToNullablePtr(v.IncludeDir),
		Defines:     definesPtr,
		ShaderStage: v.ShaderStage,
		Props:       v.Props,
	}
}
