package mixer

// NumChannelsPaused - Get the number of channels currently paused.
func NumChannelsPaused() int32 {
	return iPaused(-1)
}

// NumChannelsPlaying - Get the number of channels currently playing audio.
func NumChannelsPlaying() int32 {
	return iPlaying(-1)
}
