package main

import (
	"log"
	"net/http"
)

func main() {
	config := NewConfig()

	http.HandleFunc("/blur", buildHandler(
		blurArgsBuilder,
		blur,
	))

	http.HandleFunc("/crop", buildHandler(
		cropArgsBuilder,
		crop,
	))

	http.HandleFunc("/flip", buildHandler(
		flipArgsBuilder,
		flip,
	))

	http.HandleFunc("/rotate", buildHandler(
		rotateArgsBuilder,
		rotate,
	))

	http.HandleFunc("/thumbnail", buildHandler(
		thumbnailArgsBuilder,
		thumbnail,
	))

	http.HandleFunc("/scale", buildHandler(
		scaleArgsBuilder,
		scale,
	))

	log.Fatal(http.ListenAndServe(config.Addr, nil))
}
