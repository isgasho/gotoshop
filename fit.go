package main

import (
	"code.google.com/p/graphics-go/graphics"
	"image"
	"math"
	"net/url"
)

type fitArgs struct {
	Width  int
	Height int
	Image  image.Image
}

func fitArgsBuilder(vals url.Values) (args, error) {
	width, err := getInt(vals, "width")
	if err != nil {
		return nil, err
	}

	height, err := getInt(vals, "height")
	if err != nil {
		return nil, err
	}

	src, err := getImage(vals, "url")
	if err != nil {
		return nil, err
	}

	return fitArgs{Width: width, Height: height, Image: src}, nil
}

func fit(a args) (image.Image, error) {
	args := a.(fitArgs)

	src := args.Image
	dstWidth := args.Width
	dstHeight := args.Height

	srcRect := src.Bounds()
	srcWidth := srcRect.Max.X - srcRect.Min.X
	srcHeight := srcRect.Max.Y - srcRect.Min.Y

	widthRatio := float64(dstWidth) / float64(srcWidth)
	heightRatio := float64(dstHeight) / float64(srcHeight)
	ratio := math.Min(widthRatio, heightRatio)
	if widthRatio < heightRatio {
		if float64(srcWidth)*heightRatio < float64(dstWidth) &&
			float64(srcHeight)*heightRatio < float64(dstHeight) {
			ratio = heightRatio
		}
	} else {
		if float64(srcWidth)*widthRatio < float64(dstWidth) &&
			float64(srcHeight)*widthRatio < float64(dstHeight) {
			ratio = widthRatio
		}
	}

	rect := image.Rect(
		int(float64(srcRect.Min.X)*ratio),
		int(float64(srcRect.Min.Y)*ratio),
		int(float64(srcRect.Max.X)*ratio),
		int(float64(srcRect.Max.Y)*ratio),
	)

	dst := image.NewRGBA(rect)

	if err := graphics.Scale(dst, src); err != nil {
		return nil, err
	}

	return dst, nil
}
