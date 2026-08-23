package main

// midiAPIRef stands in for the quick reference page every other library has.
// SDL_native_midi publishes no wiki page, and its header carries no doc
// comments, so these descriptions are the only source for the ones emitted
// above each binding. Names and parameter types could be recovered from the
// ffi entries; the prose could not, which is why this lives here rather than
// beside the generated csv files.
//
// The return value matches what download produces: the body of a ```c fence,
// one prototype per line, description after "//". Every parameter needs a
// name, or paramTypes skips it.
func midiAPIRef() string {
	return `
// SDL_native_midi.h
bool NativeMidi_Init(void);  // Initialize the native MIDI subsystem.
void NativeMidi_Quit(void);  // Shut down the native MIDI subsystem.
NativeMidi_Song * NativeMidi_LoadSong_IO(SDL_IOStream *src, bool closeio);  // Load a MIDI song from an SDL_IOStream.
NativeMidi_Song * NativeMidi_LoadSong(const char *path);  // Load a MIDI song from a file path.
void NativeMidi_DestroySong(NativeMidi_Song *song);  // Free a song and the resources it holds.
void NativeMidi_Start(NativeMidi_Song *song, int loops);  // Start playing a song, replacing the one currently playing.
void NativeMidi_Pause(void);  // Pause playback of the current song.
void NativeMidi_Resume(void);  // Resume playback of a paused song.
void NativeMidi_Stop(void);  // Stop playback of the current song.
bool NativeMidi_Active(void);  // Query whether a song is currently playing.
void NativeMidi_SetVolume(float volume);  // Set the volume applied to MIDI playback.
`
}
