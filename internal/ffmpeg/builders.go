package ffmpeg

import "fmt"

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

type CompressOptions struct {
	InputPath  string // Path to the input video file
	OutputPath string // Path to the output compressed video file
	Profile    string // Compression profile (e.g., "1080p", "720p", "480p")
}

func BuildCompressArgs(opts CompressOptions) ([]string, error) {
	// -c:v libx264: A very popular and compatible video codec.
	// -crf 23: A good quality setting (lower is better quality, 23 is a sane default).
	// -c:a aac: A standard audio codec.
	// -vf "scale=-2:720": A video filter to scale the height to 720 pixels while maintaining aspect ratio.

	args := []string{
		"-i", opts.InputPath,
		"-c:v", "libx264",
		"-crf", "23",
		"-c:a", "aac",
		"-y", // Overwrite output
	}

	switch opts.Profile {
	case "1080p":
		args = append(args, "-vf", "scale=-2:1080")
	case "720p":
		args = append(args, "-vf", "scale=-2:720")
	case "480p":
		args = append(args, "-vf", "scale=-2:480")
	default:
		return nil, fmt.Errorf("unknown profile: %s", opts.Profile)
	}

	args = append(args, opts.OutputPath)
	return args, nil
}

// GifOptions holds options for creating a GIF from a video
type GifOptions struct {
	InputPath   string
	OutputPath  string
	PalettePath string
	StartTime   string
	Duration    string
	Fps         int
	Width       int
}

// BuildPaletteArgs creates the command to generate a color palette.
func BuildPaletteArgs(opts GifOptions) []string {
	// -vf is a video filter. We combine three filters:
	// fps=...: set the frames per second.
	// scale=...:-1: scale the width, height will adjust automatically.
	// palettegen: generate the color palette.
	filter := fmt.Sprintf("fps=%d,scale=%d:-1:flags=lanczos,palettegen", opts.Fps, opts.Width)
	return []string{
		"-ss", opts.StartTime,
		"-t", opts.Duration,
		"-i", opts.InputPath,
		"-vf", filter,
		"-y", // Overwrite output
		opts.PalettePath,
	}
}

// BuildGifArgs creates the command to generate the final GIF using the palette.
func BuildGifArgs(opts GifOptions) []string {
	// -filter_complex is for complex filtergraphs with multiple inputs.
	// [0:v] is the video stream from the first input (input.mp4).
	// [1:v] is the video stream from the second input (palette.png).
	// paletteuse applies the palette to the video stream.
	return []string{
		"-ss", opts.StartTime,
		"-t", opts.Duration,
		"-i", opts.InputPath,
		"-i", opts.PalettePath,
		"-filter_complex", "[0:v][1:v]paletteuse",
		"-y", // Overwrite output
		opts.OutputPath,
	}
}

// BuildExtractAudioArgs creates the command to extract audio.
func BuildExtractAudioArgs(inputPath, outputPath string) []string {
	return []string{
		"-i", inputPath,
		"-vn",       // -vn tells FFmpeg to ignore the video stream.
		"-q:a", "0", // -q:a 0 is a high-quality setting for VBR MP3s.
		"-map", "a", // -map a ensures only the audio stream is processed.
		"-y", // Overwrite output
		outputPath,
	}
}
