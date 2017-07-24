package main

import (
	"fmt"
	"image/gif"
	"image/jpeg"
	"os"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	gifname := kingpin.Arg("input", "The input GIF file").Required().String()
	savefolder := kingpin.Flag("output", "Output folder name").Default("out").Short('o').String()
	shouldsplit := kingpin.Flag("split", "Split the input gif into individual frames?").Short('s').Bool()
	savejpeg := kingpin.Flag("jpeg", "Save individual frames as JPEGs instead of GIFs").Short('j').Bool()
	kingpin.CommandLine.HelpFlag.Short('h')
	kingpin.Parse()

	f, err := os.Open(*gifname)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		fmt.Fprintf(os.Stderr, "Could not open image \"%s\"! Exiting.", *gifname)
		os.Exit(1)
	}
	defer f.Close()

	// Decode our single gif into a struct containing every frame & some other useful information
	split, err := gif.DecodeAll(f)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		fmt.Fprintf(os.Stderr, "Unable to decode \"%s\" as a GIF! Exiting.", *gifname)
		os.Exit(1)
	}

	// If the desired output folder doesn't exist we'll make it
	if _, err := os.Stat(*savefolder); os.IsNotExist(err) {
		os.Mkdir(*savefolder, 0644)
	}

	printStats(split, *savefolder, *gifname, *shouldsplit)

	if *shouldsplit {
		SplitGIF(split, *savefolder, *savejpeg)
	}
}

// Prints out a range of interesting stats about the input gif
func printStats(gifstruct *gif.GIF, path string, name string, shouldsplit bool) {
	fmt.Printf("--- GIF STATS: %s ---\n", name)
	fmt.Printf("Number of frames:\n %d\n", len(gifstruct.Image))
	fmt.Printf("Delay per frame (100ths / sec):\n %d\n", gifstruct.Delay)
	fmt.Printf("Loop count:\n %d\n", gifstruct.LoopCount)
	fmt.Printf("Image size (height x width):\n %d x %d\n", gifstruct.Config.Height, gifstruct.Config.Width)
	if shouldsplit {
		fmt.Printf("\nOutputting individual frames to folder: %s\n", path)
	}
}

// SplitGIF will split a gif (gifstruct) into it's individual frames then save them into a folder (savefolder), optionally as JPEGs (saveasjpeg)
func SplitGIF(gifstruct *gif.GIF, savefolder string, saveasjpeg bool) {
	var format string
	if saveasjpeg {
		format = ".jpeg"
	} else {
		format = ".gif"
	}
	for i, img := range gifstruct.Image {
		filename := fmt.Sprintf("%s/frame%02d%s", savefolder, i, format)
		outfile, err := os.Create(filename)
		defer outfile.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error! Could create image: %s", filename)
		} else {
			if saveasjpeg {
				jpeg.Encode(outfile, img, &jpeg.Options{Quality: 100})
			} else {
				gif.Encode(outfile, img, &gif.Options{NumColors: 256})
			}
		}
	}
}
