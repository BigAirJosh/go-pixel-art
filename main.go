package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	"image/png"
	"math"
	"os"

	"github.com/bigairjosh/pixelart/palette"
)

func main() {
	fmt.Println("Pixelart Rocks!!!")

	var inputFile string
	var outputFile string
	var pixelSize int
	var colourPaletteFlag string
	// #264653, #2a9d8f, #e9c46a, #f4a261, #e76f51

	flag.StringVar(&inputFile, "i", "/home/bigairjosh/wallpapers/0001.jpg", "Specify input file path")
	flag.StringVar(&outputFile, "o", "./pixelart.png", "Specify output file path")
	flag.IntVar(&pixelSize, "s", 25, "the size of each pixel in pixels (which doesn't make sense yet agreed)")
	flag.StringVar(&colourPaletteFlag, "c", "#264653, #2a9d8f, #e9c46a, #f4a261, #e76f51",
		"colour palette in #hex format as csv string")

	flag.Parse()

	pixelSquare := float64(pixelSize * pixelSize)

	colourPalette := palette.Parse(colourPaletteFlag)

	fmt.Println("processing " + inputFile)

	inputImage, err := getImageFromFilePath(inputFile)

	if err != nil {
		fmt.Println("error loading image " + err.Error())
		os.Exit(1)
	}

	bounds := inputImage.Bounds()

	fmt.Println(bounds)

	outputImage := image.NewRGBA(bounds)

	for x := bounds.Min.X; x < bounds.Max.X; x += pixelSize {
		for y := bounds.Min.Y; y < bounds.Max.Y; y += pixelSize {

			var sumR float64
			var sumG float64
			var sumB float64

			for px := x; px < x+pixelSize; px++ {
				for py := y; py < y+pixelSize; py++ {

					pixel := inputImage.At(px, py)
					col := color.RGBAModel.Convert(pixel).(color.RGBA)

					sumR += float64(col.R)
					sumG += float64(col.G)
					sumB += float64(col.B)
				}
			}

			avgR := math.Round(sumR / pixelSquare)
			avgG := math.Round(sumG / pixelSquare)
			avgB := math.Round(sumB / pixelSquare)
			// printColour(avgR, avgG, avgB, avgA)
			averageColour := color.RGBA{uint8(avgR), uint8(avgG), uint8(avgB), 0xff}
			pixelColour := palette.Match(averageColour, colourPalette)
			for px := x; px < x+pixelSize; px++ {
				for py := y; py < y+pixelSize; py++ {
					outputImage.Set(px, py, pixelColour)
				}
			}
		}
	}

	f, _ := os.Create(outputFile)
	png.Encode(f, outputImage)

}

//todo move to lib
func getImageFromFilePath(filePath string) (image.Image, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	i, _, err := image.Decode(f)
	return i, err
}

func ParseHexColor(s string) (c color.RGBA, err error) {
	c.A = 0xff
	switch len(s) {
	case 7:
		_, err = fmt.Sscanf(s, "#%02x%02x%02x", &c.R, &c.G, &c.B)
	case 4:
		_, err = fmt.Sscanf(s, "#%1x%1x%1x", &c.R, &c.G, &c.B)
		// Double the hex digits:
		c.R *= 17
		c.G *= 17
		c.B *= 17
	default:
		err = fmt.Errorf("invalid length, must be 7 or 4")

	}
	return
}
