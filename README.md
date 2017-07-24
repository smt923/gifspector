# gifspector
View details about a gif and split it into individual frames


## Usage
Running with just an input gif will simply print some information, run with `-s` to split the gif into individual frames, which by default save to `./out`

Basic usage, to simply print some information about the gif:
```
./gifspector input.gif
```

Splitting the file (-s), changing the default output folder (-o) and saving each frame as jpegs (-j)
```
./gifspector input.gif -s -o frames -j
```

