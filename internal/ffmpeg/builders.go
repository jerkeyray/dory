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
	InputPath	string // Path to the input video file
	OutputPath	string // Path to the output compressed video file
	Profile		string // Compression profile (e.g., "1080p", "720p", "480p")
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
