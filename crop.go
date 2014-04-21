package main

import (
	"fmt"
	"image"
	"image/draw"
	"net/url"
)

type cropArgs struct {
	Image image.Image
	Rect  image.Rectangle
}

func cropArgsBuilder(vals url.Values) (args, error) {
	x0, err := getInt(vals, "x0")
	if err != nil {
		return nil, err
	}

	y0, err := getInt(vals, "y0")
	if err != nil {
		return nil, err
	}

	x1, err := getInt(vals, "x1")
	if err != nil {
		return nil, err
	}

	y1, err := getInt(vals, "y1")
	if err != nil {
		return nil, err
	}

	src, err := getImage(vals, "url")
	if err != nil {
		return nil, err
	}

	rect := image.Rect(x0, y0, x1, y1)

	if image.Pt(x0+1, y0+1).In(src.Bounds()) == false {
		return nil, fmt.Errorf("%s is not in image %s",
			image.Pt(x0, y0), src.Bounds().String())
	}

	if image.Pt(x1-1, y1-1).In(src.Bounds()) == false {
		return nil, fmt.Errorf("%s is not in image %s",
			image.Pt(x1, y1), src.Bounds().String())
	}

	return cropArgs{Rect: rect, Image: src}, nil
}

func crop(a args) (image.Image, error) {
	args := a.(cropArgs)

	src := args.Image
	dst := image.NewRGBA(args.Rect)

	db := dst.Bounds()
	pt := image.Pt(db.Min.X, db.Min.Y)

	draw.Draw(dst, db, src, pt, draw.Src)

	return dst, nil
}
