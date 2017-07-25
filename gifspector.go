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
	//shouldtrim := kingpin.Flag("trim", "Trim the gif (saved as 'trimmed.gif') using a start and end frame number (Passed as arguments, see --help)").Short('t').Bool()
	starttrim := kingpin.Arg("start", "Start frame to trim the gif from (Inclusive)").Default("-1").Int()
	endtrim := kingpin.Arg("end", "The frame to trim the gif up to (Exclusive)").Default("-1").Int()
	kingpin.CommandLine.HelpFlag.Short('h')
	kingpin.Parse()

	f, err := os.Open(*gifname)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		fmt.Fprintf(os.Stderr, "Could not open image \"%s\"! Exiting.\n", *gifname)
		os.Exit(1)
	}
	defer f.Close()

	// Decode our single gif into a struct containing every frame & some other useful information
	gifstruct, err := gif.DecodeAll(f)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		fmt.Fprintf(os.Stderr, "Unable to decode \"%s\" as a GIF! Exiting.\n", *gifname)
		os.Exit(1)
	}

	// If the desired output folder doesn't exist we'll make it
	if _, err := os.Stat(*savefolder); os.IsNotExist(err) {
		os.Mkdir(*savefolder, 0644)
	}

	printStats(gifstruct, *savefolder, *gifname, *shouldsplit)

	if *shouldsplit {
		SplitGIF(gifstruct, *savefolder, *savejpeg)
	}

	if *starttrim != -1 {
		start, end := normalizeTrimLength(*starttrim, *endtrim, gifstruct)
		ftrim, err := os.Create("trimmed.gif")
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			fmt.Fprintf(os.Stderr, "Unable to open trimmed.gif for writing! Exiting.\n")
			os.Exit(1)
		}
		defer ftrim.Close()
		newgif := TrimGIF(start, end, gifstruct)
		err = gif.EncodeAll(ftrim, newgif)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			fmt.Fprintf(os.Stderr, "Unable to encode trimmed.gif! Exiting.\n")
			os.Exit(1)
		}
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
			fmt.Fprintf(os.Stderr, "Error! Could not create image: %s", filename)
		} else {
			if saveasjpeg {
				jpeg.Encode(outfile, img, &jpeg.Options{Quality: 100})
			} else {
				gif.Encode(outfile, img, &gif.Options{NumColors: 256})
			}
		}
	}
}

// TrimGIF will take a starting gif and trim it down to a start and end frame from the original, then return the new gif
func TrimGIF(start int, end int, original *gif.GIF) *gif.GIF {
	newgif := original
	newgif.Image = original.Image[start:end]
	newgif.Delay = original.Delay[start:end]
	newgif.Disposal = original.Disposal[start:end]
	return newgif
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

func normalizeTrimLength(startl, endl int, gifstruct *gif.GIF) (start, end int) {
	start = startl
	end = endl

	if start == -1 {
		start = 0
	}
	if end == -1 {
		end = len(gifstruct.Image)
	}
	if start > end {
		fmt.Fprintf(os.Stderr, "Invalid trim length! The end frame must be larger than the start frame. Exiting.\n")
		os.Exit(1)
	}
	return start, end
}
