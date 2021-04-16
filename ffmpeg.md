# ffmpeg

## find audio format in a video

Redirect *stderr* to *stdout* for pipe with `2>&1`

`ffprobe -hide_banner video.mp4 2>&1 | find "Audio:"`

## extract audio from video

Assumes extracting to same audio format.
Aithout `-acodec copy` it reencodes, and might compress.

`ffmpeg -i input.mp4 -vn -acodec copy output.aac`

## convert audio format

`ffmpeg -i input.aac output.wav`

## resample audio frequency (in this case 16kHz)

`ffmpeg -i input.wav -ar 16000 output.wav`
