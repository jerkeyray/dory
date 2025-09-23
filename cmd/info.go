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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Error: please provide a filepath.")
			return 
		}	

		filepath := args[0]
		info, err := ffmpeg.GetInfo(filepath)
		if err != nil {
			fmt.Printf("Error running ffprobe: %v\n", err)
			return 
		}

		// find the video and audio streams
		var videoStream ffmpeg.Stream
		var audioStream ffmpeg.Stream

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
