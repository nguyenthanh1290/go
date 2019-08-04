package identicon

import (
	"errors"

	"crypto/sha1"

	"image"
	"image/color"
	"image/draw"
)

const (
	IconWidth = 250
	CellWidth = IconWidth / 5
)

func Identicon(data string) (image.Image, error) {
	if len(data) < 15 {
		return nil, errors.New("A text of at least 15 characters is required.")
	}

	// hashing the input data
	input := []byte(data)
	hash := sha1.Sum(input)

	// background is grey by default
	background := color.RGBA{240, 240, 240, 255}

	// the last three of the hash is used to generate the foreground color
	r := hash[17]
	g := hash[18]
	b := hash[19]
	foreground := color.RGBA{r, g, b, 255}

	icon := image.NewRGBA(image.Rect(0, 0, IconWidth, IconWidth))

	var p0, p1 image.Point

	// the first 15th of the hash are used to turn pixels on or off depending on even or odd values.
	for i := 0; i < 15; i++ {
		c := &image.Uniform{background}
		if hash[i]%2 == 0 {
			c = &image.Uniform{foreground}
		}

		// 5 columns
		if i < 5 { // 3th column is drawn first
			p0 = image.Pt(2*CellWidth, i*CellWidth)
			p1 = image.Pt(3*CellWidth, (i+1)*CellWidth)
			drawBlock(icon, p0, p1, c)
		} else if i < 10 { // then 2nd and 4th
			p0 = image.Pt(1*CellWidth, (i-5)*CellWidth)
			p1 = image.Pt(2*CellWidth, (i-4)*CellWidth)
			drawBlock(icon, p0, p1, c)

			p0 = image.Pt(3*CellWidth, (i-5)*CellWidth)
			p1 = image.Pt(4*CellWidth, (i-4)*CellWidth)
			drawBlock(icon, p0, p1, c)
		} else if i < 15 { // lastly, first and fifth
			p0 = image.Pt(0*CellWidth, (i-10)*CellWidth)
			p1 = image.Pt(1*CellWidth, (i-9)*CellWidth)
			drawBlock(icon, p0, p1, c)

			p0 = image.Pt(4*CellWidth, (i-10)*CellWidth)
			p1 = image.Pt(5*CellWidth, (i-9)*CellWidth)
			drawBlock(icon, p0, p1, c)
		}
	}

	return icon, nil
}

func drawBlock(dst draw.Image, p0 image.Point, p1 image.Point, src image.Image) {
	draw.Draw(dst, image.Rectangle{p0, p1}, src, image.ZP, draw.Src)
}
