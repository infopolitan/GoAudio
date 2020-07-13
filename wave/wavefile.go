package wave

// representation of the wave file, used by reader.go and writer.go

type Sample float64

/*

╔════════╤════════════════╤══════╤═══════════════════════════════════════════════════╗
║ Offset │ Field          │ Size │ -- start of header                                ║
╠════════╪════════════════╪══════╪═══════════════════════════════════════════════════╣
║ 0      │ ChunkID        │ 4    │                                                   ║
╟────────┼────────────────┼──────┼───────────────────────────────────────────────────╢
║ 4      │ ChunkSize      │ 4    │                                                   ║
╟────────┼────────────────┼──────┼───────────────────────────────────────────────────╢
║ 8      │ Format         │ 8    │                                                   ║
╟────────┼────────────────┼──────┼───────────────────────────────────────────────────╢
║ --     │ --             │ --   │ -- start of fmt                                   ║
╟────────┼────────────────┼──────┼───────────────────────────────────────────────────╢
║ 12     │ SubchunkID     │ 4    │                                                   ║
╟────────┼────────────────┼──────┼───────────────────────────────────────────────────╢
║ 16     │ SubchunkSize   │ 4    │                                                   ║
╟────────┼────────────────┼──────┼───────────────────────────────────────────────────╢
║ 20     │ AudioFormat    │ 2    │                                                   ║
╟────────┼────────────────┼──────┼───────────────────────────────────────────────────╢
║ 22     │ NumChannels    │ 2    │                                                   ║
╟────────┼────────────────┼──────┼───────────────────────────────────────────────────╢
║ 24     │ SampleRate     │ 4    │                                                   ║
╟────────┼────────────────┼──────┼───────────────────────────────────────────────────╢
║ 28     │ ByteRate       │ 4    │                                                   ║
╟────────┼────────────────┼──────┼───────────────────────────────────────────────────╢
║ 32     │ BlockAlign     │ 2    │                                                   ║
╟────────┼────────────────┼──────┼───────────────────────────────────────────────────╢
║ 34     │ BitsPerSample  │ 2    │                                                   ║
╟────────┼────────────────┼──────┼───────────────────────────────────────────────────╢
║ * 36   │ ExtraParamSize │ 2    │ Optional! Only when not PCM                       ║
╟────────┼────────────────┼──────┼───────────────────────────────────────────────────╢
║ * 38   │ ExtraParams    │ *    │ Optional! Only when not PCM                       ║
╟────────┼────────────────┼──────┼───────────────────────────────────────────────────╢
║ --     │ --             │ --   │ -- start of data, assuming PCM                    ║
╟────────┼────────────────┼──────┼───────────────────────────────────────────────────╢
║ 36     │ Subchunk2ID    │ 4    │ (offset by extra params of subchunk 1 if not PCM) ║
╟────────┼────────────────┼──────┼───────────────────────────────────────────────────╢
║ 40     │ SubchunkSize   │ 4    │ (offset by extra params of subchunk 1 if not PCM) ║
╟────────┼────────────────┼──────┼───────────────────────────────────────────────────╢
║ 44     │ Data           │ *    │ (offset by extra params of subchunk 1 if not PCM) ║
╚════════╧════════════════╧══════╧═══════════════════════════════════════════════════╝


*/

// Wave represents an entire .wav audio file
type Wave struct {
	WaveHeader
	WaveFmt
	WaveData
}

// WaveHeader describes the header each WAVE file should start with
type WaveHeader struct {
	ChunkID   []byte // should be RIFF on little-endian or RIFX on big-endian systems..
	ChunkSize int
	Format    string // sanity-check, should be WAVE (//TODO: keep bytes??)
}

// WaveFmt describes the format of the sound-information in the data subchunks
type WaveFmt struct {
	Subchunk1ID    []byte // should contain "fmt"
	Subchunk1Size  int    // 16 for PCM
	AudioFormat    int    // PCM = 1 (Linear Quantization), if not 1, compression was used.
	NumChannels    int    // Mono 1, Stereo = 2, ..
	SampleRate     int    // 44100 for CD-Quality, etc..
	ByteRate       int    // SampleRate * NumChannels * BitsPerSample / 8
	BlockAlign     int    // NumChannels * BitsPerSample / 8 (number of bytes per sample)
	BitsPerSample  int    // 8 bits = 8, 16 bits = 16, .. :-)
	ExtraParamSize int    // if not PCM, can contain extra params
	ExtraParams    []byte // the actual extra params.
}

// waveData contains the raw sound data
type WaveData struct {
	Subchunk2ID   []byte // Identifier of subchunk
	Subchunk2Size int    // size of raw sound data
	RawData       []byte // raw sound data itself
	Samples       []Sample
}
