package ffmpeg

import "os/exec"

// GetInfo runs ffprobe and return the raw output
func GetInfo(filepath string) ([]byte, error) {
	cmd := exec.Command("ffprobe", "-v", "quiet", "-print_format", "json", "-show_format", "-show_streams", filepath)
	return cmd.CombinedOutput()
}