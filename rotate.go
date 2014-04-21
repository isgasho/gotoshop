package main

import (
	"code.google.com/p/graphics-go/graphics"
	"image"
	"math"
	"net/url"
)

type rotateArgs struct {
	Image image.Image
	Angle float64
}

func rotateArgsBuilder(vals url.Values) (args, error) {
	angle, err := getFloat(vals, "angle")
	if err != nil {
		return nil, err
	}

	src, err := getImage(vals, "url")
	if err != nil {
		return nil, err
	}

	return rotateArgs{Image: src, Angle: angle}, nil
}

func rotate(a args) (image.Image, error) {
	args := a.(rotateArgs)

	src := args.Image
	dst := image.NewRGBA(src.Bounds())
	opt := &graphics.RotateOptions{Angle: args.Angle * (math.Pi / 180)}

	if err := graphics.Rotate(dst, src, opt); err != nil {
		return nil, err
	}

	return dst, nil
}
