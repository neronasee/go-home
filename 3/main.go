package main

import (
	"fmt"
	"homework/3/data"
	"homework/3/image"
)

func main() {
	// read raw data
	rawData := data.Read()

	// create empty (white) image with default image size
	img := image.NewImage()

	// create image printer
	printer := image.NewPrinter()

	fmt.Println("Empty image:")
	printer.Print(*img)

	// load raw image data to image
	img.Load(rawData)

	fmt.Println("Loaded image:")
	printer.Print(*img)
}
