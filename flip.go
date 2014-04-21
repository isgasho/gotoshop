package main

import (
	"code.google.com/p/graphics-go/graphics"
	"code.google.com/p/graphics-go/graphics/interp"
	"errors"
	"image"
	"net/url"
)

type flipArgs struct {
	Image image.Image
	Axis  string
}

func flipArgsBuilder(vals url.Values) (args, error) {
	axis, err := getString(vals, "axis")
	if err != nil {
		return nil, err
	}

	if axis != "x" && axis != "y" {
		return nil, errors.New("axis must be x or y")
	}

	src, err := getImage(vals, "url")
	if err != nil {
		return nil, err
	}

	return flipArgs{Image: src, Axis: axis}, nil
}

func flip(a args) (image.Image, error) {
	args := a.(flipArgs)

	src := args.Image
	sb := src.Bounds()
	dst := image.NewRGBA(sb)

	var flipper graphics.Affine

	if args.Axis == "x" {
		flipper = graphics.Affine{
			1, 0, 0,
			0, -1, float64(sb.Dy()),
			0, 0, 1,
		}
	} else {
		flipper = graphics.Affine{
			-1, 0, float64(sb.Dx()),
			0, 1, 0,
			0, 0, 1,
		}
	}
	flipper.Transform(dst, src, interp.Bilinear)

	return dst, nil
}
