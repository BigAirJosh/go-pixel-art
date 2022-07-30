package palette

import (
	"image/color"
	"testing"

	"github.com/lucasb-eyer/go-colorful"
)

func TestEmptyPalette(t *testing.T) {

	palette := Parse("")

	if len(palette) > 0 {
		t.Fatal("should be empty slice")
	}
}

func TestSingleInvalidColour(t *testing.T) {

	palette := Parse("invalid")

	if len(palette) > 0 {
		t.Fatal("should be empty slice")
	}

}

func TestSingleValidColour(t *testing.T) {

	palette := Parse("#517AB8")

	if len(palette) != 1 {
		t.Fatal("should have one singel colour")
	}
	colour, _ := colorful.Hex("#517AB8")
	if palette[0] != colour {
		t.Fatal("not the right colour")
	}
}

func TestMultipleValidColoursWithSpaces(t *testing.T) {

	palette := Parse("#517AB8 , #517AB8")

	if len(palette) != 2 {
		t.Fatal("should have two colours")
	}
	colour, _ := colorful.Hex("#517AB8")
	if palette[0] != colour || palette[1] != colour {
		t.Fatal("not the right colour")
	}
}

func TestMatchWithEmptyPaletteReturnsColor(t *testing.T) {

	palette := Parse("")

	c := color.RGBA{0, 0, 0, 255}

	if Match(c, palette) != c {
		t.Fatal("no palette should return input colour")
	}

}

func TestMatchWithSinglePaletteReturnsPaletteColor(t *testing.T) {

	palette := Parse("#76FE9A")

	c := color.RGBA{0, 0, 0, 0}

	pr, pg, pb, pa := palette[0].RGBA()
	r, g, b, a := Match(c, palette).RGBA()

	if r != pr {
		t.Fatal("r colour not matched", r, pr)
	}
	if g != pg {
		t.Fatal("g colour not matched")
	}
	if b != pb {
		t.Fatal("b colour not matched")
	}
	if a != pa {
		t.Fatal("a colour not matched")
	}
}

func TestMatchWithMultiplePaletteReturnsPaletteColor(t *testing.T) {

	redish := color.RGBA{255, 0, 0, 255}
	MatchWithMultiplePaletteReturnsPaletteColor(0, redish, t)
	greenish := color.RGBA{25, 98, 11, 255}
	MatchWithMultiplePaletteReturnsPaletteColor(1, greenish, t)
	blueish := color.RGBA{10, 1, 65, 255}
	MatchWithMultiplePaletteReturnsPaletteColor(2, blueish, t)

}

func MatchWithMultiplePaletteReturnsPaletteColor(paletteIndex int, color color.RGBA, t *testing.T) {

	palette := Parse("#FF0000, #00FF00, #0000FF")

	pr, pg, pb, pa := palette[paletteIndex].RGBA()
	r, g, b, a := Match(color, palette).RGBA()

	if r != pr {
		t.Fatal("r colour not matched", r, pr)
	}
	if g != pg {
		t.Fatal("g colour not matched")
	}
	if b != pb {
		t.Fatal("b colour not matched")
	}
	if a != pa {
		t.Fatal("a colour not matched")
	}

}
