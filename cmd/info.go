/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/jerkeyray/dory/internal/ffmpeg"
	"github.com/spf13/cobra"
)

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Gives information about a media file",
	Long: `Provides a summary of key information about a media file, including:
- Duration
- Resolution
- Video Codec
- Audio Codec

Example usage:
	dory info <filepath>`,
	Args: cobra.ExactArgs(1), // expects exactly 1 argument: the input file path

	Run: func(cmd *cobra.Command, args []string) {
		// check if filepath argument is provided
		if len(args) == 0 {
			fmt.Println("Error: please provide a filepath.")
			return
		}

		// filepath is the first argument and get media info using ffprobe
		filepath := args[0]
		info, err := ffmpeg.GetInfo(filepath)
		if err != nil {
			fmt.Printf("Error running ffprobe: %v\n", err)
			return
		}

		// find the video and audio streams
		var videoStream ffmpeg.Stream
		var audioStream ffmpeg.Stream

		// Iterate through the streams to find video and audio
		for _, stream := range info.Stream {
			if stream.CodecType == "video" {
				videoStream = stream
			} else if stream.CodecType == "audio" {
				audioStream = stream
			}
		}

		// Print the clean report
		fmt.Println("--- Media Information ---")
		// Extract just the filename from the full path for cleaner output
		fileName := info.Format.Filename[strings.LastIndex(info.Format.Filename, "/")+1:]
		fmt.Printf("File:        %s\n", fileName)
		fmt.Printf("Duration:    %.2fs\n", parseDuration(info.Format.Duration))
		fmt.Printf("Resolution:  %dx%d\n", videoStream.Width, videoStream.Height)
		fmt.Printf("Video Codec: %s\n", videoStream.CodecName)
		fmt.Printf("Audio Codec: %s\n", audioStream.CodecName)
	},
}

func parseDuration(s string) float64 {
	var f float64
	fmt.Sscanf(s, "%f", &f)
	return f
}

func init() {
	rootCmd.AddCommand(infoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// infoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// infoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
