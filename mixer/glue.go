package mixer

import (
	internal "github.com/Zyko0/go-sdl3/internal"
	"github.com/Zyko0/go-sdl3/sdl"
)

// Types

type Pointer = internal.Pointer

type MIX_InitFlags uint32

type InitFlags MIX_InitFlags

// Mix_Chunk - The internal format for an audio chunk
// (https://wiki.libsdl.org/SDL3_mixer/Mix_Chunk)
type Chunk struct {
	Allocated int32
	abuf      Pointer
	alen      uint32
	Volume    uint8
}

type TrackStoppedCallback uintptr

type TrackMixCallback uintptr

type GroupMixCallback uintptr

type PostMixCallback uintptr

// Custom

type AudioParams struct {
	Frequency int32
	Format    sdl.AudioFormat
	Channels  int32
}
