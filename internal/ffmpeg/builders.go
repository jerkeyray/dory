package ffmpeg

// TrimOptions holds options for trimming a video
type TrimOptions struct {
	InputPath  string // Path to the input video file
	OutputPath string // Path to the output trimmed video file
	StartTime  string // Start time for the trim (e.g., "00:01:30")
	EndTime    string // End time for the trim (e.g., "00:02:00")
}

// BuildTrimArgs constructs the ffmpeg command arguments for trimming a video
func BuildTrimArgs(opts TrimOptions) []string {
	// -c copy tells FFmpeg to copy the video and audio streams without re-encoding.
	// This is extremely fast and preserves the original quality.
	return []string{
		"-i", opts.InputPath,
		"-ss", opts.StartTime,
		"-to", opts.EndTime,
		"-c", "copy",
		"-y", // Overwrite output file if it exists
		opts.OutputPath,
	}
}
