package ffmpeg

// ProbeResults represents the top-level structure of ffmprobe's JSON output
type ProbeResult struct {
	Format Format   `json:"format"`
	Stream []Stream `json:"streams"`
}

// Format hold information about the container format
type Format struct {
	Filename string `json:"filename"`
	Duration string `json:"duration"`
	Size     string `json:"size"`
}

// Stream holds information about each individual video or audio stream
type Stream struct {
	CodecName string `json:"codec_name"`
	CodecType string `json:"codec_type"`
	Width     int    `json:"width"`
	Height    int    `json:"height"`
}
