/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/jerkeyray/dory/internal/ffmpeg"
	"github.com/spf13/cobra"
)

// these variables will hold the values of the flags
var startTime string
var endTime string

// trimCmd represents the trim command
var trimCmd = &cobra.Command{
	Use:   "trim [input_file] [output_file]",
	Short: "Trims a video to the specified start and end times",
	Long: `Trims a video to a new, shorter clip.
This operation is very fast as it avoids re-encoding the video.
Example:
dory trim original.mp4 clip.mp4 --start 00:01:30 --end 00:02:00`,
	Args: cobra.ExactArgs(2), // expects exactly 2 arguments: input and output file paths
	Run: func(cmd *cobra.Command, args []string) {
		inputFile := args[0]
		outputFile := args[1]

		// construct the trim options from the arguements
		opts := ffmpeg.TrimOptions{
			InputPath:  inputFile,
			OutputPath: outputFile,
			StartTime:  startTime,
			EndTime:    endTime,
		}

		// build the command arguments
		ffmpegArgs := ffmpeg.BuildTrimArgs(opts)

		fmt.Printf("Trimming video from %s to %s...\n", startTime, endTime)

		// execute the ffmpeg command
		err := ffmpeg.Run(ffmpegArgs)
		if err != nil {
			fmt.Printf("Error trimming video: %v\n", err)
			return
		}

		fmt.Printf("Video trimmed successfully!: %s\n", outputFile)
	},
}

func init() {
	rootCmd.AddCommand(trimCmd)

	// The "P" in StringVarP adds a shorthand flag (-s).
	trimCmd.Flags().StringVarP(&startTime, "start", "s", "", "Start time for the trim (e.g., 00:01:30)")
	trimCmd.Flags().StringVarP(&endTime, "end", "e", "", "End time for the trim (e.g., 00:02:00)")

	// We make the flags required. If the user doesn't provide them, Cobra will show an error.
	trimCmd.MarkFlagRequired("start")
	trimCmd.MarkFlagRequired("end")
}
