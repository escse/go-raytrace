package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"os"
	"path/filepath"

	"github.com/joshdk/quantize"
)

type Image struct {
	nx, ny int
	pixel  [][]Color
}

func (im Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (im Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, im.nx, im.ny)
}

func (im Image) At(x, y int) color.Color {
	return color.RGBA{im.pixel[x][y].r, im.pixel[x][y].g, im.pixel[x][y].b, 255}
}

func ReadPPMP3(file string) image.Image {
	f, _ := os.Open(file)
	defer f.Close()
	var buffer string
	fmt.Fscanf(f, "%s\n", &buffer)
	if buffer != "P3" {
		panic("format error")
	}
	var nx, ny int
	fmt.Fscanf(f, "%d %d\n", &nx, &ny)
	fmt.Fscanf(f, "%s\n", &buffer)

	var img Image
	img.nx = nx
	img.ny = ny
	img.pixel = make([][]Color, nx)
	for k := 0; k < nx; k++ {
		img.pixel[k] = make([]Color, ny)
	}

	for j := 0; j < ny; j++ {
		for i := 0; i < nx; i++ {
			var color Color
			fmt.Fscanf(f, "%d %d %d\n", &color.r, &color.g, &color.b)
			img.pixel[i][j] = color
		}
	}
	return img
}

func ConvertToGif(ppm image.Image) *image.Paletted {

	colors := quantize.Image(ppm, 10)
	palette := make([]color.Color, len(colors))
	for index, clr := range colors {
		palette[index] = clr
	}
	rect := ppm.Bounds()
	img := image.NewPaletted(rect, palette)

	for j := 0; j < ppm.(Image).ny; j++ {
		for i := 0; i < ppm.(Image).nx; i++ {
			img.Set(i, j, ppm.At(i, j))
		}
	}
	return img
}

func saveGif() {
	files, _ := filepath.Glob("image/*.ppm")

	animation := &gif.GIF{}
	for _, file := range files {
		print(file)

		ppm := ReadPPMP3(file)

		img := ConvertToGif(ppm)
		animation.Image = append(animation.Image, img)
		animation.Delay = append(animation.Delay, 2)
	}
	f, _ := os.OpenFile("img.gif", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	gif.EncodeAll(f, animation)
}
