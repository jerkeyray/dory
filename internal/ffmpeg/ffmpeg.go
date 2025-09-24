package ffmpeg

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/schollz/progressbar/v3"
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

// RunWithProgress executes an ffmpeg command and shows a progress bar
// totalDuration is the total duration of the video in seconds needed to set the progress bar max
func RunWithProgress(info *ProbeResult, args []string) error {
    // Get the total duration from the info struct.
    totalDuration, err := strconv.ParseFloat(info.Format.Duration, 64)
    if err != nil {
        return fmt.Errorf("could not parse video duration: %w", err)
    }

    // build the ffmpeg command
    cmd := exec.Command("ffmpeg", args...)

    // Create a pipe to read FFmpeg's stdout
    stdout, err := cmd.StdoutPipe()
    if err != nil {
        return fmt.Errorf("failed to create stdout pipe: %w", err)
    }
    cmd.Stderr = io.Discard // Redirect stderr to /dev/null

    // Start the command
    if err := cmd.Start(); err != nil {
        return fmt.Errorf("failed to start ffmpeg: %w", err)
    }

    // Initialize progress bar
    // Creates a progress bar that goes from 0 to totalDuration.
    // Customizes how it looks ([=====> ]).
    // Adds a description (Compressing...)
    bar := progressbar.NewOptions(int(totalDuration),
        progressbar.OptionSetDescription("Compressing..."),
        progressbar.OptionSetTheme(progressbar.Theme{
            Saucer:        "█",
            SaucerHead:    "█",
            SaucerPadding: "░",
            BarStart:      "|",
            BarEnd:        "|",
        }))

    // Read progress from stdout in a separate goroutine
    go parseProgress(stdout, bar)

    // Wait for the command to finish
    if err := cmd.Wait(); err != nil {
        return fmt.Errorf("ffmpeg command failed: %w", err)
    }

    return nil
}

// launch ffmpeg and read its live progress output
// everytime we see a line with out_time_ms=XXXXXX we update the progress bar
// keep updating till the duration is reached
// exit when ffmpeg exits

// parseProgress reads ffmpeg output line by line and updates the progress bar
func parseProgress(stdout io.Reader, bar *progressbar.ProgressBar) {
    scanner := bufio.NewScanner(stdout)
    for scanner.Scan() {
        line := scanner.Text()
        if strings.Contains(line, "out_time_ms=") {
            timeStr := strings.Split(line, "=")[1]
            if ms, err := strconv.ParseInt(timeStr, 10, 64); err == nil {
                seconds := ms / 1_000_000 // Convert microseconds to seconds
                bar.Set(int(seconds))
            }
        }
    }
}