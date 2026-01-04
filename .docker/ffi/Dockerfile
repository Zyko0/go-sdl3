FROM ubuntu:20.04
ENV DEBIAN_FRONTEND=noninteractive

ARG SDL_VERSION=3.4.0
ARG TTF_VERSION=3.2.2
ARG MIXER_VERSION=main
ARG IMAGE_VERSION=3.2.6

RUN apt update -y

# Apparently wget is not installed by default
RUN apt-get install -y wget
# Add llvm pub key
RUN wget -qO- https://apt.llvm.org/llvm-snapshot.gpg.key | tee /etc/apt/trusted.gpg.d/apt.llvm.org.asc
# Install add-apt-repository
RUN apt-get install -y software-properties-common
# Add llvm 18 repository with sources
RUN add-apt-repository -s -y "deb http://apt.llvm.org/focal/ llvm-toolchain-focal-18 main"
RUN echo apt-get update && apt-get install -y llvm-18 clang-18 libclang-18-dev libclang-cpp18-dev python3-pip
# Easiest way to get sufficiently new cmake appears to be using pip - Ubuntu 20.10 version is too old
RUN pip3 install --upgrade cmake

# Clone c2ffi repository
RUN apt-get install -y git
RUN git clone https://github.com/rpav/c2ffi.git --depth 1
RUN cd c2ffi && git checkout llvm-18.1.0

# Build c2ffi
RUN cd c2ffi && mkdir -p build && cd build && \
        cmake -DBUILD_CONFIG=Release .. && make

# Clone SDL repository
RUN git clone https://github.com/libsdl-org/SDL.git && \
        cd SDL && git checkout tags/release-${SDL_VERSION} && \
        cp -r include/SDL3 /usr/include/SDL3

# Generate SDL ffi
RUN echo "#include <SDL3/SDL.h>\n#include <SDL3/SDL_main.h>\n#include <SDL3/SDL_vulkan.h>" > sdl.c
RUN /c2ffi/build/bin/c2ffi sdl.c > sdl_ffi.json

# Clone SDL_ttf repository
RUN git clone https://github.com/libsdl-org/SDL_ttf.git && \
        cd SDL_ttf && git checkout tags/release-${TTF_VERSION} && \
        cp -r include/SDL3_ttf /usr/include/SDL3_ttf

# Generate SDL_ttf ffi
RUN echo "#include <SDL3_ttf/SDL_ttf.h>\n#include <SDL3_ttf/SDL_textengine.h>" > ttf.c
RUN /c2ffi/build/bin/c2ffi ttf.c > ttf_ffi.json

# Clone SDL_mixer repository
RUN git clone https://github.com/libsdl-org/SDL_mixer.git && \
        cd SDL_mixer && \
        cp -r include/SDL3_mixer/ /usr/include/SDL3_mixer/

# Generate SDL_mixer ffi
RUN echo "#include <SDL3_mixer/SDL_mixer.h>\n" > mixer.c
RUN /c2ffi/build/bin/c2ffi mixer.c > mixer_ffi.json

# Clone SDL_image repository
RUN git clone https://github.com/libsdl-org/SDL_image.git && \
        cd SDL_image && git checkout tags/release-${IMAGE_VERSION} && \
        cp -r include/SDL3_image/ /usr/include/SDL3_image/

# Generate SDL_image ffi
RUN echo "#include <SDL3_image/SDL_image.h>\n" > image.c
RUN /c2ffi/build/bin/c2ffi image.c > image_ffi.json