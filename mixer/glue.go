package mixer

import (
	internal "github.com/Zyko0/go-sdl3/internal"
	"github.com/Zyko0/go-sdl3/sdl"
)

// Types

type Pointer = internal.Pointer

// MIX_TrackStoppedCallback - A callback that fires when a [MIX_Track](MIX_Track) is stopped.
// (https://wiki.libsdl.org/SDL3_mixer/MIX_TrackStoppedCallback)
type TrackStoppedCallback uintptr

// MIX_TrackMixCallback - A callback that fires when a [MIX_Track](MIX_Track) is mixing at various stages.
// (https://wiki.libsdl.org/SDL3_mixer/MIX_TrackMixCallback)
type TrackMixCallback uintptr

// MIX_GroupMixCallback - A callback that fires when a [MIX_Group](MIX_Group) has completed mixing.
// (https://wiki.libsdl.org/SDL3_mixer/MIX_GroupMixCallback)
type GroupMixCallback uintptr

// MIX_PostMixCallback - A callback that fires when all mixing has completed.
// (https://wiki.libsdl.org/SDL3_mixer/MIX_PostMixCallback)
type PostMixCallback uintptr

// Custom

type AudioParams struct {
	Frequency int32
	Format    sdl.AudioFormat
	Channels  int32
}
