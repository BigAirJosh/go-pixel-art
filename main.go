package main

import (
	"flag"
	"image"
	"image/color"
	_ "image/jpeg"
	"image/png"
	"log"
	"math"
	"os"

	"github.com/bigairjosh/pixelart/palette"
)

var verbose bool

func main() {

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
	flag.BoolVar(&verbose, "v", false, "Verbose processing output")

	flag.Parse()

	logIfVerbose("Pixelart coming up...")

	colourPalette := palette.Parse(colourPaletteFlag)

	logIfVerbose("loading input file [", inputFile, "]")

	inputImage, err := getImageFromFilePath(inputFile)

	if err != nil {
		log.Fatal("error loading image ", err.Error())
	}

	bounds := inputImage.Bounds()

	outputImage := image.NewRGBA(bounds)

	pixelSquare := float64(pixelSize * pixelSize)

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

	savePixelArt(outputFile, outputImage)

}

func getImageFromFilePath(filePath string) (image.Image, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	i, _, err := image.Decode(f)
	return i, err
}

func savePixelArt(outputFile string, outputImage image.Image) {

	logIfVerbose("Saving output file [", outputFile, "]")

	file, err := os.Create(outputFile)
	if err != nil {
		log.Fatal("Error creating output file [", outputFile, "] message [", err.Error(), "]")
	}
	defer file.Close()
	err = png.Encode(file, outputImage)
	if err != nil {
		log.Fatal("Error creating output file [", outputFile, "] message [", err.Error(), "]")
	}

	fileInfo, err := file.Stat()
	if err != nil {
		log.Fatal("Error creating output file [", outputFile, "] message [", err.Error(), "]")
	}
	logIfVerbose("Saved [", fileInfo.Size(), "] bytes")

}

func logIfVerbose(v ...any) {

	if verbose {
		log.Println(v...)
	}
}
