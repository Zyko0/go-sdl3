//go:build !js

package mixer

import (
	"unsafe"

	"github.com/Zyko0/go-sdl3/sdl"
	purego "github.com/ebitengine/purego"
)

// Callbacks

func NewTrackStoppedCallback(fn func(track *Track)) TrackStoppedCallback {
	return TrackStoppedCallback(purego.NewCallback(func(_ uintptr, track *Track) uintptr {
		fn(track)
		return 0
	}))
}

func NewTrackMixCallback(fn func(track *Track, spec *sdl.AudioSpec, pcm []float32)) TrackMixCallback {
	return TrackMixCallback(purego.NewCallback(func(_ uintptr, track *Track, spec *sdl.AudioSpec, pcm *float32, samples int32) uintptr {
		fn(track, spec, unsafe.Slice(pcm, samples))
		return 0
	}))
}

func NewGroupMixCallback(fn func(group *Group, spec *sdl.AudioSpec, pcm []float32)) GroupMixCallback {
	return GroupMixCallback(purego.NewCallback(func(_ uintptr, group *Group, spec *sdl.AudioSpec, pcm *float32, samples int32) uintptr {
		fn(group, spec, unsafe.Slice(pcm, samples))
		return 0
	}))
}

func NewPostMixCallback(fn func(mixer *Mixer, spec *sdl.AudioSpec, pcm []float32)) PostMixCallback {
	return PostMixCallback(purego.NewCallback(func(_ uintptr, mixer *Mixer, spec *sdl.AudioSpec, pcm *float32, samples int32) uintptr {
		fn(mixer, spec, unsafe.Slice(pcm, samples))
		return 0
	}))
}
