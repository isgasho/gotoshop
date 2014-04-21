package main

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"net/http"
	"net/url"
	"strconv"
)

func getString(vals url.Values, key string) (string, error) {
	val := vals.Get(key)
	if val == "" {
		return "", fmt.Errorf("%s is missing", key)
	}

	return val, nil
}

func getInt(vals url.Values, key string) (int, error) {
	val := vals.Get(key)
	if val == "" {
		return 0, fmt.Errorf("%s is missing", key)
	}

	x, err := strconv.Atoi(val)
	if err != nil {
		return 0, fmt.Errorf("%s is not int", key)
	}

	return x, nil
}

func getFloat(vals url.Values, key string) (float64, error) {
	val := vals.Get(key)
	if val == "" {
		return 0, fmt.Errorf("%s is missing", key)
	}

	f, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return 0, fmt.Errorf("%s is not int", key)
	}

	return f, nil
}

func getImage(vals url.Values, key string) (image.Image, error) {
	url := vals.Get(key)
	if url == "" {
		return nil, fmt.Errorf("%s is missing", key)
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if statusCode := resp.StatusCode; statusCode != 200 {
		return nil, fmt.Errorf(resp.Status)
	}

	image, _, err := image.Decode(resp.Body)
	if err != nil {
		return nil, err
	}

	return image, nil
}
