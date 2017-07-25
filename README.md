# gifspector
View details about a gif and split it into individual frames


## Usage
Running with just an input gif will simply print some information, run with `-s` to split the gif into individual frames, which by default save to `./out`

Basic usage, to simply print some information about the gif:
```bash
./gifspector input.gif
```

You can also trim a gif by specifying the start and, optionally, end frames to trim to (exclusive):
```bash
./gifspector input.gif 10 # trim from frame 10 to the last frame
```
```bash
./gifspector input.gif 8 10 # trim from frame 8 to frame 10 (2 frames)
```

Splitting the file (-s), changing the default output folder (-o) and saving each frame as jpegs (-j)
```bash
./gifspector input.gif -s -o frames -j
```

These are the stats printed out for a gif file:
```
--- GIF STATS: in.gif ---
Number of frames:
 33
Delay per frame (100ths / sec):
 [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
Loop count:
 0
Image size (height x width):
 240 x 560
 ```
