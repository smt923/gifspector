# gifspector
View details about a gif and split it into individual frames


## Usage
Running with just an input gif will simply print some information, run with `-s` to split the gif into individual frames, which by default save to `./out`

Basic usage, to simply print some information about the gif:
```
./gifspector input.gif
```

Splitting the file (-s), changing the default output folder (-o) and saving each frame as a jpeg (-j)
```
./gifspector input.gif -s -o frames -j
```

These are the stats it prints out for a gif
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
