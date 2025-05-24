FROM ubuntu:20.04
ENV DEBIAN_FRONTEND=noninteractive
SHELL ["/bin/bash", "-c"]

# Install dependencies
RUN apt update -y
RUN apt-get install -y wget
RUN wget -qO- https://apt.llvm.org/llvm-snapshot.gpg.key | tee /etc/apt/trusted.gpg.d/apt.llvm.org.asc
RUN apt-get install -y software-properties-common
RUN apt-get install -y git
RUN apt-get install -y cmake
RUN add-apt-repository -s -y "deb http://apt.llvm.org/focal/ llvm-toolchain-focal-18 main"
RUN apt-get install -y llvm-18 clang-18 libclang-18-dev libclang-cpp18-dev python3-pip

# Clone and install emscripten
RUN git clone https://github.com/emscripten-core/emsdk.git --depth=1 && \
        cd emsdk && \
        ./emsdk install tot && \
        ./emsdk activate tot

ENV SDL3_VERSION=3.2.8
ENV TTF_VERSION=3.2.0
ENV IMAGE_VERSION=3.2.4
ENV MIXER_VERSION=main

# Clone SDL and satellite libraries
RUN git clone --depth 1 --branch release-$SDL3_VERSION https://github.com/libsdl-org/SDL.git
RUN git clone --recurse-submodules --branch release-$TTF_VERSION https://github.com/libsdl-org/SDL_ttf.git
RUN git clone --recurse-submodules --branch release-$IMAGE_VERSION https://github.com/libsdl-org/SDL_image.git
RUN git clone --recurse-submodules https://github.com/libsdl-org/SDL_mixer.git

# Build SDL
RUN source emsdk/emsdk_env.sh && \
        mkdir build_sdl && cd build_sdl && \
        emcmake cmake -DCMAKE_BUILD_TYPE=Release -DCMAKE_DISABLE_PRECOMPILE_HEADERS=ON ../SDL && \
        emmake make -j4

# Build SDL_ttf
RUN source emsdk/emsdk_env.sh && \
        mkdir build_ttf && cd build_ttf && \
        emcmake cmake -DSDLTTF_SAMPLES=OFF -DSDLTTF_VENDORED=ON -DSDL3_DIR=/build_sdl -DCMAKE_BUILD_TYPE=Release ../SDL_ttf && \
        emmake make -j4

# Build SDL_image
RUN source emsdk/emsdk_env.sh && \
        mkdir build_image && cd build_image && \
        emcmake cmake -DSDLIMAGE_AVIF=OFF -DSDLIMAGE_WEBP=OFF -DSDLIMAGE_TIF=OFF -DSDLIMAGE_BUILD_SHARED_LIBS=OFF -DSDLIMAGE_TESTS=OFF -DHAVE_POSIX_MEMALIGN=OFF -DDAV1D_ASM=OFF -DAOM_TARGET_CPU=generic -DSDLIMAGE_SAMPLES=OFF -DSDLIMAGE_VENDORED=ON -DSDL_SHARED=OFF -DSDL3_DIR=/build_sdl -DCMAKE_BUILD_TYPE=Release ../SDL_image && \
        emmake make -j4

# Build SDL_mixer
RUN source emsdk/emsdk_env.sh && \
        mkdir build_mixer && cd build_mixer && \
        emcmake cmake -DSDLMIXER_MOD_XMP=OFF -DSDLMIXER_MOD=OFF -DSDLMIXER_WAVPACK=OFF -DSDLMIXER_GME=OFF -DSDLMIXER_OPUS=OFF -DSDLMIXER_SAMPLES=OFF -DSDLMIXER_VENDORED=ON -DSDL_SHARED=OFF -DSDL3_DIR=/build_sdl -DCMAKE_BUILD_TYPE=Release ../SDL_mixer && \
        emmake make -j4

# Build final wasm library
RUN source emsdk/emsdk_env.sh && \
        emcc --no-entry -s LINKABLE=1 -s EXPORTED_RUNTIME_METHODS="['addFunction', 'allocateUTF8','UTF8ToString', 'stringToUTF8', 'getValue', 'setValue']" -s ALLOW_MEMORY_GROWTH=1 -s ALLOW_TABLE_GROWTH=1 -s EXPORT_ALL -s DEFAULT_LIBRARY_FUNCS_TO_INCLUDE='$addFunction' build_image/libSDL3_image.a build_ttf/libSDL3_ttf.a build_mixer/libSDL3_mixer.a build_sdl/libSDL3.a -o sdl.html

RUN mkdir _out && \
        mv /build_sdl _out/build_sdl && \
        mv /build_ttf _out/build_ttf && \
        mv /build_image _out/build_image && \
        mv /build_mixer _out/build_mixer && \
        mv sdl* _out
