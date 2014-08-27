package main

import (
	"fmt"
	"net/http"
)

func inspect(writer http.ResponseWriter, request *http.Request) {
	params := request.URL.Query()

	src, err := getImage(params, "url")
	if err != nil {
		http.Error(writer, err.Error(), 406)
		return
	}

	bounds := src.Bounds()
	width := bounds.Max.X - bounds.Min.X
	height := bounds.Max.Y - bounds.Min.Y

	fmt.Fprintf(writer, `{"width": %d, "height": %d}`, width, height)
}
