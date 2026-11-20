FROM ubuntu:24.04
ENV DEBIAN_FRONTEND=noninteractive
SHELL ["/bin/bash", "-c"]

ARG EMSDK_VERSION=7.0.0
ARG SDL_VERSION=3.4.0
ARG TTF_VERSION=3.2.2
ARG MIXER_VERSION=3.2.0
ARG IMAGE_VERSION=3.2.6

# Install dependencies
RUN apt-get update -y && \
        apt-get install -y --no-install-recommends ca-certificates git cmake make python3 xz-utils && \
        rm -rf /var/lib/apt/lists/*

# Clone and install emscripten
RUN git clone https://github.com/emscripten-core/emsdk.git --depth=1 && \
        cd emsdk && \
        ./emsdk install ${EMSDK_VERSION} && \
        ./emsdk activate ${EMSDK_VERSION}

# Clone SDL and satellite libraries
RUN git clone --depth 1 --branch release-${SDL_VERSION} https://github.com/libsdl-org/SDL.git
RUN git clone --depth 1 --recurse-submodules --shallow-submodules --branch release-${TTF_VERSION} https://github.com/libsdl-org/SDL_ttf.git
RUN git clone --depth 1 --recurse-submodules --shallow-submodules --branch release-${IMAGE_VERSION} https://github.com/libsdl-org/SDL_image.git
RUN git clone --depth 1 --recurse-submodules --shallow-submodules --branch release-${MIXER_VERSION} https://github.com/libsdl-org/SDL_mixer.git

# Emscripten's clock_gettime rejects CLOCK_MONOTONIC_RAW: SDL would fall back
# to gettimeofday and round every timestamp to 1 ms. Use CLOCK_MONOTONIC.
RUN sed -i 's/defined(CLOCK_MONOTONIC_RAW) && !defined(__ANDROID__)/& \&\& !defined(__EMSCRIPTEN__)/' SDL/src/timer/unix/SDL_systimer.c && \
        grep -q '!defined(__ANDROID__) && !defined(__EMSCRIPTEN__)' SDL/src/timer/unix/SDL_systimer.c

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
        emcmake cmake -DSDLIMAGE_AVIF=OFF -DSDLIMAGE_WEBP=OFF -DSDLIMAGE_TIF=OFF -DSDLIMAGE_BUILD_SHARED_LIBS=OFF -DSDLIMAGE_TESTS=OFF -DHAVE_POSIX_MEMALIGN=OFF -DSDLIMAGE_SAMPLES=OFF -DSDLIMAGE_VENDORED=ON -DSDL3_DIR=/build_sdl -DCMAKE_BUILD_TYPE=Release ../SDL_image && \
        emmake make -j4

# Build SDL_mixer
RUN source emsdk/emsdk_env.sh && \
        mkdir build_mixer && cd build_mixer && \
        emcmake cmake -DSDLMIXER_VORBIS_VORBISFILE=OFF -DSDLMIXER_FLAC_LIBFLAC=OFF -DSDLMIXER_MP3_MPG123=OFF -DSDLMIXER_MOD_XMP=OFF -DSDLMIXER_MOD=OFF -DSDLMIXER_WAVPACK=OFF -DSDLMIXER_GME=OFF -DSDLMIXER_OPUS=OFF -DSDLMIXER_SAMPLES=OFF -DSDLMIXER_VENDORED=ON -DSDL3_DIR=/build_sdl -DCMAKE_BUILD_TYPE=Release ../SDL_mixer && \
        emmake make -j4

# Export the public API listed in each library's linker version script,
# plus the C and emscripten functions the Go side calls directly.
RUN cat SDL/src/dynapi/SDL_dynapi.sym SDL_ttf/src/SDL_ttf.sym SDL_image/src/SDL_image.sym SDL_mixer/src/SDL_mixer.sym | \
        grep -oE '^\s+(SDL|TTF|IMG|MIX)_[A-Za-z0-9_]+;' | tr -d ' \t;' | \
        grep -vx 'SDL_DYNAPI_entry' | sed 's/^/_/' > exports.txt && \
        printf '%s\n' _malloc _free _emscripten_set_main_loop _emscripten_cancel_main_loop >> exports.txt

# Build final wasm library.
# The Go side reads every runtime method below through Module, see
# internal/memory_js.go.
# AUTO_JS_LIBRARIES=0 drops emscripten's legacy SDL 1 JS emulation: its
# SDL_Init, SDL_GetError, TTF_Init... would otherwise replace the exported
# SDL3 ones. The -l*.js libraries are the ones SDL3 needs instead.
COPY post.js /post.js
RUN source emsdk/emsdk_env.sh && \
        emcc -O2 --no-entry --post-js /post.js \
        -sAUTO_JS_LIBRARIES=0 -lwebgl.js -lhtml5_webgl.js -lasync.js \
        -sALLOW_MEMORY_GROWTH -sALLOW_TABLE_GROWTH \
        -sEXPORTED_FUNCTIONS=@exports.txt \
        -sEXPORTED_RUNTIME_METHODS=addFunction,stackAlloc,stackSave,stackRestore,stringToUTF8OnStack,UTF8ToString,getValue,setValue,HEAPU8,FS \
        build_image/libSDL3_image.a build_ttf/libSDL3_ttf.a build_mixer/libSDL3_mixer.a build_sdl/libSDL3.a -o sdl.js

RUN mkdir _out && \
        mv /build_sdl _out/build_sdl && \
        mv /build_ttf _out/build_ttf && \
        mv /build_image _out/build_image && \
        mv /build_mixer _out/build_mixer && \
        mv sdl* _out
