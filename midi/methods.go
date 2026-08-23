package midi

// NativeMidi_DestroySong - Free a song and the resources it holds.
func (song *Song) Destroy() {
	iDestroySong(song)
}

// NativeMidi_Start - Start playing a song, replacing the one currently playing.
//
// loops counts repeats rather than plays: 0 plays the song once, a positive n
// replays it n more times, and any negative value loops forever.
func (song *Song) Start(loops int32) {
	iStart(song, loops)
}
