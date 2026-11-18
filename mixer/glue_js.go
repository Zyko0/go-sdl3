//go:build js

package mixer

import (
	"syscall/js"

	"github.com/Zyko0/go-sdl3/internal"
	"github.com/Zyko0/go-sdl3/sdl"
)

// Callbacks

func NewTrackStoppedCallback(fn func(track *Track)) TrackStoppedCallback {
	return TrackStoppedCallback(internal.NewCallback("vpp", func(args []js.Value) any {
		fn(internal.NewObject[Track](args[1]))
		return nil
	}))
}

func NewTrackMixCallback(fn func(track *Track, spec *sdl.AudioSpec, pcm []float32)) TrackMixCallback {
	return TrackMixCallback(newMixCallback(func(owner js.Value, spec *sdl.AudioSpec, pcm []float32) {
		fn(internal.NewObject[Track](owner), spec, pcm)
	}))
}

func NewGroupMixCallback(fn func(group *Group, spec *sdl.AudioSpec, pcm []float32)) GroupMixCallback {
	return GroupMixCallback(newMixCallback(func(owner js.Value, spec *sdl.AudioSpec, pcm []float32) {
		fn(internal.NewObject[Group](owner), spec, pcm)
	}))
}

func NewPostMixCallback(fn func(mixer *Mixer, spec *sdl.AudioSpec, pcm []float32)) PostMixCallback {
	return PostMixCallback(newMixCallback(func(owner js.Value, spec *sdl.AudioSpec, pcm []float32) {
		fn(internal.NewObject[Mixer](owner), spec, pcm)
	}))
}

// newMixCallback registers a (userdata, owner, spec, pcm, samples) mix
// callback. The pcm buffer is a copy written back after fn returns: Go
// cannot alias heap memory.
func newMixCallback(fn func(owner js.Value, spec *sdl.AudioSpec, pcm []float32)) uintptr {
	return internal.NewCallback("vppppi", func(args []js.Value) any {
		var spec sdl.AudioSpec
		internal.CopyJSToObject(&spec, args[2])
		pcm := internal.GetNumericSliceFromJSPtr[float32](args[3], args[4].Int())
		fn(args[1], &spec, pcm)
		internal.SetNumericSliceToJSPtr(args[3], pcm)
		return nil
	})
}
