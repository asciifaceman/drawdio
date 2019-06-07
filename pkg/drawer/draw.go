package drawer

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
	"time"
)

// Core Colors
var (
	WHITE = color.RGBA{R: 255, G: 255, B: 255, A: 255}
	BLACK = color.RGBA{R: 0, G: 0, B: 0, A: 255}
	RED   = color.RGBA{R: 255, G: 0, B: 0, A: 255}
)

/*
-1 = red
-1..0 = gradient red to black
0 = black
0..1 = gradient black to white
1 = white
*/

// Sound ...
func Sound(filename string, s []float64) error {
	size := int(math.Sqrt(float64(len(s))))
	w, h := size, size

	dst := image.NewRGBA(image.Rect(0, 0, w, h)) //*NRGBA (image.Image interface)

	var iter int
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			col := linearGradient(s[iter])
			dst.Set(x, y, col)
			iter++
		}
	}
	img, _ := os.Create(fmt.Sprintf("%s-%d.png", filename, time.Now().Second()))
	defer img.Close()
	png.Encode(img, dst)
	return nil
}

func linearGradient(sample float64) color.RGBA {
	/*
		if less than one, gradient between -1 and 0
		if greater than one gradient between 0 1
		else 0
	*/
	d := math.Abs(sample)
	if sample < 0 {
		return color.RGBA{
			R: RED.R + uint8(d*float64(BLACK.R-RED.R)),
			G: RED.G + uint8(d*float64(BLACK.G-RED.G)),
			B: RED.B + uint8(d*float64(BLACK.B-RED.B)),
			A: RED.A + uint8(d*float64(BLACK.A-RED.A)),
		}
	} else if sample > 0 {
		return color.RGBA{
			R: BLACK.R + uint8(d*float64(WHITE.R-BLACK.R)),
			G: BLACK.G + uint8(d*float64(WHITE.G-BLACK.G)),
			B: BLACK.B + uint8(d*float64(WHITE.B-BLACK.B)),
			A: BLACK.A + uint8(d*float64(WHITE.A-BLACK.A)),
		}
	}

	return WHITE
}
