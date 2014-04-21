package main

import (
	"code.google.com/p/graphics-go/graphics"
	"errors"
	"image"
	"net/url"
)

type scaleArgs struct {
	Image image.Image
	Ratio float64
}

func scaleArgsBuilder(vals url.Values) (args, error) {
	ratio, err := getFloat(vals, "ratio")
	if err != nil {
		return nil, err
	}
	if ratio <= 0 {
		return nil, errors.New("ratio must be greater than 0")
	}

	src, err := getImage(vals, "url")
	if err != nil {
		return nil, err
	}

	return scaleArgs{Image: src, Ratio: ratio}, nil
}

func scale(a args) (image.Image, error) {
	args := a.(scaleArgs)

	src := args.Image
	srcRect := src.Bounds()

	ratio := args.Ratio
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
