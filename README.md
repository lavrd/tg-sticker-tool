# tg-sticker-tool

Tool for prepare images to use as telegram stickers.

Tool resizes images up to 512x512px, then compresses them and saves them as PNG.

TODO:
- write full instruction how to create sticker
- use file or folder to load and prepare images

### Usage

- images should be in PNG format

```shell
go build -mod vendor -o ./tg-sticker-tool
./tg-sticker-tool -src /Users/lavr/Desktop/IMG_3734.jpg,/Users/lavr/Desktop/IMG_3773.jpg,/Users/lavr/Desktop/IMG_3832.jpg -dst /Users/lavr/Downloads
```
