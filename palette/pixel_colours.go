package palette

import (
	"image/color"
	"strings"

	"github.com/lucasb-eyer/go-colorful"
)

func Parse(palette string) []colorful.Color {

	arr := strings.Split(palette, ",")
	var colourPalette []colorful.Color
	for _, s := range arr {
		c, err := colorful.Hex(strings.TrimSpace(s))
		if err == nil {
			colourPalette = append(colourPalette, c)
		}
	}

	return colourPalette
}

func Match(colour color.RGBA, palette []colorful.Color) color.Color {

	if len(palette) == 0 {
		return colour
	}
	//convert the input color to a colorful color for comparison
	input, _ := colorful.MakeColor(colour)
	match := input
	// for each colour in the palette
	var closestDistance float64 = 100000.0 // todo what is the biggest distance?
	for _, pc := range palette {

		distance := input.DistanceLab(pc)

		if distance < closestDistance {
			match = pc
			closestDistance = distance
		}
	}

	return match
}
