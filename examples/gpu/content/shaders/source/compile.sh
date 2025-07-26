# Requires shadercross CLI installed from SDL_shadercross
for filename in *.vert.hlsl; do
    if [ -f "$filename" ]; then
        shadercross "$filename" -o "../compiled/spirv/${filename/.hlsl/.spv}"
        shadercross "$filename" -o "../compiled/msl/${filename/.hlsl/.msl}"
        shadercross "$filename" -o "../compiled/dxil/${filename/.hlsl/.dxil}"
    fi
done

for filename in *.frag.hlsl; do
    if [ -f "$filename" ]; then
        shadercross "$filename" -o "../compiled/spirv/${filename/.hlsl/.spv}"
        shadercross "$filename" -o "../compiled/msl/${filename/.hlsl/.msl}"
        shadercross "$filename" -o "../compiled/dxil/${filename/.hlsl/.dxil}"
    fi
done

for filename in *.comp.hlsl; do
    if [ -f "$filename" ]; then
        shadercross "$filename" -o "../compiled/spirv/${filename/.hlsl/.spv}"
        shadercross "$filename" -o "../compiled/msl/${filename/.hlsl/.msl}"
        shadercross "$filename" -o "../compiled/dxil/${filename/.hlsl/.dxil}"
    fi
done
