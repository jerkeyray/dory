 ____   ___  ______   __
|  _ \ / _ \|  _ \ \ / /
| | | | | | | |_) \ V / 
| |_| | |_| |  _ < | |  
|____/ \___/|_| \_\|_|  

A fast, user-friendly command-line video processing tool built with Go. Dory provides a simple interface to common video operations using FFmpeg under the hood, featuring progress bars and optimized performance for everyday video editing tasks.

## Features

- **Video Compression**: Reduce file sizes with configurable quality profiles
- **Video Trimming**: Extract specific segments from videos with frame-accurate precision
- **Video Concatenation**: Join multiple video files into a single output
- **GIF Creation**: Generate high-quality GIFs from video clips with custom parameters
- **Audio Extraction**: Extract audio tracks from video files
- **Media Information**: Display detailed metadata about video files
- **Fast Performance**: Leverages stream copying where possible to avoid re-encoding

## Prerequisites

Dory requires FFmpeg to be installed on your system:

### macOS

```bash
brew install ffmpeg
```

### Ubuntu/Debian

```bash
sudo apt update
sudo apt install ffmpeg
```

### Windows

Download from [ffmpeg.org](https://ffmpeg.org/download.html) or use:

```powershell
winget install ffmpeg
```

## Installation

### From Source

```bash
git clone https://github.com/jerkeyray/dory.git
cd dory
go build -o dory .
```

### Binary Release

Download the latest binary from the [releases page](https://github.com/jerkeyray/dory/releases).

## Usage

### Compress Videos

Reduce video file size with quality profiles:

```bash
# Compress to 720p (default)
dory compress input.mp4 output.mp4

# Compress to specific resolution
dory compress input.mp4 output.mp4 --profile 1080p
dory compress input.mp4 output.mp4 --profile 480p
```

Available profiles: `1080p`, `720p`, `480p`

### Trim Videos

Extract segments from videos:

```bash
dory trim input.mp4 clip.mp4 --start 00:01:30 --end 00:02:45
```

Uses stream copying for fast, lossless trimming.

### Concatenate Videos

Join multiple videos into one:

```bash
dory concat output.mp4 video1.mp4 video2.mp4 video3.mp4
```

### Create GIFs

Generate high-quality GIFs from video clips:

```bash
# Basic GIF creation
dory gif input.mp4 output.gif --start 00:00:10 --duration 3

# Custom parameters
dory gif input.mp4 output.gif --start 00:00:10 --duration 5 --width 600 --fps 20
```

Options:

- `--start`: Start time (default: 00:00:00)
- `--duration`: Duration in seconds (default: 3)
- `--width`: Width in pixels (default: 500)
- `--fps`: Frames per second (default: 15)

### Extract Audio

Extract audio tracks from videos:

```bash
dory extract-audio input.mp4 audio.mp3
```

### Get Media Information

Display file metadata:

```bash
dory info input.mp4
```

Shows duration, resolution, video codec, and audio codec information.

## Technical Details

### Architecture

- **Language**: Go 1.24.1
- **CLI Framework**: Cobra for command-line interface
- **Progress Display**: Custom progress bars using schollz/progressbar
- **Video Processing**: FFmpeg integration with real-time progress parsing

### Performance Optimizations

- Stream copying for trim operations (no re-encoding)
- Palette-based GIF generation for optimal quality and size
- Real-time progress tracking via FFmpeg's progress output
- Concurrent progress parsing for smooth user experience

### Project Structure

```
dory/
├── main.go                 # Application entry point
├── cmd/                    # Command implementations
│   ├── root.go            # Root command and CLI setup
│   ├── compress.go        # Video compression
│   ├── trim.go            # Video trimming
│   ├── concat.go          # Video concatenation
│   ├── gif.go             # GIF creation
│   ├── extractAudio.go    # Audio extraction
│   └── info.go            # Media information
└── internal/ffmpeg/       # FFmpeg integration layer
    ├── ffmpeg.go          # Core execution and progress handling
    ├── builders.go        # Command argument builders
    └── types.go           # Data structures for FFprobe output
```

## License

This project is licensed under the MIT License. See the LICENSE file for details.
