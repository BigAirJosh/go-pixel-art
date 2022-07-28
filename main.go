package main

import (
	"flag"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

func main() {
	fmt.Println("Pixelart Rocks!!!")

	var inputFile string
	var outputFile string
	var pixelSize int

	flag.StringVar(&inputFile, "i", "/home/bigairjosh/wallpapers/0001.jpg", "Specify input file path")
	flag.StringVar(&outputFile, "o", "./pixelart.jpg", "Specify output file path")
	flag.IntVar(&pixelSize, "s", 5, "the size of each pixel in pixels (which doesn't make sense yet agreed)")

	fmt.Println("processing " + inputFile)

	inputImage, err := getImageFromFilePath(inputFile)

	if err != nil {
		fmt.Println("error loading image " + err.Error())
		os.Exit(1)
	}

	bounds := inputImage.Bounds()

	fmt.Println(bounds)

	var sumRed uint32
	sumRed = 0
	var sumGreen uint32
	sumGreen = 0
	var sumBlue uint32
	sumBlue = 0

	var total uint32 = 0

	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			nextColour := inputImage.At(x, y)
			r, g, b, _ := nextColour.RGBA()
			sumRed += r
			sumGreen += g
			sumBlue += b

			total += 1
		}
	}

	avgR := sumRed / total
	avgG := sumGreen / total
	avgB := sumBlue / total

	fmt.Printf("r:%dg:%db:%d", avgR, avgG, avgB)
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
