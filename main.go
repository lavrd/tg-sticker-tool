package main

import (
	"flag"
	"fmt"
	"image/png"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/disintegration/imaging"
)

var src, dst string

func main() {
	flag.StringVar(&src, "src", "", "set up folder with images for resize and compress")
	flag.StringVar(&dst, "dst", "", "set up destination folder for new images")
	flag.Parse()

	if err := initializeFolder(); err != nil {
		log.Fatalln("Failed to initialize target folder; error =", err)
	}

	err := filepath.WalkDir(src, walk)
	if err != nil {
		log.Fatalln("Failed to prepare image; error =", err)
	}
}

func walk(path string, d fs.DirEntry, err error) error {
	if err != nil {
		return fmt.Errorf("failed to inspect entity while walking: %w", err)
	}
	if d.IsDir() || strings.Contains(path, ".DS") {
		return nil
	}
	if err = prepare(path); err != nil {
		return fmt.Errorf("failed to prepare image: %w", err)
	}
	return nil
}

func prepare(path string) error {
	img, err := imaging.Open(path)
	if err != nil {
		return fmt.Errorf("failed to open image: %w", err)
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
		return fmt.Errorf("failed to save image: %w", err)
	}

	log.Println("Saved prepared sticker:", filename)
	return nil
}

func initializeFolder() error {
	if _, err := os.Stat(dst); os.IsNotExist(err) {
		if err := os.Mkdir(dst, os.ModePerm); err != nil {
			return fmt.Errorf("failed to create directory: %w", err)
		}
	}
	return nil
}
