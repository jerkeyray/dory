/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "dory",
	Short: "A fast, user-friendly command-line video processing tool",
	Long: `Dory is a command-line video processing tool built with Go that provides
a simple interface to common video operations using FFmpeg under the hood.

Available Commands:
  compress      Reduce video file sizes with quality profiles
  trim          Extract specific segments from videos  
  concat        Join multiple video files into one
  gif           Create animated GIFs from video clips
  extract-audio Extract audio tracks from videos
  info          Display video metadata and properties

Use Cases:
  • Compress large videos for sharing or storage
  • Create short clips from longer videos
  • Combine multiple video segments
  • Generate GIFs for social media or documentation
  • Extract music or audio from video files
  • Check video properties before processing

Examples:
  dory compress input.mp4 output.mp4 --profile 720p
  dory trim video.mp4 clip.mp4 --start 00:01:30 --end 00:02:45
  dory concat output.mp4 part1.mp4 part2.mp4 part3.mp4
  dory gif input.mp4 output.gif --start 00:00:10 --duration 3
  dory extract-audio video.mp4 audio.mp3
  dory info video.mp4`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
