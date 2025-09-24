/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/jerkeyray/dory/internal/ffmpeg"
	"github.com/spf13/cobra"
)

// extractAudioCmd represents the extractAudio command
var extractAudioCmd = &cobra.Command{
	Use:   "extract-audio [input_file] [output_audio.mp3]",
	Short: "Extracts the audio track from a video file",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		inputFile := args[0]
		outputFile := args[1]

		fmt.Println("Extracting audio...")

		ffmpegArgs := ffmpeg.BuildExtractAudioArgs(inputFile, outputFile)
		if err := ffmpeg.Run(ffmpegArgs); err != nil {
			fmt.Printf("Error extracting audio: %v\n", err)
			return
		}

		fmt.Printf("Audio extracted successfully: %s\n", outputFile)
	},
}

func init() {
	rootCmd.AddCommand(extractAudioCmd)
}
