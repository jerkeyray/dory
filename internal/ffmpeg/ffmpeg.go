package ffmpeg

import (
	"encoding/json"
	"fmt"
	"github.com/briandowns/spinner"
	"os/exec"
	"time"
)

// GetInfo runs ffprobe and return the JSON-parsed ProbeResult
func GetInfo(filepath string) (*ProbeResult, error) {
	// flag explanations:
	// -v quiet: suppress all output except for the JSON data
	// -print_format json: output in JSON format
	// -show_format: include container format information
	// -show_streams: include information about each stream (video, audio, etc.)
	cmd := exec.Command("ffprobe", "-v", "quiet", "-print_format", "json", "-show_format", "-show_streams", filepath)

	// CombinedOutput runs the command and returns its combined standard output and standard error
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}

	var probeData ProbeResult
	// Unmarshal the JSON output into the ProbeResult struct
	if err := json.Unmarshal(output, &probeData); err != nil {
		return nil, err
	}

	return &probeData, nil
}

// Run executes an ffmpeg command with a spinner
func Run(args []string) error {
	cmd := exec.Command("ffmpeg", args...)

	// Start a spinner to show the user something is happening.
	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	s.Start()

	// Run the command and capture any output (including errors).
	output, err := cmd.CombinedOutput()

	s.Stop()

	if err != nil {
		// If FFmpeg fails, we print its output to help debug.
		return fmt.Errorf("ffmpeg command failed: %s\nOutput:\n%s", err, string(output))
	}

	return nil
}
