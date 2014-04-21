package main

import (
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"net/http"
	"net/url"
)

type args interface{}

type paramsBuilder func(url.Values) (args, error)

type handler func(args) (image.Image, error)

func buildHandler(builder paramsBuilder, handler handler) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		params := request.URL.Query()
		args, err := builder(params)
		if err != nil {
			http.Error(writer, err.Error(), 406)
			return
		}

		image, err := handler(args)

		accept := params.Get("Accept")
		switch accept {
		case "image/jpeg":
			jpeg.Encode(writer, image, &jpeg.Options{Quality: 100})
		case "image/gif":
			gif.Encode(writer, image, &gif.Options{NumColors: 256})
		default:
			png.Encode(writer, image)
		}
	}
}
