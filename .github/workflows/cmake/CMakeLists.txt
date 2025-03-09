cmake_minimum_required(VERSION 3.16)

# set the output directory for built objects.
# This makes sure that the dynamic library goes into the build directory automatically.
set(CMAKE_RUNTIME_OUTPUT_DIRECTORY "${CMAKE_BINARY_DIR}/$<CONFIGURATION>")
set(CMAKE_LIBRARY_OUTPUT_DIRECTORY "${CMAKE_BINARY_DIR}/$<CONFIGURATION>")

# prevent installing to system directories. 
set(CMAKE_INSTALL_PREFIX "${CMAKE_BINARY_DIR}" CACHE INTERNAL "")

# Declare the project
project(sdl-web)

set(BUILD_SHARED_LIBS OFF CACHE INTERNAL "")
set(SDL_SHARED OFF)

# Set the name of the executable
set(EXECUTABLE_NAME ${PROJECT_NAME})

# Create an executable or a shared library based on the platform and add our sources to it
add_executable(${EXECUTABLE_NAME})

# Add your sources to the target
target_sources(${EXECUTABLE_NAME} 
PRIVATE 
    sdl.c
)

# on Web targets, we need CMake to generate a HTML webpage. 
set(CMAKE_EXECUTABLE_SUFFIX ".html" CACHE INTERNAL "")

# Configure SDL by calling its CMake file.
# we use EXCLUDE_FROM_ALL so that its install targets and configs don't
# pollute upwards into our configuration.
add_subdirectory(SDL EXCLUDE_FROM_ALL)

set(SDLTTF_VENDORED ON) # tell SDL_ttf to build its own dependencies
add_subdirectory(SDL_ttf EXCLUDE_FROM_ALL)	

# SDL_mixer (used for playing audio)
set(SDLMIXER_MIDI_NATIVE ON)    
set(SDLMIXER_GME ON)
set(SDLMIXER_WAVPACK ON)     
set(SDLMIXER_MOD ON)
set(SDLMIXER_OPUS ON)
set(SDLMIXER_VENDORED ON)   # tell SDL_mixer to build its own dependencies
add_subdirectory(SDL_mixer EXCLUDE_FROM_ALL)

# SDL_image (used for loading various image formats)
set(SDLIMAGE_VENDORED ON)
set(SDLIMAGE_AVIF ON)
set(SDLIMAGE_BMP ON)
set(SDLIMAGE_JPEG ON)
set(SDLIMAGE_WEBP ON)

set(DAV1D_ASM OFF)
set(AOM_TARGET_CPU "generic")
add_subdirectory(SDL_image EXCLUDE_FROM_ALL)

# Link SDL to our executable. This also makes its include directory available to us. 
target_link_libraries(${EXECUTABLE_NAME} PUBLIC 
	SDL3_ttf::SDL3_ttf     
	SDL3_mixer::SDL3_mixer 
	SDL3_image::SDL3_image
    SDL3::SDL3              
)

# target_compile_definitions(${EXECUTABLE_NAME} PUBLIC SDL_MAIN_USE_CALLBACKS)