package main

import (
	"flag"
	"fmt"
	"image/png"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/disintegration/imaging"
)

func main() {
	var src, dst string
	flag.StringVar(&src, "src", "", "set up images paths for resize and compress; enter paths separated by commas")
	flag.StringVar(&dst, "dst", "", "set up destination folder for new images")
	flag.Parse()

	for _, path := range strings.Split(src, ",") {
		img, err := imaging.Open(path)
		if err != nil {
			log.Printf(`failed to open image by "%s" path; error = %v\n'`, path, err)
			os.Exit(1)
		}

		if img.Bounds().Max.X > img.Bounds().Max.Y {
			img = imaging.Resize(img, 512, 0, imaging.Lanczos)
		} else {
			img = imaging.Resize(img, 0, 512, imaging.Lanczos)
		}

		_, filename := filepath.Split(path)
		ext := filepath.Ext(path)
		dst := fmt.Sprintf("%s/%s.png", dst, filename[:len(filename)-len(ext)])

		if err := imaging.Save(img, dst, imaging.PNGCompressionLevel(png.BestCompression)); err != nil {
			log.Printf(`failed to save image by "%s" path; error = %v\n`, path, err)
			os.Exit(1)
		}
	}
}
