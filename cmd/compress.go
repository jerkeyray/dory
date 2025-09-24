/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/jerkeyray/dory/internal/ffmpeg"
	"github.com/spf13/cobra"
)

var profile string

// compressCmd represents the compress command
var compressCmd = &cobra.Command{
	Use:   "compress [input_file] [output_file]",
	Short: "Compresses a video to a smaller file size",
	Long:  `Re-encodes a video to a smaller file size using predefined quality profiles.`,
	Args:  cobra.ExactArgs(2), // expects exactly 2 arguments: input and output file paths
	Run: func(cmd *cobra.Command, args []string) {
		inputFile := args[0]
		outputFile := args[1]

		// get video info to know the duration for the progress bar
		info, err := ffmpeg.GetInfo(inputFile)
		if err != nil {
			fmt.Printf("Error getting video info: %v\n", err)
			return
		}

		// construct the compress options from the arguements
		opts := ffmpeg.CompressOptions{
			InputPath:  inputFile,
			OutputPath: outputFile,
			Profile:    profile,
		}
		ffmpegArgs, err := ffmpeg.BuildCompressArgs(opts)
		if err != nil {
			fmt.Printf("Error building ffmpeg arguments: %v\n", err)
			return
		}

		// Add the crucial -progress flag
		// We send progress to stdout by using a pipe. "pipe:1" is stdout.
		finalArgs := append([]string{"-progress", "pipe:1"}, ffmpegArgs...)

		// 3. Execute with the new progress bar runner
		err = ffmpeg.RunWithProgress(info, finalArgs)
		if err != nil {
			fmt.Printf("Error compressing video: %v\n", err)
			return
		}

		fmt.Println("\nVideo compressed successfully!")

	},
}

func init() {
	rootCmd.AddCommand(compressCmd)

	compressCmd.Flags().StringVarP(&profile, "profile", "p", "720p", "Compression profile (e.g., 1080p, 720p, 480p)")
}
