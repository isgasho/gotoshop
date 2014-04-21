package main

import (
	"code.google.com/p/graphics-go/graphics"
	"image"
	"net/url"
)

type blurArgs struct {
	Image  image.Image
	StdDev float64
	Size   int
}

func blurArgsBuilder(vals url.Values) (args, error) {
	stdDev, err := getFloat(vals, "std_dev")
	if err != nil {
		return nil, err
	}

	size, _ := getInt(vals, "size")

	src, err := getImage(vals, "url")
	if err != nil {
		return nil, err
	}

	return blurArgs{Image: src, StdDev: stdDev, Size: size}, nil
}

func blur(a args) (image.Image, error) {
	args := a.(blurArgs)

	src := args.Image
	dst := image.NewNRGBA(src.Bounds())
	opt := &graphics.BlurOptions{StdDev: args.StdDev, Size: args.Size}

	if err := graphics.Blur(dst, src, opt); err != nil {
		return nil, err
	}

	return dst, nil
}
