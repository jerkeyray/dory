/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/jerkeyray/dory/internal/ffmpeg"
	"github.com/spf13/cobra"
)

// concatCmd represents the concat command
var concatCmd = &cobra.Command{
	Use:   "concat [output_file] [input_file1] [input_file2] ...",
	Short: "Joins multiple video clips into a single file",
	Args:  cobra.MinimumNArgs(3), // We need at least 1 output and 2 inputs
	Run: func(cmd *cobra.Command, args []string) {
		outputFile := args[0]
		inputFiles := args[1:]
		tempFileName := "concat_list.txt"

		// create temp file to hold the list of files to concatenate
		tempFile, err := os.Create(tempFileName)
		if err != nil {
			fmt.Printf("Error creating temp file: %v\n", err)
			return
		}

		// Use defer to make sure the temp file is cleaned up, even if errors occur.
		defer os.Remove(tempFileName)

		for _, file := range inputFiles {
			// clean the path and write it to the file in the format FFmpeg requires.
			absPath, _ := filepath.Abs(file)
			tempFile.WriteString(fmt.Sprintf("file '%s'\n", absPath))
		}
		// close the file so FFmpeg can read it.
		tempFile.Close()

		// run the concat command
		fmt.Println("Joining videos...")
		ffmpegArgs := ffmpeg.BuildConcatArgs(tempFileName, outputFile)
		if err := ffmpeg.Run(ffmpegArgs); err != nil {
			fmt.Printf("Error joining videos: %v\n", err)
			return
		}

		fmt.Printf("Videos joined successfully: %s\n", outputFile)
	},
}

func init() {
	rootCmd.AddCommand(concatCmd)

}
