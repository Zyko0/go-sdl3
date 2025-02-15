FROM ubuntu:20.04
ENV DEBIAN_FRONTEND=noninteractive

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
RUN git clone https://github.com/rpav/c2ffi.git
RUN cd c2ffi && git checkout llvm-18.1.0

# Build c2ffi
RUN cd c2ffi && mkdir -p build && cd build && \
        cmake -DBUILD_CONFIG=Release .. && make

# Clone sdl repository
RUN git clone https://github.com/libsdl-org/SDL.git && \
        cd SDL && git checkout tags/release-3.2.2 && \
        cp -r include/SDL3 /usr/include/SDL3

# Generate SDL ffi
RUN echo "#include <SDL3/SDL.h>\n#include <SDL3/SDL_main.h>" > sdl.c
RUN /c2ffi/build/bin/c2ffi sdl.c > sdl_ffi.json

# Clone sdl repository
RUN git clone https://github.com/libsdl-org/SDL_ttf.git && \
        cd SDL_ttf && git checkout tags/preview-3.1.0 && \
        cp -r include/SDL3_ttf /usr/include/SDL3_ttf

# Generate SDL ffi
RUN echo "#include <SDL3_ttf/SDL_ttf.h>\n#include <SDL3_ttf/SDL_textengine.h>" > ttf.c
RUN /c2ffi/build/bin/c2ffi ttf.c > ttf_ffi.json