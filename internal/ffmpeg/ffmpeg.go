package ffmpeg

import (
	"encoding/json"
	"os/exec"
)

// GetInfo runs ffprobe and return the raw output
func GetInfo(filepath string) (*ProbeResult, error) {
	cmd := exec.Command("ffprobe", "-v", "quiet", "-print_format", "json", "-show_format", "-show_streams", filepath)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}

	var probeData ProbeResult
	if err := json.Unmarshal(output, &probeData); err != nil {
		return nil, err
	}

	return &probeData, nil
}