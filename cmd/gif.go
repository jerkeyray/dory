/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/jerkeyray/dory/internal/ffmpeg"
	"github.com/spf13/cobra"
)

var gifStartTime string
var gifDuration string
var gifWidth int
var gifFps int

// gifCmd represents the gif command
var gifCmd = &cobra.Command{
	Use:   "gif [input_file] [output_file.gif]",
	Short: "Creates a high-quality GIF from a video clip",
	Long: `Creates a high-quality GIF from a video clip.
Example:
dory gif input.mp4 output.gif --start 00:00:10 --duration 3 --width 500`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		inputFile := args[0]
		outputFile := args[1]
		paletteFile := "palette.png" // temporary palette file

		// genereate the options for palette generation
		opts := ffmpeg.GifOptions{
			InputPath:   inputFile,
			OutputPath:  outputFile,
			PalettePath: paletteFile,
			StartTime:   gifStartTime,
			Duration:    gifDuration,
			Fps:         gifFps,
			Width:       gifWidth,
		}

		// build the command to generate the palette
		fmt.Println("Generating optimal color palette...")
		paletteArgs := ffmpeg.BuildPaletteArgs(opts)
		if err := ffmpeg.Run(paletteArgs); err != nil {
			fmt.Printf("Error generating palette: %v\n", err)
			return
		}

		// create the GIF using the generated palette
		fmt.Println("Creating GIF...")
		gifArgs := ffmpeg.BuildGifArgs(opts)
		if err := ffmpeg.Run(gifArgs); err != nil {
			fmt.Printf("Error creating GIF: %v\n", err)
			return
		}

		os.Remove(paletteFile)

		fmt.Printf("\nGIF created successfully: %s\n", outputFile)
	},
}

func init() {
	rootCmd.AddCommand(gifCmd)

	gifCmd.Flags().StringVarP(&gifStartTime, "start", "s", "00:00:00", "Start time for the GIF (e.g., 00:01:30)")
	gifCmd.Flags().StringVarP(&gifDuration, "duration", "d", "3", "Duration of the GIF in seconds")
	gifCmd.Flags().IntVar(&gifFps, "fps", 15, "Frames per second for the GIF")
	gifCmd.Flags().IntVar(&gifWidth, "width", 500, "Width of the GIF in pixels")
}
