package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	"image/png"
	_ "image/png"
	"math"
	"os"
)

func main() {
	fmt.Println("Pixelart Rocks!!!")

	var inputFile string
	var outputFile string
	var pixelSize int

	flag.StringVar(&inputFile, "i", "/home/bigairjosh/wallpapers/0001.jpg", "Specify input file path")
	flag.StringVar(&outputFile, "o", "./pixelart.png", "Specify output file path")
	flag.IntVar(&pixelSize, "s", 25, "the size of each pixel in pixels (which doesn't make sense yet agreed)")

	flag.Parse()

	pixelSquare := float64(pixelSize * pixelSize)

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
			pixelColour := color.RGBA{uint8(avgR), uint8(avgG), uint8(avgB), 0xff}

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

func printColour(r uint32, g uint32, b uint32, a uint32) {
	fmt.Printf("r:%dg:%db:%da:%d\n", r, g, b, a)
}

//todo move to lib
func getImageFromFilePath(filePath string) (image.Image, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	image, _, err := image.Decode(f)
	return image, err
}
