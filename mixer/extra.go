package mixer

// Types

type MIX_InitFlags uint32

type (
	MixCallback             func()
	MusicFinishedCallback   func()
	ChannelFinishedCallback func()
	EffectFunc_t            func()
	EffectDone_t            func()
	EachSoundFontCallback   func()
)

// Constants (#define)

const (
	MIX_INIT_FLAC    MIX_InitFlags = 0x00000001
	MIX_INIT_MOD     MIX_InitFlags = 0x00000002
	MIX_INIT_MP3     MIX_InitFlags = 0x00000008
	MIX_INIT_OGG     MIX_InitFlags = 0x00000010
	MIX_INIT_MID     MIX_InitFlags = 0x00000020
	MIX_INIT_OPUS    MIX_InitFlags = 0x00000040
	MIX_INIT_WAVPACK MIX_InitFlags = 0x00000080
)
