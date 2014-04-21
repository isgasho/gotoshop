package main

import (
	"code.google.com/p/graphics-go/graphics"
	"image"
	"net/url"
)

type thumbnailArgs struct {
	Image image.Image
	Rect  image.Rectangle
}

func thumbnailArgsBuilder(vals url.Values) (args, error) {
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

	rect := image.Rect(0, 0, width, height)

	return thumbnailArgs{Rect: rect, Image: src}, nil
}

func thumbnail(p args) (image.Image, error) {
	args := p.(thumbnailArgs)

	src := args.Image
	dst := image.NewRGBA(args.Rect)

	if err := graphics.Thumbnail(dst, src); err != nil {
		return nil, err
	}

	return dst, nil
}
