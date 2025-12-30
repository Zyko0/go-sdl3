# shadercross

These are bindings to the [SDL_shadercross](https://github.com/libsdl-org/SDL_shadercross) library. It allows you to compile HLSL code to shader formats of different platforms (DXBC, DXIL, SPIRV, MSL), allowing you to target DirectX, Metal, Vulkan with a single shader language (HLSL).

You can statically compile your HLSL code to all formats when releasing your game, using the `shadercross` binary. But doing so, every time you make a change to your HLSL code might slow you down during development.

So this is where using it as a library can make sense, as it is giving you much more runtime control, allowing features such as hot-reloading shader while your program is running.

## Requirements

For these bindings to work, you will need a few dependencies that differ based on your platform.

You can either follow the build/installation guidelines from https://github.com/libsdl-org/SDL_shadercross or you can directly go to their [github actions](https://github.com/libsdl-org/SDL_shadercross/actions) and download the artifacts of your system, containing all the required libraries.

## Usage

```go
package main

import (
	"fmt"
	"log"

	"github.com/Zyko0/go-sdl3/bin/binsdl"
	"github.com/Zyko0/go-sdl3/shadercross"
)

var (
	shaderSrc = `
cbuffer UBO : register(b0, space3)
{
    int mode : packoffset(c0);
};

Texture2D<float4> Texture : register(t0, space2);

float4 main(float2 TexCoord : TEXCOORD0) : SV_Target0
{
    float w, h;
    Texture.GetDimensions(w, h);
    int2 texelPos = int2(float2(w, h) * TexCoord);
    float4 mainTexel = Texture[texelPos];
    if (mode == 0)
    {
        return mainTexel;
    }
    else
    {
        float4 bottomTexel = Texture[texelPos + int2(0, 1)];
        float4 leftTexel = Texture[texelPos + int2(-1, 0)];
        float4 topTexel = Texture[texelPos + int2(0, -1)];
        float4 rightTexel = Texture[texelPos + int2(1, 0)];
        return ((((mainTexel * 0.2f) + (bottomTexel * 0.2f)) + (leftTexel * 0.20000000298023223876953125f)) + (topTexel * 0.20000000298023223876953125f)) + (rightTexel * 0.2f);
    }
}
`
)

func main() {
	defer binsdl.Load().Unload()

	err := shadercross.LoadLibrary(shadercross.Path())
	if err != nil {
		panic("couldn't load shadercross library: " + err.Error())
	}
	defer shadercross.CloseLibrary()

	spirv, err := shadercross.CompileSPIRVFromHLSL(&shadercross.HLSLInfo{
		Source:      shaderSrc,
		Entrypoint:  "main",
		ShaderStage: shadercross.SHADERSTAGE_FRAGMENT,
	})
	if err != nil {
		log.Fatal("couldn't get HLSL shader formats:", err)
	}
	fmt.Println("SPIR-V bytes count:", len(spirv))

	spirvInfo := &shadercross.SPIRVInfo{
		Bytecode:    spirv,
		Entrypoint:  "main",
		ShaderStage: shadercross.SHADERSTAGE_FRAGMENT,
	}

	msl, err := shadercross.TranspileMSLFromSPIRV(spirvInfo)
	if err != nil {
		log.Fatal("couldn't transpile MSL from SPIR-V:", err)
	}
	fmt.Println("MSL from SPIR-V:", string(msl))

	hlsl, err := shadercross.TranspileHLSLFromSPIRV(spirvInfo)
	if err != nil {
		log.Fatal("couldn't transpile HLSL from SPIR-V:", err)
	}
	fmt.Println("HLSL from SPIR-V:", string(hlsl))
}
```

<ins>**Outputs:**</ins>
<details>
<summary>MSL (from SPIR-V)</summary>

```c
#include <metal_stdlib>
#include <simd/simd.h>

using namespace metal;

struct type_UBO
{
    int mode;
};

struct main0_out
{
    float4 out_var_SV_Target0 [[color(0)]];
};

struct main0_in
{
    float2 in_var_TEXCOORD0 [[user(locn0)]];
};

fragment main0_out main0(main0_in in [[stage_in]], constant type_UBO& UBO [[buffer(0)]], texture2d<float> Texture [[texture(0)]])
{
    main0_out out = {};
    float4 _79;
    do
    {
        uint2 _37 = uint2(Texture.get_width(), Texture.get_height());
        int2 _44 = int2(float2(float(_37.x), float(_37.y)) * in.in_var_TEXCOORD0);
        float4 _47 = Texture.read(uint2(uint2(_44)), 0u);
        if (UBO.mode == 0)
        {
            _79 = _47;
            break;
        }
        else
        {
            _79 = ((((_47 * 0.20000000298023223876953125) + (Texture.read(uint2(uint2(_44 + int2(0, 1))), 0u) * 0.20000000298023223876953125)) + (Texture.read(uint2(uint2(_44 + int2(-1, 0))), 0u) * 0.20000000298023223876953125)) + (Texture.read(uint2(uint2(_44 + int2(0, -1))), 0u) * 0.20000000298023223876953125)) + (Texture.read(uint2(uint2(_44 + int2(1, 0))), 0u) * 0.20000000298023223876953125);
            break;
        }
        break; // unreachable workaround
    } while(false);
    out.out_var_SV_Target0 = _79;
    return out;
}
```
</details>

<details>
<summary>HLSL (from SPIR-V)</summary>

```c
cbuffer type_UBO : register(b0, space3)
{
    int UBO_mode : packoffset(c0);
};

Texture2D<float4> Texture : register(t0, space2);

static float2 in_var_TEXCOORD0;
static float4 out_var_SV_Target0;

struct SPIRV_Cross_Input
{
    float2 in_var_TEXCOORD0 : TEXCOORD0;
};

struct SPIRV_Cross_Output
{
    float4 out_var_SV_Target0 : SV_Target0;
};

uint2 spvTextureSize(Texture2D<float4> Tex, uint Level, out uint Param)
{
    uint2 ret;
    Tex.GetDimensions(Level, ret.x, ret.y, Param);
    return ret;
}

void main_inner()
{
    float4 _79;
    do
    {
        uint _37_dummy_parameter;
        uint2 _37 = spvTextureSize(Texture, uint(0), _37_dummy_parameter);
        int2 _44 = int2(float2(float(_37.x), float(_37.y)) * in_var_TEXCOORD0);
        float4 _47 = Texture.Load(int3(uint2(_44), 0u));
        if (UBO_mode == 0)
        {
            _79 = _47;
            break;
        }
        else
        {
            _79 = ((((_47 * 0.20000000298023223876953125f) + (Texture.Load(int3(uint2(_44 + int2(0, 1)), 0u)) * 0.20000000298023223876953125f)) + (Texture.Load(int3(uint2(_44 + int2(-1, 0)), 0u)) * 0.20000000298023223876953125f)) + (Texture.Load(int3(uint2(_44 + int2(0, -1)), 0u)) * 0.20000000298023223876953125f)) + (Texture.Load(int3(uint2(_44 + int2(1, 0)), 0u)) * 0.20000000298023223876953125f);
            break;
        }
        break; // unreachable workaround
    } while(false);
    out_var_SV_Target0 = _79;
}

SPIRV_Cross_Output main(SPIRV_Cross_Input stage_input)
{
    in_var_TEXCOORD0 = stage_input.in_var_TEXCOORD0;
    main_inner();
    SPIRV_Cross_Output stage_output;
    stage_output.out_var_SV_Target0 = out_var_SV_Target0;
    return stage_output;
}
```
</details>